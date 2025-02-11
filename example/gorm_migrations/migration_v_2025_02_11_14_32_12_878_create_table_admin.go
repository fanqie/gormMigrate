package gorm_migrations

import (
	migrate "github.com/fanqie/gormMigrate/pkg/core"
)

type Admin struct {
	Id        int64 `gorm:"primaryKey;autoIncrement"`
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
	DeletedAt int64 `gorm:"index"`
}
type MigrateV20250211143212878CreateTableAdmin struct {
	migrate.MigrateBasic
}

func NewMigrateV20250211143212878CreateTableAdmin() *MigrateV20250211143212878CreateTableAdmin {
	return &MigrateV20250211143212878CreateTableAdmin{}
}
func (r *MigrateV20250211143212878CreateTableAdmin) Register() {
	r.Tag = "v_2025_02_11_14_32_12_878_create_table_admin"

}
func (r *MigrateV20250211143212878CreateTableAdmin) Up() {
	err := migrate.Migrator().CreateTable(&Admin{})
	if err != nil {
		panic(err)
	}
}
func (r *MigrateV20250211143212878CreateTableAdmin) Down() {
	err := migrate.Migrator().DropTable(&Admin{})
	if err != nil {
		panic(err)
	}
}
