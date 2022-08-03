package source

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/challenai/conveyer/pkg/codec"
	"github.com/challenai/conveyer/pkg/graph/desc"
	"github.com/challenai/conveyer/pkg/source"
	mysqltable "github.com/challenai/conveyer/pkg/store/mysql/table"
	"github.com/challenai/conveyer/pkg/table"
	_ "github.com/go-sql-driver/mysql"
)

const (
	KindMysql = "mysql"

	DefaultHost = "localhost"
	DefaultPort = 3306
	DefaultUser = "root"
	// 3 minutes
	DefaultConnMaxLifetime = 60 * 3
	DefaultMaxOpenConns    = 3
	DefaultMaxIdleConns    = 3

	DefaultCharset   = "utf8mb4"
	DefaultParseTime = "True"
	DefaultLoc       = "Local"

	KeywordLimit  = "LIMIT"
	KeywordOffset = "Offset"
)

type mysqlSource struct {
	db *sql.DB
	table.TableManager
	desc  desc.Source
	extra *MysqlSourceExtraDescription
}

type MysqlSourceExtraDescription struct {

	// basic connection information
	Database string
	Host     string
	Port     int
	User     string
	Passwd   string

	// table information
	Charset   string
	ParseTime string
	Loc       string

	// connections setting
	ConnMaxLifetime int
	MaxOpenConns    int
	MaxIdleConns    int
}

func NewMysqlSource(desc desc.Source) (source.Source, error) {
	extra, ok := desc.Extra.(MysqlSourceExtraDescription)
	if !ok {
		return nil, errors.New("bad source extra description")
	}

	err := extra.validate()
	if err != nil {
		return nil, err
	}

	s := &mysqlSource{
		desc:  desc,
		extra: &extra,
	}

	s.extra.setDefault()

	err = validateDSL(desc.DSL)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (ms *mysqlSource) Open() error {
	var err error

	ms.db, err = sql.Open(KindMysql, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s", ms.extra.User, ms.extra.Passwd, ms.extra.Host, ms.extra.Port, ms.extra.Database, ms.extra.Charset, ms.extra.ParseTime, ms.extra.Loc))
	if err != nil {
		return err
	}

	ms.TableManager = mysqltable.NewMySQLTableManager(ms.db)

	return nil
}

func (ms *mysqlSource) Count(queryDSL string) (int, error) {
	// rows, err := ms.db.Query(fmt.Sprintf("COUNT(%s)", queryDSL))
	// if err != nil {
	// 	return nil, err
	// }

	return 0, nil
}

func (ms *mysqlSource) Query(queryDSL string, offset, limit int) ([][]codec.Bytes, error) {
	fmt.Println(fmt.Sprintf("%s LIMIT %d OFFSET %d", queryDSL, limit, offset))
	rows, err := ms.db.Query(fmt.Sprintf("%s LIMIT %d OFFSET %d", queryDSL, limit, offset))
	if err != nil {
		return nil, err
	}

	var records [][]codec.Bytes
	r := make([]any, ms.TableManager.GetFieldsCount())
	for rows.Next() {
		for i := range r {
			r[i] = new(sql.RawBytes)
		}

		err = rows.Scan(r...)
		if err != nil {
			return nil, err
		}

		record := make([]codec.Bytes, ms.TableManager.GetFieldsCount())
		for i, v := range r {
			ptr, _ := v.(*sql.RawBytes)
			record[i] = (codec.Bytes)(*ptr)
		}

		records = append(records, record)
	}

	return records, nil
}

func (ms *mysqlSource) Close() error {
	return nil
}

func (desc *MysqlSourceExtraDescription) validate() error {
	if desc.Database == "" {
		return errors.New("config error: mysql database can't be empty")
	}
	if desc.User == "" {
		desc.User = DefaultUser
	}

	return nil
}

func validateDSL(queryDSL string) error {
	upperDSL := strings.ToUpper(queryDSL)
	if strings.Contains(upperDSL, KeywordLimit) {
		return errors.New("config error: mysql query dsl include keywords LIMIT")
	}
	if strings.Contains(upperDSL, KeywordOffset) {
		return errors.New("config error: mysql query dsl include keywords OFFSET")
	}

	return nil
}

func (desc *MysqlSourceExtraDescription) setDefault() {
	if desc.Host == "" {
		desc.Host = DefaultHost
	}
	if desc.Port <= 0 {
		desc.Port = DefaultPort
	}

	if desc.ConnMaxLifetime <= 0 {
		desc.ConnMaxLifetime = DefaultConnMaxLifetime
	}
	if desc.MaxOpenConns <= 0 {
		desc.MaxOpenConns = DefaultMaxOpenConns
	}
	if desc.MaxIdleConns <= 0 {
		desc.MaxIdleConns = DefaultMaxIdleConns
	}

	if desc.Charset == "" {
		desc.Charset = DefaultCharset
	}
	if desc.ParseTime == "" {
		desc.ParseTime = DefaultParseTime
	}
	if desc.Loc == "" {
		desc.Loc = DefaultLoc
	}
}
