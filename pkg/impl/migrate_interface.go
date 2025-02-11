package impl

import (
	"github.com/fanqie/gormMigrate/pkg/core"
	"gorm.io/gorm"
)

type GormMigrateInterface interface {
	Up(tx *gorm.DB) error
	Down(tx *gorm.DB) error
	Register()
	UpAfter()
	DownAfter()
	GetTypeTag() string
	TableName() string
	GetData() *core.MigrateBasic
}
