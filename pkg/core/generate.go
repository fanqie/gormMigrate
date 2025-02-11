package core

import (
	"fmt"
	"github.com/fanqie/gormMigrate/pkg/tpl"
	"github.com/fanqie/gormMigrate/pkg/utility"
	"gorm.io/gorm"
	"os"
	"strings"
)

type GenArgs struct {
	Action    string
	TableName string
}

func gen(genArgs GenArgs, migrationsManage *MigratesManage) error {

	mb := &MigrateBasic{}
	mb.genRecord(genArgs)
	return Db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(mb)
		if result.Error != nil {
			utility.ErrPrintf("insert error", result.Error)
			return result.Error
		}
		err := migrationsManage.RefreshMigrationsData(tx)
		if err != nil {
			return err
		}
		return GenFile(mb, migrationsManage, genArgs)
	})
}
func GenFile(r *MigrateBasic, migrationsManage *MigratesManage, genArgs GenArgs) error {

	migrateFileName := fmt.Sprintf("./gorm_migrations/migration_%s.go", r.Tag)
	registerFileName := "./gorm_migrations/register.go"

	err := saveFile(migrateFileName, makeMigrateFile(r, genArgs))
	if err != nil {
		return err
	}
	err = overwriteFile(registerFileName, refreshRenderRegisterCode(migrationsManage))
	if err != nil {
		return err
	}
	return nil
}
func refreshRenderRegisterCode(migrationsManage *MigratesManage) string {
	content := tpl.RegisterCode
	var registerMigration []string
	for _, migration := range migrationsManage.MigrateList {
		migrationCode := fmt.Sprintf("migrate.RegisterMigration(\"%s\", NewMigrate%s())", migration.GetTypeTag(), migration.GetTypeTag())
		registerMigration = append(registerMigration, migrationCode)
	}
	return strings.Replace(content, "{{RegisterMigration}}", strings.Join(registerMigration, "\n\t"), -1)
}
func makeMigrateFile(r *MigrateBasic, genArgs GenArgs) string {
	var content string
	switch genArgs.Action {
	case "create":
		content = tpl.MigrationCreateTableCode
		break
	case "alter":
		content = tpl.MigrationAlterTableCode
		break
	default:
		panic("not support action")
	}
	fmt.Printf("makeMigrateFile: %v", content)
	content = strings.Replace(content, "{{Tag}}", r.Tag, -1)
	content = strings.Replace(content, "{{TypeTag}}", r.GetTypeTag(), -1)
	content = strings.Replace(content, "{{TableStructName}}", utility.FirstToUpper(genArgs.TableName), -1)
	fmt.Printf("makeMigrateFileEnd: %v", content)
	return content
}

// saveFile saves the content to the specified file path.
func saveFile(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return err
	//defer file.Close().Error()
}
func overwriteFile(path string, content string) error {
	err := os.Remove(path)
	if err != nil {
		utility.ErrPrintf("Error deleting file:", err)
		return err
	}
	return saveFile(path, content)

}
