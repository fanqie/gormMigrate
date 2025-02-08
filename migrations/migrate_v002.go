package migrations

import (
	"fmt"
	migrate "github.com/fanqie/gormMigrate/pkg/migrate"
)

type MigrateV002 struct {
	migrate.Migrate
}

func NewMigrateV002() *MigrateV002 {
	return &MigrateV002{}
}
func (r *MigrateV002) Register() {
	r.Tag = "V002"

}
func (r *MigrateV002) Up() {
	// todo:migrate the database
	fmt.Printf("test,%v", r.Tag)
	r.UpAfter()
}
func (r *MigrateV002) Down() {
	// todo:migrate the database
	fmt.Printf("test,%v", r.Tag)
	r.DownAfter()
}
