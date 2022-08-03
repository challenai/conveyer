package sink

import (
	"github.com/challenai/conveyer/pkg/codec"
	"github.com/challenai/conveyer/pkg/table"
	"github.com/challenai/conveyer/pkg/transform"
)

type Sink interface {
	Open() error
	InitTransformer(src []transform.Field)
	WriteRows([][]codec.Bytes) error
	WriteRow([]codec.Bytes) error
	Close() error
	table.TableManager
	// generateBatchWatermark() string
}
