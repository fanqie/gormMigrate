package migrations

import (
	"fmt"
	migrate "github.com/fanqie/gormMigrate/pkg/migrate"
)

type Migrate{{TypeTag}} struct {
	migrate.MigrateBasic
}

func NewMigrate{{TypeTag}}() *Migrate{{TypeTag}} {
	return &Migrate{{TypeTag}}{}
}
func (r *Migrate{{TypeTag}}) Register() {
	r.Tag = "{{Tag}}"

}
func (r *Migrate{{TypeTag}}) Up() {
	// todo:migrate the database
	fmt.Printf("test,%v", r.Tag)
	r.UpAfter()
}
func (r *Migrate{{TypeTag}}) Down() {
	// todo:migrate the database
	fmt.Printf("test,%v", r.Tag)
	r.DownAfter()
}
