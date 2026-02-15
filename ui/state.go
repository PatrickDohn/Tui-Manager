package ui

import (
	"go-tui/db"

	"github.com/google/go-github/v60/github"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

// UIState holds the references needed to control the app from any page
type UIState struct {
	App            *tview.Application
	MainPages      *tview.Pages // The container that swaps between Home/Settings
	UserName       string       // Example of "Auth" state
	DB             *gorm.DB
	GHClient       *github.Client
	CurrentProject *db.Project
	RefreshSidebar func() // function pointer: "hook" for other components to call
}
