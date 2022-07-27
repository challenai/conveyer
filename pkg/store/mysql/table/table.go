package table

import (
	"database/sql"
	"fmt"

	"github.com/challenai/conveyer/pkg/table"
	"github.com/challenai/conveyer/pkg/transform"
)

type tableMySQL struct {
	db        *sql.DB
	tableName string
	types     []transform.Type
}

func NewMySQLTableManager(db *sql.DB, tableName string) table.TableManager {
	tm := &tableMySQL{
		db:        db,
		tableName: tableName,
	}
	return tm
}

func (t *tableMySQL) ScanTypes(dsl string) ([]transform.Type, error) {
	if len(t.types) > 0 {
		return t.types, nil
	}
	rows, err := t.db.Query(fmt.Sprintf("%s LIMIT 0", dsl))
	if err != nil {
		return nil, err
	}
	cts, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	t.types = make([]transform.Type, 0, len(cts))
	for _, v := range cts {
		t.types = append(t.types, transform.Type(v.DatabaseTypeName()))
	}
	return t.types, nil
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
