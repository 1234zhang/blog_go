package models
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"project/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID 				int `gorm:"primary_key" json:"id"`
	CreatedOn 		int `json:"created_on"`
	ModifiedOn		int `json:"modified_on"`
}

func init() {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	database, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("get database section happen some error: %v", err)
	}

	dbType 		= database.Key("TYPE").String()
	dbName 		= database.Key("NAME").String()
	user   		= database.Key("USER").String()
	password 	= database.Key("PASSWORD").String()
	host		= database.Key("HOST").String()
	tablePrefix = database.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, dbName))
	if err != nil {
		log.Fatalf("connect database happen some error: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

