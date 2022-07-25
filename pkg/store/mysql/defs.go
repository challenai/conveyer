package mysql

import "github.com/challenai/conveyer/pkg/transform"

// https://dev.mysql.com/doc/internals/en/com-query-response.html#packet-Protocol::ColumnType

const (
	MysqlTypeDecimal transform.Type = iota
	MysqlTypeTiny
	MysqlTypeShort
	MysqlTypeLong
	MysqlTypeFloat
	MysqlTypeDouble
	MysqlTypeNULL
	MysqlTypeTimestamp
	MysqlTypeLongLong
	MysqlTypeInt24
	MysqlTypeDate
	MysqlTypeTime
	MysqlTypeDateTime
	MysqlTypeYear
	MysqlTypeNewDate
	MysqlTypeVarChar
	MysqlTypeBit
)
const (
	MysqlTypeJSON transform.Type = iota + 0xf5
	MysqlTypeNewDecimal
	MysqlTypeEnum
	MysqlTypeSet
	MysqlTypeTinyBLOB
	MysqlTypeMediumBLOB
	MysqlTypeLongBLOB
	MysqlTypeBLOB
	MysqlTypeVarString
	MysqlTypeString
	MysqlTypeGeometry
)
