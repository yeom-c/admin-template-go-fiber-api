package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/yeom-c/admin-template-go-fiber-api/app"
	"log"
	"sync"
	"xorm.io/xorm"
)

type database struct {
	Conn *xorm.Engine
}

var once sync.Once
var instance *database

func Database() *database {
	once.Do(func() {
		if instance == nil {
			instance = &database{}
			engine, err := xorm.NewEngine(app.Config().DbDriver, app.Config().DbConn)
			if err != nil {
				log.Fatal("failed to connect database: ", err)
			}

			if app.Config().SqlShow {
				engine.ShowSQL(true)
			}
			instance.Conn = engine
		}
	})

	return instance
}
