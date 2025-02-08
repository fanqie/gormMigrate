package migrations

import (
	"fmt"
	migrate "github.com/fanqie/gormMigrate/pkg/migrate"
)

type Migrate{{Tag}} struct {
	migrate.Basic
}

func NewMigrate{{Tag}}() *Migrate{{Tag}} {
	return &Migrate{{Tag}}{}
}
func (r *Migrate{{Tag}}) Register() {
	r.Tag = "{{Tag}}"

}
func (r *Migrate{{Tag}}) Up() {
	// todo:migrate the database
	fmt.Printf("test,%v", r.Tag)
	r.UpAfter()
}
func (r *Migrate{{Tag}}) Down() {
	// todo:migrate the database
	fmt.Printf("test,%v", r.Tag)
	r.DownAfter()
}
