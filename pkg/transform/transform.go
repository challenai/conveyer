package transform

type Type string

type Field struct {
	Name string
	Type
}

type Transformer interface {
	// notify
	SetSourceFields([]Field)
	SetSinkFields([]Field)

	Castable() (bool, error)

	CastFields() []Field

	Cast(src any) (any, error)

	// 	CastableType(src Field, dest Field) error
	// 	CastBytes(index int, src codec.Bytes) (any, error)
	// 	CastField(Field) Field
}
