package storage

import "github.com/fanqie/gormMigrate/pkg/impl"

type MigratesManage struct {
	migrateList []*impl.GormMigrateInterface
	TableName   string
}

func NewMigratesManage() *MigratesManage {
	return &MigratesManage{
		TableName: "migrates",
	}
}
func (receiver MigratesManage) CheckTable() {

}
