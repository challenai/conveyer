package source

import (
	"github.com/challenai/conveyer/pkg/codec"
	"github.com/challenai/conveyer/pkg/table"
)

type Source interface {
	Count(dsl string) (int, error)
	Query(dsl string, offset, limit int) ([][]codec.Bytes, error)
	Close() error
	table.TableManager
}
