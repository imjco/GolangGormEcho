package db

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DatabaseConnection struct {
	Server   string
	UserID   string
	Password string
	Database string
	Alias    string
}

func (Db *DatabaseConnection) Open() gorm.Dialector {
	return sqlserver.Open(fmt.Sprintf("odbc:server=%s;user id=%s;password={%s}; database=%s", Db.Server, Db.UserID, Db.Password, Db.Database))
}
