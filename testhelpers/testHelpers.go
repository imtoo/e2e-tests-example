package testhelpers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	txdb "github.com/DATA-DOG/go-txdb"
	"github.com/gorilla/mux"
	"github.com/imtoo/e2e-tests-example/config"
	"github.com/imtoo/e2e-tests-example/models"
	"github.com/jinzhu/gorm"

	// needs to be here because of the DB Open
	_ "github.com/lib/pq"
)

// ExecuteRequest executes request for testing purposes
func ExecuteRequest(req *http.Request, router *mux.Router) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

// RegisterTxDB registers new db for single transaction tests
func RegisterTxDB(name string) {
	txdb.Register(name, "postgres", config.EnvDatabaseURL)
}

// PrepareTestDB prepare test DB according to txdb name
func PrepareTestDB(withName string) (*gorm.DB, error) {
	sqlDB, err := sql.Open(withName, fmt.Sprintf("connection_%d", time.Now().UnixNano()))
	db, err := gorm.Open("postgres", sqlDB)
	models.AutoMigrate(db)

	if err != nil {
		panic(err)
	}

	return db, err
}

// CleanTestDB drops all tables from test DB
func CleanTestDB(db *gorm.DB) {
	db.Close()
}
