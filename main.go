package main

import (
	"ke-db-migration/core"
	"ke-db-migration/domain"
	"ke-db-migration/src"
)

func main() {
	core.InitApp()
	_ = core.DB.AutoMigrate(domain.Migration{})
	src.Migration()
}
