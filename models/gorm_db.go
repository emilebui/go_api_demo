package models

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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
}

var DB = initDB()
