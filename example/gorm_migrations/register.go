package gorm_migrations

import (
	"github.com/fanqie/gormMigrate/pkg"
)

func Register(migrate *pkg.GormMigrate) {
	migrate.RegisterMigration("V20250211115311401CreateTableAdmin", NewMigrateV20250211115311401CreateTableAdmin())
	migrate.RegisterMigration("V20250211143212878CreateTableAdmin", NewMigrateV20250211143212878CreateTableAdmin())
}
