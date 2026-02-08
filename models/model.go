package models

import "gorm.io/gorm"

type Valentines struct {
	gorm.Model
	Receipient    string
	Sender        string
	CreateId      string `gorm:"unique"`
	ClickYesFirst string
	YesClick      int
	NoClick       int
}
