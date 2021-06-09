package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

type DBConfig struct {
	User string
	Pass string
	Name string
	Host string
	Port string
}

func getConfig() DBConfig {
	conf := DBConfig{
		User: "test",
		Pass: "test",
		Name: "test",
		Host: "localhost",
		Port: "3307",
	}
	return conf
}

func newDB() *xorm.Engine {
	config := getConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=UTC",
		config.User, config.Pass, config.Host, config.Port, config.Name)

	db, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	logger := xlog.NewSimpleLogger(os.Stdout)
	db.SetLogger(logger)
	db.ShowSQL(true)
	db.SetLogLevel(xlog.LOG_DEBUG)

	//db.SetTZDatabase(time.UTC)
	//db.SetTZLocation(time.UTC)

	return db
}

type User struct {
	Created time.Time `xorm:"created"`
}

func main() {
	db := newDB()
	_, err := db.Exec("truncate table user")
	if err != nil {
		log.Fatal(err)
	}

	date := time.Now()
	fmt.Println(date)
	if _, err := db.Insert(User{Created: date}); err != nil {
		log.Fatal(err)
	}

	users := []User{}
	if err := db.Find(&users); err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
}
