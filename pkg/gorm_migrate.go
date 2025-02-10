package pkg

import (
	"fmt"
	"github.com/fanqie/gormMigrate/pkg/core"
	"github.com/fanqie/gormMigrate/pkg/impl"
)

type GormMigrate struct {
	migrations       map[string]impl.GormMigrateInterface
	DbTool           *core.DbTool
	MigrationsManage *core.MigratesManage
	isDebug          bool
}

func (r *GormMigrate) RegisterMigration(name string, migrationFunc impl.GormMigrateInterface) {
	r.migrations[name] = migrationFunc
}
func NewGormMigrate(isDebug bool) *GormMigrate {
	boot := &GormMigrate{
		migrations: make(map[string]impl.GormMigrateInterface),
		DbTool:     core.NewDbTool(),
		isDebug:    isDebug,
	}

	return boot
}

func (r *GormMigrate) Setup(db core.GromParams, afterHandle func()) {
	r.databaseInit(db)
	if r.isDebug {
		r.DbTool.Db.Debug()
	}
	r.MigrationsManage = core.NewMigratesManage()

	core.DefinedCommand(r.MigrationsManage)
	afterHandle()
	r.MigrationsManage.CheckTable()
}
func (r *GormMigrate) databaseInit(db core.GromParams) {
	err := r.DbTool.Open(db)
	if err != nil {
		fmt.Println(err)
		panic("the database connect error")
	}

}
