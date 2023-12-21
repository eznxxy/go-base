package db

import (
	"fmt"

	"github.com/eznxxy/go-base/config"
	"github.com/eznxxy/go-base/db/seeders"

	_ "github.com/go-sql-driver/mysql" // nolint
	"github.com/jinzhu/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name)

	db, err := gorm.Open(cfg.DB.Driver, dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	userSeeder := seeders.NewUserSeeder(db)
	userSeeder.SetUsers()

	return db
}
