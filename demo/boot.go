package demo

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Init() {

	definedCommand()
}

func definedCommand() {
	rootCmd := &cobra.Command{
		Use:   "",
		Short: "",
		Long:  "",
		//Run: func(cmd *cobra.Command, args []string) {
		//	// 这是根命令的执行逻辑
		//	fmt.Printf("欢迎使用示例应用程序！%v", args)
		//	register()
		//},
	}
	genCommand := &cobra.Command{
		Use: "gen",
		Short: "generate a new migrate file" +
			"\n\tsyntax：gmcmd gen  [create|alter] {table_name}" +
			"\n\tusage：`gmcmd gen create user` //or `gmcmd gen alter user`",

		Run: func(cmd *cobra.Command, args []string) {
			// 这是 server 子命令的执行逻辑
			if len(args) < 2 {
				fmt.Printf("[Error]syntax error, gmcmd gen [create|alter] {table_name}")
				return
			}
			action := args[0]
			tableName := args[1]
			if tableName == "" {
				fmt.Printf("[Error]tableName is required")
				return
			}
			actions := []string{"alter", "create"}

			found := false
			for _, a := range actions {
				if a == action {
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("[Error]action type is required， the type value equal must “alter ”or “create”")
				return
			}
			gen(action, tableName)

		},
	}
	genCommand.Flags().StringP("create", "c", "yes", "set action to create the table")
	genCommand.Flags().StringP("modify", "m", "no", "set action to modify the table")
	rootCmd.AddCommand(genCommand)

	rollbackCommand := &cobra.Command{
		Use:   "rollback",
		Short: "rollback migrations version by step",
		Run: func(cmd *cobra.Command, args []string) {
			// 这是 server 子命令的执行逻辑
			fmt.Printf("欢迎使用示例应用程序！%v", args)
			fmt.Printf("欢迎使用示例应用程序！%v", args)
		},
	}
	rollbackCommand.Flags().Int8P("step", "s", 1, "The default is to rollback one version. You can specify the number of versions to be rolled back and execute them in reverse order!")
	// 添加子命令
	rootCmd.AddCommand(rollbackCommand)

	migrateCommand := &cobra.Command{
		Use:   "migrate",
		Short: "all new migration file versions will be migrated or target step size version",
		Run: func(cmd *cobra.Command, args []string) {
			// 这是 server 子命令的执行逻辑
			fmt.Printf("欢迎使用示例应用程序！%v", args)
			fmt.Printf("欢迎使用示例应用程序！%v", args)
		},
	}
	migrateCommand.Flags().Int8P("step", "s", 1, "By default, all new migration file versions will be migrated, and you can also set the migration step size")
	rootCmd.AddCommand(migrateCommand)
	println(rootCmd.Execute().Error())

}

func gen(action string, tableName string) {
	fmt.Printf("gen: %v,%s", action, tableName)
	//解析命令行参数
	//获取所有迁移文件
	//生成一个迁移文件
	//从数据库里读取并刷新注册文件

}
