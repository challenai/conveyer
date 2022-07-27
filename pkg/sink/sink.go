package sink

import (
	"github.com/challenai/conveyer/pkg/table"
	"github.com/challenai/conveyer/pkg/transform"
)

type Sink interface {
	GetTransformer(src []transform.Type, dest []transform.Type) transform.Transformer
	WriteRows([][]any) error
	WriteRow([]any) error
	Close() error
	table.TableManager
	// generateBatchWatermark() string
}
