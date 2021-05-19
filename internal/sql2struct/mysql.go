package sql2struct

import (
	"database/sql"
	"fmt"
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?" +
		"charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(s, m.DBInfo.UserName, m.DBInfo.Password, m.DBInfo.Host, m.DBInfo.Charset)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}
