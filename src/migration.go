package src

import (
	"fmt"
	"gorm.io/gorm"
	"ke-db-migration/config"
	"ke-db-migration/core"
	"ke-db-migration/domain"
	"ke-db-migration/utils"
	"os"
	"strings"
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
	core.DB.Order("version asc").Find(&migrations)
	migrationMap := make(map[string]uint)
	failMap := make(map[string]uint)
	for _, migration := range migrations {
		if migration.Complete == 1 {
			migrationMap[migration.Version] = migration.ID
		} else {
			failMap[migration.Version] = migration.ID
		}
	}

	num := 0
	notify := NewNotify()

	for _, src := range files {
		filename := getFilename(src)
		if strings.HasPrefix(filename, "ignore") {
			continue
		}
		if _, ok := migrationMap[filename]; ok {
			continue
		}

		dataBytes, err := os.ReadFile(src)
		if err != nil {
			core.Logger.Errorf("migrate fail: %v\n", err)
			if config.Config.SkipError {
				continue
			}
			_ = notify.Qywx(fmt.Sprintf("数据库迁移失败 %s", err))
			break
		}
		sql := string(dataBytes)

		data := domain.Migration{
			Version:  filename,
			File:     src,
			Complete: 0,
		}
		if _, ok := failMap[filename]; !ok {
			core.DB.Create(&data)
		}

		err = core.DB.Transaction(func(tx *gorm.DB) error {
			sqlList := strings.Split(utils.FilterSpaceLine(sql), ";")
			for _, sqlItem := range sqlList {
				if strings.TrimSpace(sqlItem) == "" {
					continue
				}
				er := tx.Exec(sqlItem).Error
				if er != nil {
					return er
				}
			}
			return nil
		})
		if err != nil {
			_ = notify.Qywx(fmt.Sprintf("数据库迁移失败 [%s] %s", data.Version, err))
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
