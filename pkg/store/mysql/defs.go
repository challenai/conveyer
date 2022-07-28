package mysql

import "github.com/challenai/conveyer/pkg/transform"

// https://dev.mysql.com/doc/internals/en/com-query-response.html#packet-Protocol::ColumnType

const (
	TypeDecimal          = transform.Type("DECIMAL")
	TypeUnsignedTiny     = transform.Type("UNSIGNED TINYINT")
	TypeTiny             = transform.Type("TINYINT")
	TypeUnsignedSmallInt = transform.Type("UNSIGNED SMALLINT")
	TypeSmallInt         = transform.Type("SMALLINT")
	TypeLongText         = transform.Type("LONGTEXT")
	TypeLongBLOB         = transform.Type("LONGBLOB")
	TypeFloat            = transform.Type("FLOAT")
	TypeDouble           = transform.Type("DOUBLE")
	TypeNULL             = transform.Type("NULL")
	TypeTimestamp        = transform.Type("TIMESTAMP")
	TypeInt24            = transform.Type("MEDIUMINT")
	TypeDate             = transform.Type("DATE")
	TypeTime             = transform.Type("TIME")
	TypeDateTime         = transform.Type("DATETIME")
	TypeYear             = transform.Type("YEAR")
	TypeVarChar          = transform.Type("VARCHAR")
	TypeBit              = transform.Type("BIT")
	TypeText             = transform.Type("TEXT")
	TypeBLOB             = transform.Type("BOLB")
	TypeEnum             = transform.Type("ENUM")
	TypeGeometry         = transform.Type("GEOMETRY")
	TypeJSON             = transform.Type("JSON")
	TypeUnsignedInt      = transform.Type("UNSIGNED INT")
	TypeUnsignedBigInt   = transform.Type("UNSIGNED BIGINT")
	TypeBigInt           = transform.Type("BIGINT")
	TypeMediumText       = transform.Type("MEDIUMTEXT")
	TypeMediumBLOB       = transform.Type("MEDIUMBLOB")
	TypeTinyText         = transform.Type("TINYTEXT")
	TypeTinyBLOB         = transform.Type("TINYBLOB")
	TypeSET              = transform.Type("SET")
	TypeChar             = transform.Type("CHAR")
	TypeBinary           = transform.Type("BINARY")
	TypeVarBinary        = transform.Type("VARBINARY")
	TypeUnknown          = transform.Type("UNKNOWN")
)
