package models

import "gorm.io/gorm"

type Repo struct {
	gorm.Model
	Name    string   `json:"name" gorm:"primaryKey;unique"`
	Url     string   `json:"url"`
	Results []Result `gorm:"constraint:OnDelete:CASCADE;foreignkey:RepositoryName;references:Name"`
}

type Result struct {
	Id             string          `json:"id" gorm:"primaryKey;unique"`
	Status         string          `json:"status"`
	RepositoryName string          `json:"repositoryName"`
	RepositoryUrl  string          `json:"repositoryUrl"`
	Findings       []Vulnerability `gorm:"constraint:OnDelete:CASCADE;foreignkey:ResultID;references:Id"`
	ErrorLog       string          `json:"errorLog"`
	QueuedAt       int64           `json:"queuedAt"`
	ScanningAt     int64           `json:"scanningAt"`
	FinishedAt     int64           `json:"finishedAt"`
}

type Vulnerability struct {
	gorm.Model
	ID       string `json:"ID" gorm:"primaryKey;unique"`
	ResultID string `json:"resultID"`
	Type     string `json:"type"`
	RuleID   string `json:"ruleID"`
	Rule     Rule
	Path     string `json:"path"`
	Line     int    `json:"line"`
}

type Rule struct {
	gorm.Model
	ID              string          `json:"Id" gorm:"primaryKey;unique"`
	Description     string          `json:"description"`
	Severity        string          `json:"severity"`
	StringCompare   string          `json:"stringCompare" gorm:"unique"`
	Vulnerabilities []Vulnerability `gorm:"constraint:OnDelete:CASCADE;foreignkey:RuleID;references:ID"`
}
