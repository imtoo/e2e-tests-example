package database

import (
	"github.com/imtoo/e2e-tests-example/config"
	"github.com/imtoo/e2e-tests-example/models"
	"github.com/jinzhu/gorm"

	// needs to be here because of the DB Open
	_ "github.com/lib/pq"
)

// DB is a database struct
type DB struct {
	db *gorm.DB
}

// OpenDB opens connection do DB with credentials
func OpenDB() *gorm.DB {
	db, err := gorm.Open("postgres", config.EnvDatabaseURL)

	if err != nil {
		panic(err)
	}

	return db
}

// SetupDB runs migrations and seeds (if flag is set up)
func SetupDB() {
	db := OpenDB()

	models.AutoMigrate(db)

	if config.EnvRunSeeds {
		runSeeds(db)
	}
}
