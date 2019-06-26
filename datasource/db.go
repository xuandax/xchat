package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"sync"
)

var engine *xorm.Engine
var mux sync.Mutex

func InstanceDb() *xorm.Engine {
	if engine != nil {
		return engine
	}
	mux.Lock()
	defer mux.Unlock()
	if engine != nil {
		return engine
	}
	return newInstanceDb()
}

func newInstanceDb() *xorm.Engine {
	driverName := "mysql"
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		"root",
		"123456",
		"120.79.167.42",
		3306,
		"xchat")
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	return engine
}
