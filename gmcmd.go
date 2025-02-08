package main

import (
	"github.com/fanqie/gormMigrate/migrations"
	gm "github.com/fanqie/gormMigrate/pkg"
	"github.com/fanqie/gormMigrate/pkg/storage"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	gormMigrate := gm.NewGormMigrate()
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	gormMigrate.Setup(storage.GromParams{
		Dialector: mysql.Open(dsn),
		Opts: &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	}, func() {
		migrations.Register(gormMigrate)
	})

}
