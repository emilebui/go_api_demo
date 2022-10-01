package models

import "gorm.io/gorm"

type Repo struct {
	gorm.Model
	Name    string   `json:"name" gorm:"primaryKey"`
	Url     string   `json:"url"`
	Results []Result `gorm:"constraint:OnDelete:CASCADE;foreignkey:RepositoryName;references:Name"`
}

type Result struct {
	Id             string          `json:"id" gorm:"primaryKey"`
	Status         string          `json:"status"`
	RepositoryName string          `json:"repositoryName"`
	RepositoryUrl  string          `json:"repositoryUrl"`
	Findings       []Vulnerability `gorm:"constraint:OnDelete:CASCADE;foreignkey:ResultID;references:Id"`
	QueuedAt       int64           `json:"queuedAt"`
	ScanningAt     int64           `json:"scanningAt"`
	FinishedAt     int64           `json:"finishedAt"`
}

type Vulnerability struct {
	gorm.Model
	ID       string `json:"ID" gorm:"primaryKey"`
	ResultID string `json:"resultID"`
	Type     string `json:"type"`
	RuleID   string `json:"ruleID"`
	Rule     Rule
	Path     string `json:"path"`
	Line     string `json:"line"`
}

type Rule struct {
	gorm.Model
	RuleID          string          `json:"ruleId" gorm:"primaryKey"`
	Description     string          `json:"description"`
	Severity        string          `json:"severity"`
	StringCompare   string          `json:"stringCompare"`
	Vulnerabilities []Vulnerability `gorm:"constraint:OnDelete:CASCADE;foreignkey:RuleID;references:RuleID"`
}
