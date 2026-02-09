package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// Open the SQLite database file
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate creates/updates tables based on your structs
	err = db.AutoMigrate(&Project{}, &Task{})
	return db, err
}
