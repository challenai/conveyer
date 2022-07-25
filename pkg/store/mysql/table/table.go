package table

import (
	"github.com/challenai/conveyer/pkg/table"
)

type tableMySQL struct {
	databaseName string
	tableName    string
}

func NewMySQLTableManager() table.TableManager {
	tm := &tableMySQL{}
	return tm
}

func (t *tableMySQL) ScanTypes() error {
	return nil
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
