package sink

import (
	"github.com/challenai/conveyer/pkg/table"
)

type Sink interface {
	Open() error
	WriteRows(any) error
	Close() error
	table.TableManager
	// generateBatchWatermark() string
}
