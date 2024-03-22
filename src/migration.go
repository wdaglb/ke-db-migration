package src

import (
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
		filename := getFilename(src)
		if _, ok := migrationMap[filename]; ok {
			continue
		}

		dataBytes, err := os.ReadFile(src)
		if err != nil {
			log.Printf("migrate fail: %v\n", err)
			break
		}
		sql := string(dataBytes)

		data := domain.Migration{
			Version:  filename,
			File:     src,
			Complete: 0,
		}
		core.DB.Create(&data)

		err = core.DB.Exec(sql).Error
		if err != nil {
			log.Printf("migrate fail: %v\n", err)
			break
		}

		data.Complete = 1
		core.DB.Save(&data)
		num++
		log.Printf("migrate %s ok\n", data.Version)
	}
	log.Printf("migrate number:%d completed\n", num)
}
