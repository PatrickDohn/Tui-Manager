package db

import "gorm.io/gorm"

// --- CREATE ---

func CreateProject(db *gorm.DB, name string, desc string, notes string) (*Project, error) {
	project := Project{Name: name, Description: desc, Notes: notes}
	result := db.Create(&project)
	return &project, result.Error
}

func CreateTask(db *gorm.DB, title string, priority string, projectID *uint) (*Task, error) {
	task := Task{
		Title:     title,
		Priority:  priority,
		Status:    "Pending",
		ProjectID: projectID, // Will be nil for Personal tasks
	}
	result := db.Create(&task)
	return &task, result.Error
}

// --- READ ---

func GetAllProjects(db *gorm.DB) ([]Project, error) {
	var projects []Project
	err := db.Preload("Tasks").Find(&projects).Error
	return projects, err
}

func GetTasksByView(db *gorm.DB, viewMode string, projectID uint) ([]Task, error) {
	var tasks []Task
	query := db

	switch viewMode {
	case "personal":
		query = query.Where("project_id IS NULL")
	case "project":
		query = query.Where("project_id = ?", projectID)
	}

	err := query.Find(&tasks).Error
	return tasks, err
}

// --- UPDATE ---

func UpdateTaskStatus(db *gorm.DB, taskID uint, newStatus string) error {
	return db.Model(&Task{}).Where("id = ?", taskID).Update("status", newStatus).Error
}

// --- DELETE ---

func DeleteTask(db *gorm.DB, taskID uint) error {
	return db.Delete(&Task{}, taskID).Error
}
