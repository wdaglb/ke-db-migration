package src

import (
	"fmt"
	"ke-db-migration/config"
	"ke-db-migration/core"
	"ke-db-migration/domain"
	"os"
)

func Migration() {
	absDir := getAbsDir(config.Config.MigrationDir)
	core.Logger.Infof("migrationDir: %s\n", absDir)
	files := scanMigration(absDir)
	if len(files) == 0 {
		core.Logger.Infof("migrate empty\n")
		return
	}
	var migrations []domain.Migration
	core.DB.Where("complete=1").Order("version asc").Find(&migrations)
	migrationMap := make(map[string]uint)
	for _, migration := range migrations {
		migrationMap[migration.Version] = migration.ID
	}

	num := 0
	notify := NewNotify()

	for _, src := range files {
		filename := getFilename(src)
		if _, ok := migrationMap[filename]; ok {
			continue
		}

		dataBytes, err := os.ReadFile(src)
		if err != nil {
			_ = notify.Qywx(fmt.Sprintf("数据库迁移失败 %s", err))
			core.Logger.Errorf("migrate fail: %v\n", err)
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
			_ = notify.Qywx(fmt.Sprintf("数据库迁移失败 %s", err))
			core.Logger.Errorf("migrate fail: %v\n", err)
			break
		}

		data.Complete = 1
		core.DB.Save(&data)
		num++
		core.Logger.Infof("migrate %s ok\n", data.Version)
	}
	core.Logger.Infof("migrate number:%d completed\n", num)
}
