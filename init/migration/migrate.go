package migration

import (
	"gin-dubbogo-consumer/global"
	"gin-dubbogo-consumer/init/migration/migrations"
	"github.com/go-gormigrate/gormigrate/v2"
	"log"
)

func Init() {
	m := gormigrate.New(global.DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		migrations.AddTableHost,
	})
	if err := m.Migrate(); err != nil {
		global.LOG.Error(err)
		panic(err)
	}
	log.Println("Migration run successfully")
}
