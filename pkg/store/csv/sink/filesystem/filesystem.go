package filesystem

import (
	"encoding/csv"
	"errors"
	"os"

	"github.com/challenai/conveyer/pkg/codec"
	"github.com/challenai/conveyer/pkg/graph/desc"
	"github.com/challenai/conveyer/pkg/sink"
	csvtable "github.com/challenai/conveyer/pkg/store/csv/table"
	mysqltransform "github.com/challenai/conveyer/pkg/store/csv/transform/mysql"
	"github.com/challenai/conveyer/pkg/table"
	"github.com/challenai/conveyer/pkg/transform"
)

const (
	KindCSV = "csv"
)

type filesystemSink struct {
	f     *os.File
	w     *csv.Writer
	desc  desc.Sink
	extra *FilesystemSinkExtraDescription
	table.TableManager
	trans transform.Transformer
}

type FilesystemSinkExtraDescription struct {
	CSVName    string
	Path       string
	WithHeader bool
	Headers    string
}

func NewFilesystemCSVSink(desc desc.Sink) (sink.Sink, error) {
	extra, ok := desc.Extra.(FilesystemSinkExtraDescription)
	if !ok {
		return nil, errors.New("bad sink extra description")
	}

	s := &filesystemSink{
		desc:  desc,
		extra: &extra,
	}

	err := extra.validate()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *filesystemSink) Open() error {
	var err error
	s.f, err = os.OpenFile(s.extra.CSVName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	s.TableManager = csvtable.NewCSVTableManager(s.f)
	s.w = csv.NewWriter(s.f)

	return nil
}

func (fs *filesystemSink) InitTransformer(src []transform.Field) {
	fs.trans = mysqltransform.NewMysqlTransformer(src)
	fs.SetFields(fs.trans.CastFields())
}

func (fs *filesystemSink) WriteRows(rows [][]codec.Bytes) error {
	for _, rowBytes := range rows {
		err := fs.WriteRow(rowBytes)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fs *filesystemSink) WriteRow(rowBytes []codec.Bytes) error {
	row, err := fs.trans.Cast(rowBytes)
	if err != nil {
		return err
	}

	temp := make([]string, 0, len(row))
	for _, v := range row {
		temp = append(temp, v.(string))
	}

	err = fs.w.Write(temp)
	if err != nil {
		return err
	}

	return nil
}

func (fs *filesystemSink) Close() error {
	fs.w.Flush()
	err := fs.f.Sync()
	if err != nil {
		return err
	}

	return nil
}

func (desc *FilesystemSinkExtraDescription) validate() error {
	if desc.CSVName == "" {
		return errors.New("config error: csv file need a name, like \"dump.csv\"")
	}
	return nil
}

func (desc *FilesystemSinkExtraDescription) setDefault() {

}
