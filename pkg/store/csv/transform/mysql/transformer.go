package transform

import (
	"errors"

	"github.com/challenai/conveyer/pkg/codec"
	"github.com/challenai/conveyer/pkg/store/csv"
	"github.com/challenai/conveyer/pkg/store/mysql"
	"github.com/challenai/conveyer/pkg/transform"
)

type mysqlTransformer struct {
	sz     int
	source []transform.Field
	sink   []transform.Field
}

func NewMysqlTransformer(sourceFields []transform.Field) transform.Transformer {
	return &mysqlTransformer{
		sz:     len(sourceFields),
		source: sourceFields,
	}
}

func (t *mysqlTransformer) SourceFields() []transform.Field {
	return t.source
}

func (t *mysqlTransformer) SinkFields() []transform.Field {
	return t.sink
}

func (t *mysqlTransformer) SetSinkFields(sinkFields []transform.Field) {
	t.sink = sinkFields
}

func (t *mysqlTransformer) Castable() error {
	if len(t.source) != len(t.sink) {
		return errors.New("can't cast: source and sink have different fields length")
	}
	for i, v := range t.source {
		if err := t.CastableType(v, t.sink[i]); err != nil {
			return err
		}
	}
	return nil
}

func (t *mysqlTransformer) Cast(src []codec.Bytes) ([]any, error) {
	if len(src) != t.sz {
		return nil, errors.New("source types and source rows mismatch")
	}
	dest := make([]any, 0, t.sz)
	for i, v := range src {
		f, err := t.CastBytes(i, v)
		if err != nil {
			return nil, err
		}
		dest = append(dest, f)
	}
	return dest, nil
}

func (t *mysqlTransformer) CastFields() []transform.Field {
	dest := make([]transform.Field, 0, t.sz)
	for _, v := range t.source {
		t.sink = append(t.sink, t.CastField(v))
	}
	return dest
}

func (t *mysqlTransformer) CastableType(src transform.Field, dest transform.Field) error {
	// every mysql type could be turned to csv string
	return nil
}

func (t *mysqlTransformer) CastBytes(index int, src codec.Bytes) (any, error) {
	switch t.sink[index].Type {
	case csv.TypeString:
		switch t.source[index].Type {
		case mysql.TypeDecimal, mysql.TypeUnsignedTiny, mysql.TypeTiny, mysql.TypeSmallInt, mysql.TypeLongText, mysql.TypeLongBLOB, mysql.TypeFloat, mysql.TypeDouble, mysql.TypeNULL, mysql.TypeTimestamp, mysql.TypeInt24, mysql.TypeDate, mysql.TypeTime, mysql.TypeDateTime, mysql.TypeYear, mysql.TypeVarChar, mysql.TypeBit, mysql.TypeText, mysql.TypeBLOB, mysql.TypeEnum, mysql.TypeGeometry, mysql.TypeJSON, mysql.TypeUnsignedInt, mysql.TypeUnsignedBigInt, mysql.TypeBigInt, mysql.TypeMediumText, mysql.TypeMediumBLOB, mysql.TypeTinyText, mysql.TypeTinyBLOB, mysql.TypeSET, mysql.TypeChar, mysql.TypeBinary, mysql.TypeVarBinary:
			return string(src), nil
		}
	}
	return nil, errors.New("unknown source or sink type")
}

func (t *mysqlTransformer) CastField(src transform.Field) transform.Field {
	return transform.Field{
		Name: src.Name,
		Type: csv.TypeString,
	}
}
