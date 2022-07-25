package transform

import "github.com/challenai/conveyer/pkg/codec"

type Type byte

type Transformer interface {
	// notify
	Castable(src Type, dest Type) error
	CastableRow(src []Type, dest []Type) error
	Cast(src codec.Bytes) (any, error)
	CastRow(src []codec.Bytes) ([]any, error)
	CastType(Type) Type
	CastTypesRow([]Type) []Type
}
