package core

import (
	"database/sql"
	"fmt"
	"github.com/fanqie/gormMigrate/pkg/utility"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type MigrateBasic struct {
	ID              int          `json:"id" yaml:"id" dc:"id" gorm:"primaryKey"`
	Tag             string       `json:"tag" yaml:"tag" dc:"timestampTag" gorm:"index"`
	AlreadyMigrated bool         `json:"alreadyMigrated" yaml:"alreadyMigrated" dc:"migrated status"`
	CreatedAt       time.Time    `json:"createdAt" yaml:"createdAt" dc:"createdAt"`
	UpdatedAt       time.Time    `json:"updatedAt" yaml:"updatedAt" dc:"updatedAt"`
	ExecutedAt      sql.NullTime `json:"executedAt" yaml:"executedAt" dc:"executedAt"`
	RevertedAt      sql.NullTime `json:"revertedAt" yaml:"revertedAt" dc:"revertedAt"`
}

func (*MigrateBasic) TableName() string {
	return "gorm_migrations"
}
func (r *MigrateBasic) GetData() *MigrateBasic {
	return r
}
func (r *MigrateBasic) GetTypeTag() string {
	// 把tag转为驼峰命名
	tags := strings.Split(r.Tag, "_")
	for i, tag := range tags {
		tags[i] = utility.FirstToUpper(tag)
	}
	return strings.Join(tags, "")
}
func (r *MigrateBasic) genRecord(args GenArgs) {
	dateStr := time.Now().Format("2006_01_02_15_04_05")

	r.Tag = fmt.Sprintf("v_%s_%s_%s_table_%s", dateStr, strconv.Itoa(rand.Intn(899)+100), args.Action, args.TableName)
	r.CreatedAt = time.Now()

}

func (r *MigrateBasic) UpAfter() {
}
func (r *MigrateBasic) DownAfter() {
	r.AlreadyMigrated = false
	r.RevertedAt = sql.NullTime{Time: time.Now(), Valid: true}
	//todo:update migration table
}
func (r *MigrateBasic) Up(tx *gorm.DB) error {
	return nil
}
func (r *MigrateBasic) Down(tx *gorm.DB) error { return nil }

func (r *MigrateBasic) Register() {
}
