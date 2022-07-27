package table

import (
	"os"

	"github.com/challenai/conveyer/pkg/table"
	"github.com/challenai/conveyer/pkg/transform"
)

type tableCSV struct {
	f           os.File
	fieldsCount int
	fields      []transform.Field
}

func NewCSVTableManager(f os.File) table.TableManager {
	tc := &tableCSV{
		f: f,
	}
	return tc
}

func (t *tableCSV) ScanFields(dsl string) ([]transform.Field, error) {
	if t.fieldsCount > 0 {
		return t.fields, nil
	}

	// todo

	return t.fields, nil
}

func (t *tableCSV) SetFields(fields []transform.Field) {
	t.fieldsCount = len(fields)
	t.fields = fields
}

func (t *tableCSV) CreateTable() error {
	return nil
}

func (t *tableCSV) TruncateTable() error {
	return t.f.Truncate(0)
}

func (t *tableCSV) RemoveTable() error {
	return t.TruncateTable()
}

func (t *tableCSV) CreateWatermarkField() error {
	return nil
}

func (t *tableCSV) RemoveWatermarkField() error {
	return nil
}

func (t *tableCSV) SetTableName(string) {}

func (t *tableCSV) GetFieldsCount() int {
	return t.fieldsCount
}
