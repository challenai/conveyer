package table

import "github.com/challenai/conveyer/pkg/transform"

type TableManager interface {
	SetTableName(tableName string)
	SetFields(fields []transform.Field)
	ScanFields(queryDSL string) ([]transform.Field, error)
	CreateTable() error
	TruncateTable() error
	RemoveTable() error
	CreateWatermarkField() error
	RemoveWatermarkField() error
	GetFieldsCount() int
}
