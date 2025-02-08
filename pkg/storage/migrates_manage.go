package storage

import (
	"github.com/fanqie/gormMigrate/pkg/impl"
	"github.com/fanqie/gormMigrate/pkg/migrate"
)

type MigratesManage struct {
	migrateList []*impl.GormMigrateInterface
	TableName   string
	DbTool      *DbTool
}

func NewMigratesManage(tool *DbTool) *MigratesManage {
	return &MigratesManage{
		DbTool: tool,
	}
}

func (r MigratesManage) CheckTable() {
	if !r.DbTool.db.Migrator().HasTable(&migrate.Basic{}) {
		err := r.DbTool.db.AutoMigrate(&migrate.Basic{})
		if err != nil {
			return
		}
	}
}
