package demo

import (
	"fmt"
	migrate "github.com/fanqie/gormMigrate/pkg/migrate"
)

type MigrateV001 struct {
	migrate.Basic
}

func NewMigrateV001() *MigrateV001 {
	return &MigrateV001{}
}
func (r *MigrateV001) Register() {
	r.Tag = "V001"

}
func (r *MigrateV001) Up() {
	// todo:migrate the database
	fmt.Printf("test,%v", r.Tag)
	r.UpAfter()
}
func (r *MigrateV001) Down() {
	// todo:migrate the database
	fmt.Printf("test,%v", r.Tag)
	r.DownAfter()
}
