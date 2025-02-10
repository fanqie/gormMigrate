package core

import (
	"github.com/fanqie/gormMigrate/pkg/impl"
	"github.com/fanqie/gormMigrate/pkg/utility"
)

type MigratesManage struct {
	MigrateList []impl.GormMigrateInterface
	AlreadyList []impl.GormMigrateInterface
	PendingList []impl.GormMigrateInterface
	TableName   string
}

func NewMigratesManage() *MigratesManage {
	return &MigratesManage{}
}
func (r *MigratesManage) RefreshMigrationsData() {
	var migrateList []*MigrateBasic
	result := Db.Order("created_at asc").Find(&migrateList)
	if result.Error != nil {
		utility.ErrPrintf("the database connect error:%s", result.Error.Error())
		return
	}
	r.MigrateList = make([]impl.GormMigrateInterface, 0)
	r.AlreadyList = make([]impl.GormMigrateInterface, 0)
	r.PendingList = make([]impl.GormMigrateInterface, 0)
	for _, s := range migrateList {
		r.MigrateList = append(r.MigrateList, s)
		if s.AlreadyMigrated {
			r.AlreadyList = append(r.AlreadyList, s)
		} else {
			r.PendingList = append(r.PendingList, s)
		}
	}
}

func (r *MigratesManage) CheckTable() {
	if !Db.Migrator().HasTable(&MigrateBasic{}) {
		err := Db.AutoMigrate(&MigrateBasic{})
		if err != nil {
			return
		}
	}
}
