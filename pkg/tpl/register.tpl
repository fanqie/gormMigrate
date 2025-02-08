package demo

import "github.com/fanqie/gormMigrate/pkg/impl"

var migrations = make(map[string]impl.GormMigrateInterface)

func RegisterMigration(name string, migrationFunc impl.GormMigrateInterface) {
	migrations[name] = migrationFunc
}
func register() {
    {{RegisterMigration}}
        //todo:最终替换为
	    RegisterMigration("{{@tag}}", NewMigrateV{{@tag}})
	    RegisterMigration("{{@tag}}", NewMigrateV{{@tag}})
	    RegisterMigration("{{@tag}}", NewMigrateV{{@tag}})
}
