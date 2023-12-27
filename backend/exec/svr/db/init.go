package db

import (
	"github.com/KevinZonda/float32/exec/svr/dbmodel"
	"gorm.io/gorm"
	"log"
)
import "gorm.io/driver/mysql"

var _db *gorm.DB

func InitDb(sqlUrl string) {
	var err error
	_db, err = gorm.Open(
		mysql.Open(sqlUrl),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			PrepareStmt:                              true,
		})
	if err != nil {
		_db = nil
		log.Println("Failed to connect to database", err)
		return
	}
	_db.AutoMigrate(
		&dbmodel.Answer{},
	)
}
