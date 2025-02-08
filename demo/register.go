package demo

import "github.com/fanqie/gormMigrate/pkg/impl"

var migrations = make(map[string]impl.GormMigrateInterface)

func RegisterMigration(name string, migrationFunc impl.GormMigrateInterface) {
	migrations[name] = migrationFunc
}
func register() {
	RegisterMigration("migrate_v001", NewMigrateV001())
	RegisterMigration("migrate_v002", NewMigrateV002())
}
