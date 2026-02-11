package db

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Description string
	Tasks       []Task // One-to-Many relationship
}

type Task struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Status   string // e.g., "Pending", "Done"
	Priority string // e.g., "High", "Low"
	Desc     string
	DueDate  time.Time
	// Using a pointer makes this field optional (nullable)
	ProjectID *uint `gorm:"index"` // Foreign Key linking to Project
}
