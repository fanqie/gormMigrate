package migrate

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Basic struct {
	Tag             string    `json:"tag" yaml:"tag" dc:"timestampTag"`
	AlreadyMigrated bool      `json:"alreadyMigrated" yaml:"alreadyMigrated" dc:"migrated status"`
	CreatedAt       time.Time `json:"createdAt" yaml:"createdAt" dc:"createdAt"`
	ExecutedAt      time.Time `json:"executedAt" yaml:"executedAt" dc:"executedAt"`
	RevertedAt      time.Time `json:"revertedAt" yaml:"revertedAt" dc:"revertedAt"`
}

func (Basic) TableName() string {
	return "gorm_migrations"
}
func (receiver *Basic) insertRecord() {
	receiver.Tag = fmt.Sprintf("_v_%s", time.Now())
	receiver.CreatedAt = time.Now()
}
func (receiver *Basic) GenFile() any {
	receiver.insertRecord()
	migrateFileName := fmt.Sprintf("./migrations/migrate_%s.go", receiver.Tag)

	receiver.saveFile(migrateFileName, receiver.makeMigrateFile())

	return nil
}
func (receiver *Basic) makeMigrateFile() string {
	file, err := os.Open("../tpl/migrate_{{tag}}.tpl")
	if err != nil {
		return fmt.Errorf("error opening file: %v", err).Error()

	}
	var contentBytes []byte

	_, err = file.Read(contentBytes)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err).Error()

	}
	content := string(contentBytes)
	strings.Replace(content, "{{Tag}}", receiver.Tag, -1)
	file, err = os.Create("../migrations/" + receiver.Tag + ".go")
	defer file.Close().Error()
	return content
}

// saveFile saves the content to the specified file path.
func (receiver *Basic) saveFile(path string, content string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	defer file.Close().Error()
}
func (receiver *Basic) UpAfter() {
	receiver.AlreadyMigrated = true
	receiver.ExecutedAt = time.Now()
	//todo:update migration table
}
func (receiver *Basic) DownAfter() {
	receiver.AlreadyMigrated = false
	receiver.RevertedAt = time.Now()
	//todo:update migration table
}
