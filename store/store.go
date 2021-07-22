package store

import (
	"fmt"
	"log"

	"github.com/danishk121/go-frame/store/adaptee/mysql"
	"github.com/danishk121/go-frame/store/adapter"
	"github.com/joho/godotenv"
	sql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Store adapter.Store
)

func Initialize() {
	var appConfig map[string]string
	appConfig, confErr := godotenv.Read()

	if confErr != nil {
		log.Fatal("Error reading .env file")
		return
	}

	// Ex: user:password@tcp(host:port)/dbname
	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		appConfig["MYSQL_USER"],
		appConfig["MYSQL_PASSWORD"],
		appConfig["MYSQL_PROTOCOL"],
		appConfig["MYSQL_HOST"],
		appConfig["MYSQL_PORT"],
		appConfig["MYSQL_DBNAME"],
	)

	db, confErr := gorm.Open(sql.Open(mysqlCredentials), &gorm.Config{})

	if confErr != nil {
		log.Printf("Error starting server: %s\n", confErr)
		return
	}
	Init(db)

}

func Init(db *gorm.DB) {
	Store = mysql.NewAdapter(db)
}
