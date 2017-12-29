package main

import (
	"fmt"

	"./config"
	"./entities"
	"github.com/go-xorm/xorm"
)
import _ "github.com/go-sql-driver/mysql"
import _ "github.com/mattn/go-sqlite3"

func main() {
	engine, err := xorm.NewEngine(config.Config.DriverName, config.Config.DataSourceName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer engine.Close()
	engine.ShowSQL(true)
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)
	err = engine.CreateTables(&entities.User{})
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = engine.Insert(&entities.User{Name: "xlw"})
	if err != nil {
		fmt.Println(err)
		return
	}

}
