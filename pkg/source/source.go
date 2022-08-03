package source

import (
	"github.com/challenai/conveyer/pkg/table"
)

type Source interface {
	Open() error
	Count(dsl string) (int64, error)
	Query(dsl string, offset, limit int) (any, error)
	Close() error
	table.TableManager
}
