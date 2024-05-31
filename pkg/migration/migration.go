package migration

import (
	"s21_go/pkg/db"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)


func MigrateDB(database *gorm.DB) error {
	m := gormigrate.New(database, gormigrate.DefaultOptions, []*gormigrate.Migration{
		
		{
			ID: "202405310001",
			Migrate: func(tx *gorm.DB) error {
				
				if err := tx.AutoMigrate(&db.TransmitterData{}); err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				
				if err := tx.Migrator().DropTable(&db.TransmitterData{}); err != nil {
					return err
				}
				return nil
			},
		},
	})


	if err := m.Migrate(); err != nil {
		return err
	}
	return nil
}
