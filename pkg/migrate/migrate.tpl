package migrate

import (
	"fmt"
	"os"
	"strings"
	"time"
	gm "github.com/fanqie/gormMigrate/pkg"
)

type Migrate struct {
	gm.GormMigrate
}
func NewMigrate() *Migrate {
	return &Migrate{
	    Tag:"{{Tag}}"
	}
}
func (receiver *Migrate) Init() {
	receiver.GormMigrate.Init()
	receiver.GormMigrate.Tag = receiver.Tag
	receiver.GormMigrate.Up = receiver.Up
	receiver.GormMigrate.Down = receiver.Down
}

func (receiver *Migrate) Up() {
    // todo:migrate the database

}
func (receiver *Migrate) Down() {
    //todo:revert the migration
}
