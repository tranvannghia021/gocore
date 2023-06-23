package config

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

var Connection *gorm.DB
var Redis *redis.Client

func ConnectDB() {
	host, _ := os.LookupEnv("DB_CORE_HOST")
	port, _ := os.LookupEnv("DB_CORE_PORT")
	user, _ := os.LookupEnv("DB_CORE_USER")
	pass, _ := os.LookupEnv("DB_CORE_PASSWORD")
	db, _ := os.LookupEnv("DB_CORE_NAME")
	connect, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, db, port)))
	helpers.CheckNilErr(err)
	Connection = connect
	if !CheckTable(repositories.Core{}, connect) {
		MigrateCore(connect)
	}
	fmt.Println("------------postgres CORE ready")

}

func ConnectCache() {
	host_redis, _ := os.LookupEnv("CORE_REDIS_HOST")
	port_redis, _ := os.LookupEnv("CORE_REDIS_PORT")
	pass_redis, _ := os.LookupEnv("CORE_REDIS_PASSWORD")
	db_redis, _ := os.LookupEnv("CORE_REDIS_DB")
	db, _ := strconv.Atoi(db_redis)
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host_redis, port_redis),
		Password: pass_redis,
		DB:       db,
	})
	pong, err := client.Ping().Result()
	helpers.CheckNilErr(err)
	Redis = client
	fmt.Println("-----------------redis CORE ready " + pong)
}

type ConfigCore struct {
	Database struct {
		TableName string            `json:"table_name"`
		Fields    map[string]string `json:"fields"`
	} `json:"database"`
}

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
