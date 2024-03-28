package global

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	configs "micro-k8s-client/conf"
)

var (
	DB    *gorm.DB
	LOG   *logrus.Logger
	CONF  configs.ServerConfig
	VALID *validator.Validate
	//SESSION *psession.PSession
	//CACHE   *badger_db.Cache
	Viper *viper.Viper

	//Cron          *cron.Cron
	MonitorCronID int
)
