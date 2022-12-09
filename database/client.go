package database

import (
	"time"

	"github.com/go-bai/forward/model"
	"github.com/spf13/viper"
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func InitEngine() {
	var err error
	Engine, err = xorm.NewEngine(viper.GetString("db.driver"), viper.GetString("db.dsn"))
	if err != nil {
		panic(err)
	}
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	Engine.TZLocation = location
	Engine.DatabaseTZ = location
	Engine.ShowSQL(true)
	Engine.SetMaxIdleConns(viper.GetInt("db.max_idle_conns"))
	Engine.SetMaxOpenConns(viper.GetInt("db.max_open_conns"))
	connMaxLifetime, err := time.ParseDuration(viper.GetString("db.conn_max_lifetime"))
	if err != nil {
		panic(err)
	}
	Engine.SetConnMaxLifetime(connMaxLifetime)
}

func InitTable() {
	_, err := Engine.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}

	user := &model.User{Username: "admin", Password: "admin"}
	user.HashPassword()
	Engine.InsertOne(user)
}
