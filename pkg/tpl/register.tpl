package migrations

import "github.com/fanqie/gormMigrate/pkg"

func Register(migrate *pkg.GormMigrate) {
    {{RegisterMigration}}
        //todo:最终替换为
	    // migrate.RegisterMigration("{{@tag0}}", NewMigrateV{{@tag0}})
	    // migrate.RegisterMigration("{{@tag1}}", NewMigrateV{{@tag1}})
	    // migrate.RegisterMigration("{{@tag2}}", NewMigrateV{{@tag2}})
}
