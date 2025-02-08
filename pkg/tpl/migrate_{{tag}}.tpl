package migrate

import (
	gormMigrate "github.com/fanqie/gormMigrate"
)
func Register() {
    migrate := gormMigrate.NewMigrate()
    migrate.Tag = "{{Tag}}"
    migrate.Up = func Up() {
                     // todo:migrate the database

                 }
    migrate.Up = func Down() {
                    // todo:migrate the database

                  }
}
