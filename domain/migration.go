package domain

import "gorm.io/gorm"

type Migration struct {
	gorm.Model
	Version  string
	File     string
	Complete uint
}
