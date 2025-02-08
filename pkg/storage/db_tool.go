package storage

import "gorm.io/gorm"

type DbTool struct {
	db *gorm.DB
}
type GromParams struct {
	Dialector gorm.Dialector
	Opts      *gorm.Config
}

func NewDbTool() *DbTool {
	return &DbTool{}
}
func (receiver *DbTool) Open(params GromParams) error {
	db, err := gorm.Open(params.Dialector, params.Opts)
	if err != nil {
		return err
	}
	db = db.Debug()
	receiver.db = db

	return err
}
