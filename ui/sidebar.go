package ui

import (
	"go-tui/db"

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
	var projects []db.Project

	projList := tview.NewList().SetCurrentItem(0)

	projList.Clear()

	state.DB.Find(&projects)

	if len(projects) == 0 {
		projList.AddItem(
			"[gray]No projects[-]",
			"",
			0,
			nil,
		)
		projList.SetCurrentItem(0)
	} else {
		for i := range projects {
			proj := &projects[i] // safe pointer to slice element

			projList.AddItem(
				proj.Name,
				proj.Description,
				0,
				func() {
					state.CurrentProject = proj

					projectPage := CreateProjectDetailPage(state)

					state.MainPages.AddAndSwitchToPage(
						"project",
						projectPage,
						true,
					)
				},
			)
		}
	}

	settings := tview.NewList().
		AddItem("App Settings", "", 's', func() {
			state.MainPages.SwitchToPage("settings")
		}).
		AddItem("Quit", "Future tasks", 'q', func() {
			state.App.Stop()
		})

	// 3. SECTION HEADERS
	headerMain := tview.NewTextView().SetText("--- TASKS ---").SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorYellow)

	headerProj := tview.NewTextView().SetText("--- PROJECTS ---").SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorYellow)

	// 4. ASSEMBLE SIDEBAR FLEX
	sidebarFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		// 1. TOP SECTION: Tasks
		AddItem(headerMain, 1, 1, false).
		AddItem(mainTasks, 8, 1, true).

		// 2. MIDDLE SECTION: Projects
		AddItem(tview.NewBox(), 1, 1, false). // Small gap between Tasks and Projects
		AddItem(headerProj, 1, 1, false).
		AddItem(projList, 8, 1, false).

		// 3. THE "SPRING": This fills all empty space
		// Fixed height: 0, Proportion: 1
		AddItem(tview.NewBox(), 0, 1, false).

		// 4. BOTTOM SECTION: Settings
		// Fixed height: 5 (adjust based on number of items), Proportion: 0
		AddItem(settings, 5, 0, false)

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
