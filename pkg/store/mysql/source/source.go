package source

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/challenai/conveyer/pkg/codec"
	"github.com/challenai/conveyer/pkg/graph/desc"
	"github.com/challenai/conveyer/pkg/source"
	mysqltable "github.com/challenai/conveyer/pkg/store/mysql/table"
	"github.com/challenai/conveyer/pkg/table"
)

type mysqlSource struct {
	db *sql.DB
	table.TableManager
	desc MysqlSourceDescription
}

const (
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
)

type MysqlSourceDescription struct {
	desc.Source

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

func (desc *MysqlSourceDescription) validate() error {
	if desc.Host == "" {
		desc.Host = DefaultHost
	}
	if desc.Port <= 0 {
		desc.Port = DefaultPort
	}
	if desc.Database == "" {
		return errors.New("config error: mysql database can't be empty")
	}
	if desc.User == "" {
		desc.User = DefaultUser
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

	if desc.ConnMaxLifetime <= 0 {
		desc.ConnMaxLifetime = DefaultConnMaxLifetime
	}
	if desc.MaxOpenConns <= 0 {
		desc.MaxOpenConns = DefaultMaxOpenConns
	}
	if desc.MaxIdleConns <= 0 {
		desc.MaxIdleConns = DefaultMaxIdleConns
	}
	return nil
}

func NewMysqlSource(desc MysqlSourceDescription) (source.Source, error) {
	err := desc.validate()
	if err != nil {
		return nil, err
	}

	return &mysqlSource{
		desc: desc,
	}, nil
}

func (ms *mysqlSource) Open() error {
	var err error

	ms.db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s", ms.desc.User, ms.desc.Passwd, ms.desc.Host, ms.desc.Port, ms.desc.Database, ms.desc.Charset, ms.desc.ParseTime, ms.desc.Loc))
	if err != nil {
		return err
	}

	ms.TableManager = mysqltable.NewMySQLTableManager(ms.db, ms.desc.Table)

	return nil
}

func (ms *mysqlSource) Count(query string) (int, error) {
	return 0, nil
}

func (ms *mysqlSource) Query(query string, offset, limit int) ([][]codec.Bytes, error) {
	return nil, nil
}

func (ms *mysqlSource) Close() error {
	return nil
}
