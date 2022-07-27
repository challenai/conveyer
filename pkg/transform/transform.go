package transform

import "github.com/challenai/conveyer/pkg/codec"

type Type string

type Transformer interface {
	// notify
	SourceTypes() []Type

	SinkTypes() []Type
	SetSinkTypes([]Type)

	Castable() error
	Cast(src []codec.Bytes) ([]any, error)
	CastTypes() []Type

	CastableType(src Type, dest Type) error
	CastByte(src codec.Bytes) (any, error)
	CastType(Type) Type
}
