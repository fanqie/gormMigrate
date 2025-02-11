package tpl

const MigrationCreateTableCode = `package gorm_migrations

import (
	migrate "github.com/fanqie/gormMigrate/pkg/core"
	"gorm.io/gorm"
)

type Migrate{{TypeTag}} struct {
	migrate.MigrateBasic
	currentTable *Struct{{TypeTag}}
}
type Struct{{TypeTag}} struct{
	Id        int64 ` + "`" + `gorm:"primaryKey;autoIncrement"` + "`" + `
	CreatedAt int64 ` + "`" + `gorm:"autoCreateTime"` + "`" + `
	UpdatedAt int64 ` + "`" + `gorm:"autoUpdateTime"` + "`" + `
	DeletedAt int64 ` + "`" + `gorm:"index"` + "`" + `
}
func (*Struct{{TypeTag}}) TableName() string {
	return "{{TableName}}"
}
func NewMigrate{{TypeTag}}() *Migrate{{TypeTag}} {
	return &Migrate{{TypeTag}}{
		currentTable:&Struct{{TypeTag}}{},
	}
}
func (r *Migrate{{TypeTag}}) Register() {
	r.Tag = "{{Tag}}"

}
func (r *Migrate{{TypeTag}}) Up(tx *gorm.DB) error{
	
	err := tx.Migrator().CreateTable(r.currentTable)
	if err != nil {
		return err
	}
	return nil
}
func (r *Migrate{{TypeTag}}) Down(tx *gorm.DB) error{
	err := tx.Migrator().DropTable(r.currentTable)
	if err != nil {
		return err
	}
	return nil
}
`
const MigrationAlterTableCode = `package gorm_migrations

import (
	migrate "github.com/fanqie/gormMigrate/pkg/core"
	"gorm.io/gorm"
)

type Migrate{{TypeTag}} struct {
	migrate.MigrateBasic
	currentTable *Struct{{TypeTag}}
}
type Struct{{TypeTag}} struct{
	UserName        string ` + "`" + `gorm:"index"` + "`" + `
}
func (*Struct{{TypeTag}}) TableName() string {
	return "{{TableName}}"
}

func NewMigrate{{TypeTag}}() *Migrate{{TypeTag}} {
	return &Migrate{{TypeTag}}{
		currentTable:&Struct{{TypeTag}}{},
	}
}
func (r *Migrate{{TypeTag}}) Register() {
	r.Tag = "{{Tag}}"

}
func (r *Migrate{{TypeTag}}) Up(tx *gorm.DB) error{
	err := tx.Migrator().AddColumn(r.currentTable,"UserName")
	if err != nil {
		return err
	}
	return nil
}
func (r *Migrate{{TypeTag}}) Down(tx *gorm.DB) error{
	err := tx.Migrator().DropColumn(r.currentTable,"UserName")
	if err != nil {
		return err
	}
	return nil
}
`
