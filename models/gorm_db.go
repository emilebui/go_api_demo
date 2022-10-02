package models

import (
	"github.com/google/uuid"
	"go_api/helpers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// initDB Init DB object logic (for singleton DB)
func initDB() *gorm.DB {
	dbFile := helpers.GetDirPath("test.db")
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Repo{}, &Rule{}, &Result{}, &Vulnerability{})
	initRuleDB(db)

	return db
}

// initRuleDB Create initial rules to scan
func initRuleDB(db *gorm.DB) {
	db.Create(&Rule{ID: uuid.NewString(), Description: "repo", Severity: "HIGH", StringCompare: "public_key"})
	db.Create(&Rule{ID: uuid.NewString(), Description: "repo", Severity: "HIGH", StringCompare: "private_key"})
}

// InitTestDB Init Test DB
func InitTestDB() *gorm.DB {
	cxn := "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(cxn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	initTestData(db)
	return db
}

// initTestData Init Test DB Data
func initTestData(db *gorm.DB) {
	db.Migrator().DropTable(&Repo{})
	db.Migrator().DropTable(&Rule{})
	db.Migrator().DropTable(&Result{})
	db.Migrator().DropTable(&Vulnerability{})
	db.AutoMigrate(&Repo{}, &Rule{}, &Result{}, &Vulnerability{})
	db.Create(&Repo{Name: "test", Url: "https://github.com/emilebui/Minigames-on-Processing"})
	db.Create(&Rule{ID: "test", Description: "repo", Severity: "HIGH", StringCompare: "chosenRocket"})
	db.Create(&Result{
		Id: "test", Status: "Success", RepositoryName: "test",
		RepositoryUrl: "https://github.com/emilebui/Minigames-on-Processing",
		QueuedAt:      1, ScanningAt: 1, FinishedAt: 2,
	})
	db.Create(&Vulnerability{
		ID: "test", ResultID: "test", Type: "sast", RuleID: "test", Path: "blah", Line: 12,
	})
}

// DB Singleton DB object for entire app to use
var DB = initDB()
