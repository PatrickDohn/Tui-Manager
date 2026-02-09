package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CreateHomePage(state *UIState) tview.Primitive {

	titleText := "No Project Selected"

	// Check if there is an active project in our state
	if state.CurrentProject != nil {
		titleText = "Project: " + state.CurrentProject.Name
	}

	header := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("\n" + titleText).
		SetTextColor(tcell.ColorYellow)

	table := tview.NewTable().SetBorders(true).SetSelectable(true, false)

	details := tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true)
	details.SetBorder(true).SetTitle(" Task Details ")

	taskPriority := map[int]string{
		1: "red",    // High priority
		2: "green",  // Done
		3: "yellow", // In Progress
		4: "blue",   // planning
	}
	// 1. Define your data in a Map (Key = Task Name, Value = Description)
	taskData := map[string]string{
		"Install Go on M4 Mac": "[cyan]Status:[white] High Priority\n\nEnsure you use the ARM64 version for Apple Silicon.",
		"Connect SQLite DB":    "[red]Status:[white] In Progress\n\nSetting up GORM with the SQLite driver.",
		"Build TUI Dashboard":  fmt.Sprintf("[%s]Status:[white] Planning\n\nDesigning the layout with tview Flex and Pages.", taskPriority[4]),
		"Testing":              fmt.Sprintf("[%s]Status:[white] In Progress\n\nSetting up GORM with the SQLite driver.", taskPriority[3]),
	}

	// 2. Populate the table using the map keys
	table.SetCell(0, 0, tview.NewTableCell("Task Name").SetTextColor(tcell.ColorYellow).SetAttributes(tcell.AttrBold))
	table.SetCell(0, 1, tview.NewTableCell("Status").SetTextColor(tcell.ColorYellow).SetAttributes(tcell.AttrBold))
	table.SetCell(0, 2, tview.NewTableCell("Priority").SetTextColor(tcell.ColorYellow).SetAttributes(tcell.AttrBold))

	// 2. Populate Data Rows
	row := 1
	for name := range taskData {
		// Column 0: Name
		table.SetCell(row, 0, tview.NewTableCell(name))

		// Column 1: Status (Static example or from data)
		table.SetCell(row, 1, tview.NewTableCell("[green]Active"))

		// Column 2: Priority (Static example or from data)
		table.SetCell(row, 2, tview.NewTableCell("High"))

		row++
	}

	// 3. THE FIX: Use taskName to update details
	table.SetSelectedFunc(func(r, c int) {
		if r == 0 {
			return
		}

		// Get the text from the clicked cell
		taskName := table.GetCell(r, 0).Text

		// Look up the description in our map using that name
		if desc, full := taskData[taskName]; full {
			details.SetText("[yellow]Details for: " + taskName + "[white]\n\n" + desc)
		} else {
			details.SetText("No details found for this task.")
		}
	})

	return tview.NewFlex().
		AddItem(header, 3, 1, false).
		AddItem(table, 0, 3, true).
		AddItem(details, 0, 2, false)
}
