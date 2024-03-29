package transform

import "github.com/challenai/conveyer/pkg/codec"

type Type string

type Field struct {
	Name string
	Type
}

type Transformer interface {
	// notify
	SourceFields() []Field

	SinkFields() []Field
	SetSinkFields([]Field)

	Castable() error
	Cast(src []codec.Bytes) ([]any, error)
	CastFields() []Field

	CastableType(src Field, dest Field) error
	CastBytes(index int, src codec.Bytes) (any, error)
	CastField(Field) Field
}
