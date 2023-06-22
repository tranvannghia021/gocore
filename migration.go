package gocore

import "github.com/tranvannghia021/gocore/src"

func MigrateCore(db *gorm.DB) {
	er := db.Migrator().CreateTable(&src.Core{})
	helpers.CheckNilErr(er)
}

func RollbackMigrate(db *gorm.DB) {
	er := db.Migrator().DropTable(&src.Core{})
	helpers.CheckNilErr(er)
}

func CheckTable(table src.Core, db *gorm.DB) bool {
	return db.Migrator().HasTable(&table)
}
