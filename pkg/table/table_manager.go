package table

type TableManager interface {
	CreateTable() error
	TruncateTable() error
	RemoveTable() error
	ScanTypes() error
	CreateWatermarkField() error
	RemoveWatermarkField() error
}
