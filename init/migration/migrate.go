package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"log"
	"micro-k8s-client/global"
	"micro-k8s-client/init/migration/migrations"
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
