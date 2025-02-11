package gorm_migrations

import (
	migrate "github.com/fanqie/gormMigrate/pkg/core"
	"gorm.io/gorm"
)

type MigrateV20250211175923254CreateTableUser struct {
	migrate.MigrateBasic
	currentTable *StructV20250211175923254CreateTableUser
}
type StructV20250211175923254CreateTableUser struct {
	Id        int64 `gorm:"primaryKey;autoIncrement"`
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
	DeletedAt int64 `gorm:"index"`
}

func (*StructV20250211175923254CreateTableUser) TableName() string {
	return "user"
}
func NewMigrateV20250211175923254CreateTableUser() *MigrateV20250211175923254CreateTableUser {
	return &MigrateV20250211175923254CreateTableUser{
		currentTable: &StructV20250211175923254CreateTableUser{},
	}
}
func (r *MigrateV20250211175923254CreateTableUser) Register() {
	r.Tag = "v_2025_02_11_17_59_23_254_create_table_user"

}
func (r *MigrateV20250211175923254CreateTableUser) Up(tx *gorm.DB) error {

	err := tx.Migrator().CreateTable(r.currentTable)
	if err != nil {
		return err
	}
	return nil
}
func (r *MigrateV20250211175923254CreateTableUser) Down(tx *gorm.DB) error {
	err := tx.Migrator().DropTable(r.currentTable)
	if err != nil {
		return err
	}
	return nil
}
