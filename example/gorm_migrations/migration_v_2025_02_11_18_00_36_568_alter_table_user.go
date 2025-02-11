package gorm_migrations

import (
	migrate "github.com/fanqie/gormMigrate/pkg/core"
	"gorm.io/gorm"
)

type MigrateV20250211180036568AlterTableUser struct {
	migrate.MigrateBasic
	currentTable *StructV20250211180036568AlterTableUser
}
type StructV20250211180036568AlterTableUser struct {
	UserName string `gorm:"index"`
}

func (*StructV20250211180036568AlterTableUser) TableName() string {
	return "user"
}

func NewMigrateV20250211180036568AlterTableUser() *MigrateV20250211180036568AlterTableUser {
	return &MigrateV20250211180036568AlterTableUser{
		currentTable: &StructV20250211180036568AlterTableUser{},
	}
}
func (r *MigrateV20250211180036568AlterTableUser) Register() {
	r.Tag = "v_2025_02_11_18_00_36_568_alter_table_user"

}
func (r *MigrateV20250211180036568AlterTableUser) Up(tx *gorm.DB) error {
	err := tx.Migrator().AddColumn(r.currentTable, "UserName")
	if err != nil {
		return err
	}
	return nil
}
func (r *MigrateV20250211180036568AlterTableUser) Down(tx *gorm.DB) error {
	err := tx.Migrator().DropColumn(r.currentTable, "UserName")
	if err != nil {
		return err
	}
	return nil
}
