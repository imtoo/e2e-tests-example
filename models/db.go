package models

import (
	"github.com/jinzhu/gorm"
)

// StoreType is struct of the DB
type StoreType struct {
	DB *gorm.DB
}
