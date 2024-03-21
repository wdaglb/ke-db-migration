package src

import (
	"gorm.io/gorm"
	"ke-db-migration/config"
	"ke-db-migration/core"
	"ke-db-migration/domain"
	"log"
	"os"
)

func Migration() {
	files := scanMigration(config.Config.MigrationDir)
	var migrations []domain.Migration
	core.DB.Where("complete=1").Order("version asc").Find(&migrations)
	migrationMap := make(map[string]uint)
	for _, migration := range migrations {
		migrationMap[migration.Version] = migration.ID
	}

	err := core.DB.Transaction(func(tx *gorm.DB) error {
		for _, src := range files {
			filename := getFilename(src)
			if _, ok := migrationMap[filename]; ok {
				continue
			}

			dataBytes, err := os.ReadFile(src)
			if err != nil {
				continue
			}
			sql := string(dataBytes)

			data := domain.Migration{
				Version:  filename,
				File:     src,
				Complete: 0,
			}
			tx.Create(&data)

			err = tx.Exec(sql).Error
			if err != nil {
				return err
			}

			data.Complete = 1
			tx.Save(&data)
			log.Printf("migrate %s ok\n", data.Version)
		}
		return nil
	})
	if err != nil {
		log.Printf("migrate fail: %v\n", err)
		return
	}
	log.Printf("migrate number:%d completed\n", len(files))
}
