package core

import (
	"github.com/fanqie/gormMigrate/pkg/utility"
	"gorm.io/gorm"
)

type MigratesManage struct {
	MigrateList []MigrateBasic
	AlreadyList []MigrateBasic
	PendingList []MigrateBasic
	TableName   string
}

func NewMigratesManage() *MigratesManage {
	return &MigratesManage{}
}
func (r *MigratesManage) RefreshMigrationsData(tx *gorm.DB) error {
	var migrateList []*MigrateBasic
	result := tx.Order("created_at asc").Find(&migrateList)
	if result.Error != nil {
		utility.ErrPrintf("the database connect error:%s", result.Error.Error())
		return result.Error
	}
	r.MigrateList = make([]MigrateBasic, 0)
	r.AlreadyList = make([]MigrateBasic, 0)
	r.PendingList = make([]MigrateBasic, 0)
	for _, s := range migrateList {
		r.MigrateList = append(r.MigrateList, *s)
		if s.AlreadyMigrated {
			r.AlreadyList = append(r.AlreadyList, *s)
		} else {
			r.PendingList = append(r.PendingList, *s)
		}
	}
	return nil
}

func (r *MigratesManage) CheckTable() {
	if !Db.Migrator().HasTable(&MigrateBasic{}) {
		err := Db.AutoMigrate(&MigrateBasic{})
		if err != nil {
			return
		}
	}
}
