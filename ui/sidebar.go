package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CreateSidebar(state *UIState) tview.Primitive {
	// 1. MAIN TASKS SECTION
	mainTasks := tview.NewList().
		AddItem("Daily", "Today's focus", 'd', func() {
			state.MainPages.SwitchToPage("home")
		}).
		AddItem("Upcoming", "Next few days", 'u', nil).
		AddItem("Backlog", "Future tasks", 'b', func() {
			state.MainPages.SwitchToPage("settings")
		})

	mainTasks.SetBorder(false) // We'll put the border on the outer flex instead

	// 2. PROJECTS SECTION
	projects := tview.NewList().
		AddItem("Project Alpha", "Go TUI development", '1', nil).
		AddItem("Project Beta", "SQLite Integration", '2', nil).
		AddItem("New Project", "+ Add dynamic project", 'n', nil)

	projects.SetBorder(false)

	// 3. SECTION HEADERS
	headerMain := tview.NewTextView().SetText("--- TASKS ---").SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorYellow)
	headerProj := tview.NewTextView().SetText("--- PROJECTS ---").SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorYellow)

	// 4. ASSEMBLE SIDEBAR FLEX
	sidebarFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(headerMain, 1, 1, false).
		AddItem(mainTasks, 8, 1, true).       // Give fixed height for main list
		AddItem(tview.NewBox(), 1, 1, false). // Small gap
		AddItem(headerProj, 1, 1, false).
		AddItem(projects, 0, 1, false).      // This expands to fill the rest
		AddItem(tview.NewBox(), 0, 1, false) // Pushes everything up

	sidebarFlex.SetBorder(true).SetTitle("Menu")

	return sidebarFlex
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
