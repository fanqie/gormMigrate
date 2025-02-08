package storage

import "gorm.io/gorm"

type DbTool struct {
	db *gorm.DB
}

func NewDbTool() *DbTool {
	return &DbTool{}
}
func (receiver *DbTool) Open(dialector gorm.Dialector, opts ...gorm.Option) error {
	db, err := gorm.Open(dialector, opts...)
	if err != nil {
		return err
	}
	receiver.db = db
	return err
}
