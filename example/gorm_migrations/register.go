package gorm_migrations

import (
	"github.com/fanqie/gormMigrate/pkg"
)

func Register(migrate *pkg.GormMigrate) {
	migrate.RegisterMigration("V20250211175923254CreateTableUser", NewMigrateV20250211175923254CreateTableUser())
	migrate.RegisterMigration("V20250211180036568AlterTableUser", NewMigrateV20250211180036568AlterTableUser())
}
