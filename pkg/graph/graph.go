package graph

import (
	"errors"
	"fmt"

	"github.com/challenai/conveyer/pkg/belt"
	"github.com/challenai/conveyer/pkg/graph/desc"
	"github.com/challenai/conveyer/pkg/graph/meta"
	"github.com/challenai/conveyer/pkg/graph/state"
	"github.com/challenai/conveyer/pkg/sink"
	"github.com/challenai/conveyer/pkg/source"
	filesystemcsv "github.com/challenai/conveyer/pkg/store/csv/sink/filesystem"
	mysqlsource "github.com/challenai/conveyer/pkg/store/mysql/source"
)

type Graph struct {
	meta.Meta
	desc.Description
	*state.State
	src    source.Source
	sink   sink.Sink
	belt   *belt.Belt
	errch  chan error
	nextch chan struct{}
}

func (g *Graph) Run() error {
	var err error

	g.errch = make(chan error)
	g.nextch = make(chan struct{})

	switch g.Description.Source.Kind {
	case mysqlsource.KindMysql:
		g.src, err = mysqlsource.NewMysqlSource(g.Description.Source)
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported source")
	}

	err = g.src.Open()
	if err != nil {
		return err
	}
	defer g.src.Close()

	srcFields, err := g.src.ScanFields(g.Description.Source.DSL)
	if err != nil {
		return err
	}

	switch g.Description.Sink.Kind {
	case filesystemcsv.KindCSV:
		g.sink, err = filesystemcsv.NewFilesystemCSVSink(g.Description.Sink)
	default:
		return errors.New("unsupported sink")
	}

	err = g.sink.Open()
	if err != nil {
		return err
	}
	defer g.sink.Close()

	g.sink.InitTransformer(srcFields)
	if err != nil {
		return err
	}

	g.belt = belt.NewBelt(1)

	fields, err := g.src.ScanFields(g.Source.DSL)
	if err != nil {
		return err
	}
	g.src.SetFields(fields)

	err = g.deliever()
	if err != nil {
		return err
	}

	return nil
}

func (g *Graph) deliever() error {
	go func() {
		rows := g.belt.Accept()
		err := g.sink.WriteRows(rows)
		if err != nil {
			g.errch <- err
		}
		g.nextch <- struct{}{}
	}()

	go func() {
		// for-loop
		rows, err := g.src.Query(g.Source.DSL, 0, 10)
		if err != nil {
			g.errch <- err
		}
		fmt.Println("got rows: ", len(rows))

		err = g.belt.Deliever(rows)
		if err != nil {
			g.errch <- err
		}
	}()

	for {
		select {
		case err := <-g.errch:
			return err
		case <-g.nextch:
			return nil
		}
	}
}

func (g *Graph) check() error {
	return nil
}
