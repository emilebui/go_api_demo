package models

import (
	"github.com/google/uuid"
	"go_api/helpers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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

func initRuleDB(db *gorm.DB) {
	db.Create(&Rule{RuleID: uuid.NewString(), Description: "demo", Severity: "HIGH", StringCompare: "public_key"})
	db.Create(&Rule{RuleID: uuid.NewString(), Description: "demo", Severity: "HIGH", StringCompare: "private_key"})
	db.Create(&Rule{RuleID: uuid.NewString(), Description: "demo", Severity: "HIGH", StringCompare: "chosenRocket"})
}

var DB = initDB()
