package main

import (
	"github.com/fanqie/gormMigrate-example/gorm_migrations"
	gm "github.com/fanqie/gormMigrate/pkg"
	"github.com/fanqie/gormMigrate/pkg/storage"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	gormMigrate := gm.NewGormMigrate()

	gormMigrate.Setup(storage.GromParams{
		Dialector: mysqlDialector(),
		Opts: &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		},
	}, func() {
		gorm_migrations.Register(gormMigrate)
	})
}
func mysqlDialector() gorm.Dialector {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	return mysql.Open(dsn)
}
