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
	log.Printf("migrationDir: %s\n", config.Config.MigrationDir)
	files := scanMigration(config.Config.MigrationDir)
	if len(files) == 0 {
		log.Printf("migrate empty\n")
		return
	}
	var migrations []domain.Migration
	core.DB.Where("complete=1").Order("version asc").Find(&migrations)
	migrationMap := make(map[string]uint)
	for _, migration := range migrations {
		migrationMap[migration.Version] = migration.ID
	}

	num := 0

	for _, src := range files {
		err := core.DB.Transaction(func(tx *gorm.DB) error {
			filename := getFilename(src)
			if _, ok := migrationMap[filename]; ok {
				return nil
			}

			dataBytes, err := os.ReadFile(src)
			if err != nil {
				return err
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
			num++
			log.Printf("migrate %s ok\n", data.Version)
			return nil
		})
		if err != nil {
			log.Printf("migrate fail: %v\n", err)
			return
		}
	}
	log.Printf("migrate number:%d completed\n", num)
}
