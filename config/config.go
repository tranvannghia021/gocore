package config

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/vars"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func ConnectDB() {
	host, _ := os.LookupEnv("DB_CORE_HOST")
	port, _ := os.LookupEnv("DB_CORE_PORT")
	user, _ := os.LookupEnv("DB_CORE_USER")
	pass, _ := os.LookupEnv("DB_CORE_PASSWORD")
	db, _ := os.LookupEnv("DB_CORE_NAME")
	connect, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, db, port)))
	helpers.CheckNilErr(err)
	vars.Connection = connect
	if !helpers.CheckTable(connect) {
		helpers.MigrateCore(connect)
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
	vars.Redis = client
	fmt.Println("-----------------redis CORE ready " + pong)
}
func init() {
	godotenv.Load()
	ConnectDB()
	ConnectCache()
}
