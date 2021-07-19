// database/database.go

package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GlobalDB a global db object will be used across different packages
var GlobalDB *gorm.DB
var GlobalDBAcc *gorm.DB
var GlobalDBTrans *gorm.DB

// InitDatabase creates a sqlite db
func InitDatabase(dsn string) (err error) {
	GlobalDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	return
}

func InitDatabaseAcc(dsn string) (err error) {
	GlobalDBAcc, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return
}

func InitDatabaseTrans(dsn string) (err error) {
	GlobalDBTrans, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return
}
