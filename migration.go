package gocore

import (
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/repositories"
	"gorm.io/gorm"
)

func MigrateCore(db *gorm.DB) {
	er := db.Migrator().CreateTable(&repositories.Core{})
	helpers.CheckNilErr(er)
}

func RollbackMigrate(db *gorm.DB) {
	er := db.Migrator().DropTable(&repositories.Core{})
	helpers.CheckNilErr(er)
}

func CheckTable(table repositories.Core, db *gorm.DB) bool {
	return db.Migrator().HasTable(&table)
}
