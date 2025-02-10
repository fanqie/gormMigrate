package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 定义文件夹路径
	dir := "gorm_migrations"

	// 检查文件夹是否存在
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 文件夹不存在，创建它
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败:", err)
			return
		}
		fmt.Println("文件夹已创建:", dir)
	} else {
		fmt.Println("文件夹已存在:", dir)
	}
	// 定义文件路径

	registerPath := filepath.Join("gorm_migrations", "register.go")
	// 检查 gorm_migrations/register.go 是否存在
	if err := checkAndCreateFile(registerPath, registerContent()); err != nil {
		fmt.Println("创建 register.go 失败:", err)
	}
	gmcmdPath := "gmcmd.go"
	// 检查 gmcmd.go 是否存在
	if err := checkAndCreateFile(gmcmdPath, gmcmdContent()); err != nil {
		fmt.Println("创建 gmcmd.go 失败:", err)
	}

}

// checkAndCreateFile 检查文件是否存在，不存在则创建并写入内容
func checkAndCreateFile(path, content string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 文件不存在，创建它
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// 写入内容
		_, err = file.WriteString(content)
		if err != nil {
			return err
		}
		fmt.Printf("文件已创建: %s\n", path)
	} else {
		fmt.Printf("文件已存在: %s\n", path)
	}
	return nil
}

// gmcmdContent 返回 gmcmd.go 的内容
func gmcmdContent() string {
	return `package main
import (
	"github.com/fanqie/gormMigrate/gorm_migrations"
	gm "github.com/fanqie/gormMigrate/pkg"
	"github.com/fanqie/gormMigrate/pkg/storage"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	gormMigrate := gm.NewGormMigrate()
	
	gormMigrate.Setup(storage.GromParams{
		Dialector: mysqlDialector(),
		Opts: &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	}, func() {
		gorm_migrations.Register(gormMigrate)
	})
}
func mysqlDialector() *Dialector{
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	return mysql.Open(dsn);
}

`
}

// registerContent 返回 register.go 的内容
func registerContent() string {
	return `package gorm_migrations

import (
	"github.com/fanqie/gormMigrate/pkg"
)

func Register(migrate *pkg.GormMigrate) {
}
`
}
