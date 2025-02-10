package storage

import (
	"github.com/fanqie/gormMigrate/pkg/utility"
	"gorm.io/gorm"
)

type DbTool struct {
	db *gorm.DB
}
type GromParams struct {
	Dialector gorm.Dialector
	Opts      *gorm.Config
}

func NewDbTool() *DbTool {
	return &DbTool{}
}
func (receiver *DbTool) Open(params GromParams) error {
	db, err := gorm.Open(params.Dialector, params.Opts)
	if err != nil {
		utility.ErrPrintf("the database connect error:%s", err.Error())
		return err
	}
	//db = db.Debug()
	receiver.db = db

	return err
}
