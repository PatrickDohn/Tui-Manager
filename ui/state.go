package ui

import (
	"github.com/rivo/tview"
)

// UIState holds the references needed to control the app from any page
type UIState struct {
	App       *tview.Application
	MainPages *tview.Pages // The container that swaps between Home/Settings
	UserName  string       // Example of "Auth" state
	// DB             *gorm.DB
	// CurrentProject *db.Project
}
