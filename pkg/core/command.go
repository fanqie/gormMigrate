package core

import (
	"fmt"
	"github.com/fanqie/gormMigrate/pkg/utility"
	"github.com/spf13/cobra"
)

func DefinedCommand(migrationsManage *MigratesManage) {
	rootCmd := &cobra.Command{}
	genCommand := &cobra.Command{
		Use: "gen",
		Short: "generate a new core file" +
			"\n\tsyntax：gmcmd gen  [create|alter] {table_name}" +
			"\n\tusage：`gmcmd gen create user` //or `gmcmd gen alter user`",

		Run: func(cmd *cobra.Command, args []string) {
			// 这是 server 子命令的执行逻辑
			if len(args) < 2 {
				utility.ErrPrint("syntax error, gmcmd gen [create|alter] {table_name}")
				return
			}
			action := args[0]
			tableName := args[1]
			if tableName == "" {

				utility.ErrPrint("tableName is required")
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
				utility.ErrPrint("action type is required， the type value equal must “alter ”or “create”")
				return
			}
			err := gen(GenArgs{
				Action:    action,
				TableName: tableName,
			}, migrationsManage)
			if err != nil {
				utility.ErrPrint(err.Error())
				return
			}

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
		Use:   "core",
		Short: "all new migration file versions will be migrated or target step size version",
		Run: func(cmd *cobra.Command, args []string) {
			// 这是 server 子命令的执行逻辑
			fmt.Printf("欢迎使用示例应用程序！%v", args)
			fmt.Printf("欢迎使用示例应用程序！%v", args)
		},
	}
	migrateCommand.Flags().Int8P("step", "s", 1, "By default, all new migration file versions will be migrated, and you can also set the migration step size")
	rootCmd.AddCommand(migrateCommand)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}

}
