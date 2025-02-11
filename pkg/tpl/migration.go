package tpl

const MigrationCreateTableCode = `package gorm_migrations

import (
	migrate "github.com/fanqie/gormMigrate/pkg/core"
)
type {{TableStructName}} struct {
	Id        int64 ` + "`" + `gorm:"primaryKey;autoIncrement"` + "`" + `
	CreatedAt int64 ` + "`" + `gorm:"autoCreateTime"` + "`" + `
	UpdatedAt int64 ` + "`" + `gorm:"autoUpdateTime"` + "`" + `
	DeletedAt int64 ` + "`" + `gorm:"index"` + "`" + `
}
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
	err := migrate.Migrator().CreateTable(&{{TableStructName}}{})
		if err != nil {
			panic(err)
		}
}
func (r *Migrate{{TypeTag}}) Down() {
	err := migrate.Migrator().DropTable(&{{TableStructName}}{})
		if err != nil {
			panic(err)
		}
}
`
const MigrationAlterTableCode = `package gorm_migrations

import (
	migrate "github.com/fanqie/gormMigrate/pkg/core"
)
type {{TableStructName}} struct {
	Id        int64 ` + "`" + `gorm:"primaryKey;autoIncrement"` + "`" + `
	CreatedAt int64 ` + "`" + `gorm:"autoCreateTime"` + "`" + `
	UpdatedAt int64 ` + "`" + `gorm:"autoUpdateTime"` + "`" + `
	DeletedAt int64 ` + "`" + `gorm:"index"` + "`" + `
}
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
	
	err := migrate.Migrator().CreateTable(&{{TableStructName}}{})
		if err != nil {
			panic(err)
		}
}
func (r *Migrate{{TypeTag}}) Down() {
	err := migrate.Migrator().DropTable(&{{TableStructName}}{})
		if err != nil {
			panic(err)
		}
}
`
