package main

import (
	"go-tui/ui"

	"github.com/rivo/tview"
)

func main() {

	// Initialize database
	// conn, _ := db.InitDB()

	// proj, _ := db.CreateProject(conn, "Work", "Office related tasks")

	app := tview.NewApplication()

	// 1. rootPages will handle Login vs. The Whole App
	rootPages := tview.NewPages()

	// 2. contentPages handles the switching between Home and Settings
	contentPages := tview.NewPages()

	state := &ui.UIState{
		App:       app,
		MainPages: contentPages, // Content switcher
		// CurrentProject: proj,
		// DB:             conn,
		UserName: "Gopher",
	}

	// --- YOUR EXISTING SIDEBAR LOGIC ---
	// sidebar := tview.NewList().
	// 	AddItem("Home", "Go to dashboard", 'h', func() {
	// 		contentPages.SwitchToPage("home")
	// 	}).
	// 	AddItem("Settings", "Change preferences", 's', func() {
	// 		contentPages.SwitchToPage("settings")
	// 	}).
	// 	AddItem("Quit", "Press to exit", 'q', func() {
	// 		app.Stop()
	// 	})
	// sidebar.SetBorder(true).SetTitle("Menu")

	sidebar := ui.CreateSidebar(state)

	// divider := tview.NewTextView().
	// 	SetTextColor(tcell.ColorIndianRed).
	// 	SetText("│"). // Vertical line character
	// 	SetTextAlign(tview.AlignCenter)

	contentPages.AddPage("home", ui.CreateHomePage(state), true, true)
	contentPages.AddPage("settings", ui.CreateSettingsPage(state), true, false)

	// --- THE MAIN APP LAYOUT ---
	fullLayout := tview.NewFlex().
		AddItem(sidebar, 25, 1, true).
		AddItem(contentPages, 0, 1, false)

	// --- THE LOGIN PAGE ---
	// We pass rootPages so the login button can switch to "main_app"
	loginPage := ui.CreateLoginPage(state, rootPages, fullLayout)

	// --- WIRE IT UP ---
	rootPages.AddPage("login_screen", loginPage, true, true) // Visible first
	rootPages.AddPage("main_app", fullLayout, true, false)   // Hidden first

	if err := app.SetRoot(rootPages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
