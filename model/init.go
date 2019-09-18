package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Self:   getSelfDB(),
		Docker: getDockerDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
	DB.Docker.Close()
}

func getDockerDB() *gorm.DB {
	return openDB(
		viper.GetString("docker_db.name"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.user"),
		viper.GetString("docker_db.passwd"),
	)
}

func getSelfDB() *gorm.DB {
	return openDB(
		viper.GetString("db.name"),
		viper.GetString("db.addr"),
		viper.GetString("db.user"),
		viper.GetString("db.passwd"),
	)
}

func openDB(name, addr, user, passwd string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		user,
		passwd,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "open database[%s] error", name)
		return nil
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(0)
	return db
}
