package main

import (
	"go-tui/api"
	"go-tui/db"
	"go-tui/ui"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/joho/godotenv"
	"github.com/rivo/tview"
)

var (
	DraculaBg      = tcell.NewRGBColor(40, 42, 54)    // #282A36
	DraculaCurrent = tcell.NewRGBColor(68, 71, 90)    // #44475A
	DraculaFg      = tcell.NewRGBColor(248, 248, 242) // #F8F8F2
	DraculaComment = tcell.NewRGBColor(98, 114, 164)  // #6272A4
	DraculaCyan    = tcell.NewRGBColor(139, 233, 253) // #8BE9FD
	DraculaGreen   = tcell.NewRGBColor(80, 250, 123)  // #50FA7B
	DraculaOrange  = tcell.NewRGBColor(255, 184, 108) // #FFB86C
	DraculaPink    = tcell.NewRGBColor(255, 121, 198) // #FF79C6
	DraculaPurple  = tcell.NewRGBColor(189, 147, 249) // #BD93F9
	DraculaRed     = tcell.NewRGBColor(255, 85, 85)   // #FF5555
	DraculaYellow  = tcell.NewRGBColor(241, 250, 140) // #F1FA8C
)

func main() {

	//Load env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, ")
	}

	// 1. Initialize database
	conn, _ := db.InitDB()

	// 2. Setup GitHub (Use os.Getenv so you don't hardcode your token!)
	token := os.Getenv("GITHUB_FG_TOKEN")
	ghClient := api.CreateGithubClient(token)

	app := tview.NewApplication()

	tview.Styles.PrimitiveBackgroundColor = DraculaBg
	tview.Styles.ContrastBackgroundColor = DraculaCurrent
	tview.Styles.PrimaryTextColor = DraculaFg
	tview.Styles.SecondaryTextColor = DraculaComment
	tview.Styles.BorderColor = DraculaPurple

	// db.CreateProject(conn, "Work", "Office related tasks", "REturn office supplies")
	// db.CreateProject(conn, "Res Creator", "resume creation app", "These are my notes.")

	// 1. rootPages will handle Login vs. The Whole App
	rootPages := tview.NewPages()

	// 2. contentPages handles the switching between Home and Settings
	contentPages := tview.NewPages()

	state := &ui.UIState{
		App:            app,
		MainPages:      contentPages, // Content switcher
		CurrentProject: nil,
		DB:             conn,
		GHClient:       ghClient,
		UserName:       "Gopher",
	}

	sidebar := ui.CreateSidebar(state)

	contentPages.AddPage("home", ui.CreateHomePage(state), true, true)
	contentPages.AddPage("settings", ui.CreateSettingsPage(state), true, false)
	contentPages.AddPage("project", ui.CreateProjectDetailPage(state), true, false)
	contentPages.AddPage("backlog", ui.CreateBacklogPage(state), true, false)

	// --- THE MAIN APP LAYOUT ---
	fullLayout := tview.NewFlex().
		AddItem(sidebar, 25, 1, true).
		AddItem(contentPages, 0, 1, false)

	// --- THE LOGIN PAGE ---
	// We pass rootPages so the login button can switch to "main_app"
	loginPage := ui.CreateLoginPage(state, rootPages, fullLayout)

	// --- WIRE IT UP ---
	rootPages.AddPage("login_screen", loginPage, true, false) // Visible first
	rootPages.AddPage("main_app", fullLayout, true, true)     // Hidden first

	if err := app.SetRoot(rootPages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
