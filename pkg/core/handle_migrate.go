package core

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

func MigrateHandle(step int, manage MigratesManage) error {
	return Db.Transaction(func(tx *gorm.DB) error {
		err := manage.RefreshMigrationsData(tx)
		if err != nil {
			return err
		}
		var count int
		for _, m := range manage.PendingList {

			if step == 0 || count < step {
				err := m.Up(tx)
				if err != nil {
					return err
				}
				m.UpAfter()
				count++
			}
			m.AlreadyMigrated = true
			m.ExecutedAt = sql.NullTime{Time: time.Now(), Valid: true}
			err := tx.Save(&m).Error
			if err != nil {
				return err
			}
		}
		return nil

	})

}
