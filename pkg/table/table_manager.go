package table

import "github.com/challenai/conveyer/pkg/transform"

type TableManager interface {
	CreateTable() error
	TruncateTable() error
	RemoveTable() error
	ScanTypes(dsl string) ([]transform.Type, error)
	CreateWatermarkField() error
	RemoveWatermarkField() error
}
