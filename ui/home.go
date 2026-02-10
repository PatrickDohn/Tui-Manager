package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Task struct {
	ID       int
	Title    string
	Status   string // e.g., "Pending", "Done"
	Priority string // e.g., "High", "Low"
	Desc     string
	DueDate  string
	// Using a pointer makes this field optional (nullable)
	ProjectID *uint // Foreign Key linking to Project
}

func CreateHomePage(state *UIState) tview.Primitive {
	table := tview.NewTable().
		SetBorders(true).
		SetSelectable(true, false).
		// This ensures the selection stays within existing cells
		SetEvaluateAllRows(true)
	detailContainer := tview.NewFlex().SetDirection(tview.FlexRow)
	mainContentContainer := tview.NewFlex().SetDirection(tview.FlexRow)

	// 1. Create Dummy Data using your Database Models
	// (In the future, this comes from state.DB.Find(&tasks))
	tasks := []Task{
		{
			ID:       1,
			Title:    "Install Go on M4 Mac",
			Status:   "Done",
			Priority: "High",
			DueDate:  "01/24/2026",
			Desc:     "Download the ARM64 installer from go.dev",
		},
		{
			ID:       2,
			Title:    "Setup SQLite",
			Status:   "Pending",
			Priority: "Medium",
			DueDate:  "05/02/2026",
			Desc:     "Configure GORM with the glebarez driver",
		},
	}

	// 2. Helper to render the Form using the Model
	renderForm := func(task Task) {
		detailContainer.Clear()
		form := CreateTaskDetailForm(state, task, func() {
			// Delete logic: remove from slice and refresh table
			state.App.SetFocus(table)
		})
		detailContainer.AddItem(form, 0, 1, true)
	}

	// 3. Populate Table
	table.SetCell(0, 0, tview.NewTableCell("Task").SetAttributes(tcell.AttrBold))
	table.SetCell(0, 1, tview.NewTableCell("Priority").SetAttributes(tcell.AttrBold))
	table.SetCell(0, 2, tview.NewTableCell("Due Date").SetAttributes(tcell.AttrBold))
	table.SetCell(0, 3, tview.NewTableCell("Status").SetAttributes(tcell.AttrDim))

	for i, t := range tasks {
		table.SetCell(i+1, 0, tview.NewTableCell(t.Title))
		table.SetCell(i+1, 1, tview.NewTableCell(t.Priority))
		table.SetCell(i+1, 2, tview.NewTableCell(t.DueDate))
		table.SetCell(i+1, 3, tview.NewTableCell(t.Status))
	}

	// 4. Update on Selection
	table.SetSelectionChangedFunc(func(r, c int) {
		if r <= 0 || r > len(tasks) {
			return
		}
		renderForm(tasks[r-1]) // Pass the actual task struct
	})

	mainContentContainer.AddItem(table, 0, 1, true).SetBorder(true).SetTitle(" Daily Tasks ")

	return tview.NewFlex().
		AddItem(mainContentContainer, 0, 1, true).
		AddItem(detailContainer, 0, 1, false)
}
