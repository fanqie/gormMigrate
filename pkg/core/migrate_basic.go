package core

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type MigrateBasic struct {
	Tag             string    `json:"tag" yaml:"tag" dc:"timestampTag"`
	AlreadyMigrated bool      `json:"alreadyMigrated" yaml:"alreadyMigrated" dc:"migrated status"`
	CreatedAt       time.Time `json:"createdAt" yaml:"createdAt" dc:"createdAt"`
	ExecutedAt      time.Time `json:"executedAt" yaml:"executedAt" dc:"executedAt"`
	RevertedAt      time.Time `json:"revertedAt" yaml:"revertedAt" dc:"revertedAt"`
}

func (*MigrateBasic) TableName() string {
	return "gorm_migrations"
}
func (r *MigrateBasic) GetTypeTag() string {
	// 把tag转为驼峰命名
	tags := strings.Split(r.Tag, "_")
	for i, tag := range tags {
		tags[i] = strings.ToTitleSpecial(unicode.SpecialCase{}, tag)
	}
	return strings.Join(tags, "")
}
func (r *MigrateBasic) genRecord(action string, migrateTableName string) {
	dateStr := time.Now().Format("2006_01_02_15_04_05")

	r.Tag = fmt.Sprintf("v_%s_%s_%s_table_%s", dateStr, strconv.Itoa(rand.Intn(899)+100), action, migrateTableName)
	r.CreatedAt = time.Now()

}

func (r *MigrateBasic) UpAfter() {
	r.AlreadyMigrated = true
	r.ExecutedAt = time.Now()
	//todo:update migration table
}
func (r *MigrateBasic) DownAfter() {
	r.AlreadyMigrated = false
	r.RevertedAt = time.Now()
	//todo:update migration table
}
func (r *MigrateBasic) Up() {
}
func (r *MigrateBasic) Down() {
}

func (r *MigrateBasic) Register() {
}
