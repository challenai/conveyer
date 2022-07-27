package table

import (
	"database/sql"
	"fmt"

	"github.com/challenai/conveyer/pkg/table"
	"github.com/challenai/conveyer/pkg/transform"
)

type tableMySQL struct {
	db          *sql.DB
	fieldsCount int
	table       string
	fields      []transform.Field
}

func NewMySQLTableManager(db *sql.DB) table.TableManager {
	tm := &tableMySQL{
		db: db,
	}
	return tm
}

func (t *tableMySQL) ScanFields(dsl string) ([]transform.Field, error) {
	if t.fieldsCount > 0 {
		return t.fields, nil
	}
	rows, err := t.db.Query(fmt.Sprintf("%s LIMIT 0", dsl))
	if err != nil {
		return nil, err
	}
	cts, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	t.fieldsCount = len(cts)
	t.fields = make([]transform.Field, 0, t.fieldsCount)
	for _, v := range cts {
		t.fields = append(t.fields, transform.Field{
			Name: v.Name(),
			Type: transform.Type(v.DatabaseTypeName()),
		})
	}
	return t.fields, nil
}

func (t *tableMySQL) SetFields(fields []transform.Field) {
	t.fieldsCount = len(fields)
	t.fields = fields
}

func (t *tableMySQL) CreateTable() error {
	return nil
}

func (t *tableMySQL) TruncateTable() error {
	return nil
}

func (t *tableMySQL) RemoveTable() error {
	return nil
}

func (t *tableMySQL) CreateWatermarkField() error {
	return nil
}

func (t *tableMySQL) RemoveWatermarkField() error {
	return nil
}

func (t *tableMySQL) SetTableName(tableName string) {
	t.table = tableName
}

func (t *tableMySQL) GetFieldsCount() int {
	return t.fieldsCount
}
