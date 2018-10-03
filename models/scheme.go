package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

// Models is list of all models in application
// For migrations and clean purposes
var Models = []interface{}{&Article{}}

// AutoMigrate runs migrations according to scheme and models passed into func
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(Models...)
}

// Article - db Article type
type Article struct {
	gorm.Model
	URL                string         `gorm:"type:text;not null;"`
	Channel            string         `gorm:"type:character varying(255);not null;"`
	Username           string         `gorm:"type:character varying(100);not null;"`
	ApprovedForSlack   *bool          `gorm:"type:boolean;default:true;"`
	SentToSlack        *bool          `gorm:"type:boolean;default:false;"`
	ApprovedForTwitter *bool          `gorm:"type:boolean;default:false;"`
	SentToTwitter      *bool          `gorm:"type:boolean;default:false;"`
	TwitterMessage     sql.NullString `gorm:"type:text;"`
}
