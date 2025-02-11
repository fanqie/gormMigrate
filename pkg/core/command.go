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
			"\n\tsyntax：gmcmd gen  [--create|--alter]  {table_name}" +
			"\n\tusage：`gmcmd gen --create user` //or `gmcmd gen --alter user`",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				utility.ErrPrint("syntax error, gmcmd gen [--create|--alter]  {table_name}")
				return
			}
			action := ""
			if cmd.Flags().Changed("create") {
				action = "create"
			} else if cmd.Flags().Changed("alter") {
				action = "alter"
			}
			if action == "" {
				utility.ErrPrint("syntax error, gmcmd gen [--create|--alter] {table_name}")
			}
			tableName := args[0]

			if tableName == "" {
				utility.ErrPrint("tableName is required")
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
	genCommand.Flags().Bool("create", false, "set action to create the table,is default")
	genCommand.Flags().Bool("alter", false, "set action to alter the table")
	rootCmd.AddCommand(genCommand)

	rollbackCommand := &cobra.Command{
		Use:   "rollback",
		Short: "rollback migrations version by step",
		Run: func(cmd *cobra.Command, args []string) {
			var step int
			if cmd.Flags().Changed("step") {
				value, err := cmd.Flags().GetInt("step")
				if err != nil {
					utility.ErrPrintf("step is required, %v", err.Error())
					return
				}
				step = value
			}
			fmt.Printf("欢迎使用示例应用程序！%v", step)
		},
	}
	rollbackCommand.Flags().Int8P("step", "s", 1, "The default is to rollback one version. You can specify the number of versions to be rolled back and execute them in reverse order!")
	// 添加子命令
	rootCmd.AddCommand(rollbackCommand)

	migrateCommand := &cobra.Command{
		Use:   "migrate",
		Short: "all new migration file versions will be migrated or target step size version",
		Run: func(cmd *cobra.Command, args []string) {
			var step int
			if cmd.Flags().Changed("step") {
				value, err := cmd.Flags().GetInt("step")
				if err != nil {
					utility.ErrPrintf("step is required, %v", err.Error())
					return
				}
				step = value
			}
			fmt.Printf("欢迎使用示例应用程序！%v", step)
			err := MigrateHandle(step, *migrationsManage)
			if err != nil {
				utility.ErrPrintf("migrate error, %v", err.Error())
				return
			}
		},
	}

	migrateCommand.Flags().Int("step", 1, "By default, all new migration file versions will be migrated, and you can also set the migration step size")
	rootCmd.AddCommand(migrateCommand)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}

}
