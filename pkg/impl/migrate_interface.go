package impl

type GormMigrateInterface interface {
	Up()
	Down()
	Register()
	UpAfter()
	DownAfter()
}
