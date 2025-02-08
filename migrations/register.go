package migrations

import (
	"fmt"
	"github.com/fanqie/gormMigrate/pkg"
)

func Register(migrate *pkg.GormMigrate) {
	fmt.Println("register migrate beging")
	migrate.RegisterMigration("migrate_v001", NewMigrateV001())
	migrate.RegisterMigration("migrate_v002", NewMigrateV002())
	fmt.Println("register migrate end")
}
