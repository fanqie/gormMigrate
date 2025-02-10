package core

import (
	"fmt"
	"os"
	"strings"
)

func gen(action string, tableName string, migrationsManage *MigratesManage) {

	migrationsManage.RefreshMigrationsData()
	mb := &MigrateBasic{}
	mb.genRecord(action, tableName)
	Db.Create(mb)
	GenFile(mb)
	fmt.Printf("gen: %v,%s", action, tableName)
	//migrateFiles, err := os.ReadDir("./")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//获取所有迁移文件
	//生成一个迁移文件
	//从数据库里读取并刷新注册文件

}
func GenFile(r *MigrateBasic) any {

	migrateFileName := fmt.Sprintf("./gorm_migrations/migration_%s.go", r.Tag)

	saveFile(migrateFileName, makeMigrateFile(r))

	return nil
}
func makeMigrateFile(r *MigrateBasic) string {
	file, err := os.Open("../tpl/migration_{{tag}}.tpl")
	if err != nil {
		return fmt.Errorf("error opening file: %v", err).Error()

	}
	var contentBytes []byte

	_, err = file.Read(contentBytes)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err).Error()

	}
	content := string(contentBytes)
	fmt.Printf("makeMigrateFile: %v", content)
	strings.Replace(content, "{{Tag}}", r.Tag, -1)
	strings.Replace(content, "{{TypeTag}}", r.GetTypeTag(), -1)
	fmt.Printf("makeMigrateFileEnd: %v", content)
	//file, err = os.Create(fmt.Sprintf("%s/gorm_migrations/migrate_%s.go", utility.GetDir(), r.Tag))
	defer file.Close().Error()
	return content
}

// saveFile saves the content to the specified file path.
func saveFile(path string, content string) {
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
