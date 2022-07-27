package filesystem

import (
	"errors"
	"os"

	"github.com/challenai/conveyer/pkg/sink"
	"github.com/challenai/conveyer/pkg/table"
	"github.com/challenai/conveyer/pkg/transform"
)

type filesystemSink struct {
	f    os.File
	desc FilesystemSinkDescription
	table.TableManager
}

type FilesystemSinkDescription struct {
	CSVName    string
	Path       string
	WithHeader bool
	Headers    string
}

func NewFilesystemCSVSink(desc FilesystemSinkDescription) (sink.Sink, error) {
	err := desc.validate()
	if err != nil {
		return nil, err
	}

	return &filesystemSink{
		desc: desc,
	}, nil
}

func (fs *filesystemSink) GetTransformer(src []transform.Type, dest []transform.Type) transform.Transformer {
	return nil
}

func (fs *filesystemSink) WriteRows([][]any) error {
	return nil
}

func (fs *filesystemSink) WriteRow([]any) error {
	return nil
}

func (fs *filesystemSink) Close() error {
	return nil
}

func (desc *FilesystemSinkDescription) validate() error {
	if desc.CSVName == "" {
		return errors.New("config error: csv file need a name, like \"dump.csv\"")
	}
	return nil
}

func (desc *FilesystemSinkDescription) setDefault() {

}
