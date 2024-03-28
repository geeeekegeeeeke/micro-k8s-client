package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"micro-k8s-client/model"
)

var AddTableHost = &gormigrate.Migration{
	ID: "20231102",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&model.K8sCluster{}); err != nil {
			return err
		}
		if err := tx.AutoMigrate(&model.Group{}); err != nil {
			return err
		}
		if err := tx.AutoMigrate(&model.Command{}); err != nil {
			return err
		}
		group := model.Group{
			Name: "default", Type: "host", IsDefault: true,
		}
		if err := tx.Create(&group).Error; err != nil {
			return err
		}
		/*host := model.Host{
			Name: "localhost", Addr: "127.0.0.1", User: "root", Port: 22, AuthMode: "password", GroupID: group.ID,
		}
		if err := tx.Create(&host).Error; err != nil {
			return err
		}*/
		return nil
	},
}
