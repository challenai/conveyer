package avro

type AvroType int32

const (
	TypeNULL AvroType = iota + 1
	TypeBool
	TypeBytes
	TypeFloat
	TypeDouble
	TypeLong
	TypeInt
	TypeString
	TypeArray
	TypeEnum
	TypeFixed
	TypeMap
	TypeRecord
	TypeUnion
)

func (t AvroType) AvroType() string {
	switch t {
	case TypeNULL:
		return "null"
	case TypeBool:
		return "bool"
	case TypeBytes:
		return "bytes"
	case TypeFloat:
		return "float"
	case TypeDouble:
		return "double"
	case TypeLong:
		return "long"
	case TypeInt:
		return "int"
	case TypeString:
		return "string"
	case TypeArray:
		return "array"
	case TypeEnum:
		return "enum"
	case TypeFixed:
		return "fixed"
	case TypeMap:
		return "map"
	case TypeRecord:
		return "record"
	case TypeUnion:
		return "union"
	}
	return ""
}
