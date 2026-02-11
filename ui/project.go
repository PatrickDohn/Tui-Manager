package ui

import (
	"fmt"
	"go-tui/db"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CreateProjectDetailPage(state *UIState) tview.Primitive {

	table := tview.NewTable().
		SetBorders(true).
		SetSelectable(true, false).
		// This ensures the selection stays within existing cells
		SetEvaluateAllRows(true)

	mainContentContainer := tview.NewFlex().SetDirection(tview.FlexRow)
	// projDetailContainer := tview.NewFlex().SetDirection(tview.FlexRow)

	var tasks []db.Task
	// 1. Fetch the tasks
	refreshTable := func() {
		table.Clear()
		// Headers
		table.SetCell(0, 0, tview.NewTableCell("Task").SetAttributes(tcell.AttrBold))
		table.SetCell(0, 1, tview.NewTableCell("Priority").SetAttributes(tcell.AttrBold))
		table.SetCell(0, 2, tview.NewTableCell("Due Date").SetAttributes(tcell.AttrBold))
		table.SetCell(0, 3, tview.NewTableCell("Status").SetAttributes(tcell.AttrDim))

		state.DB.
			Where("project_id = ?", state.CurrentProject.ID).
			Find(&tasks)

		for i, t := range tasks {
			dateDisplay := ""
			if !t.DueDate.IsZero() {
				dateDisplay = t.DueDate.Format("01/02/2026")
			}

			table.SetCell(i+1, 0, tview.NewTableCell(t.Title))
			table.SetCell(i+1, 1, tview.NewTableCell(t.Priority))
			table.SetCell(i+1, 2, tview.NewTableCell(dateDisplay))
			table.SetCell(i+1, 3, tview.NewTableCell(t.Status))
		}

	}
	quickinput := tview.NewInputField().
		SetLabel(" [green]+[white] New Project Task: ").
		SetFieldWidth(0).
		SetPlaceholder("Type task title and press enter...")

	quickinput.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			title := quickinput.GetText()
			if title == "" {
				return
			}

			newTask := db.Task{
				Title:     title,
				Status:    "Pending",
				Priority:  "",
				Desc:      "",
				ProjectID: &state.CurrentProject.ID,
			}
			state.DB.Create(&newTask)

			quickinput.SetText("")
			refreshTable()
		}
	})

	var message string
	var defaultProjView tview.Primitive

	if state.CurrentProject == nil {
		message = " No Project Selected "
	} else {
		message = fmt.Sprintf(" %s Tasks", state.CurrentProject.Name)
		defaultProjView = ProjectDetailForm(state, *state.CurrentProject, func() {
			state.RefreshSidebar()
			state.App.SetFocus(table)
		})
	}

	detailContainer := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(defaultProjView, 0, 1, true)

	renderForm := func(task db.Task) {
		detailContainer.Clear()

		form := CreateTaskDetailForm(state, task, func() {
			refreshTable()

			// Restore default view
			detailContainer.Clear()
			detailContainer.AddItem(defaultProjView, 0, 1, true)

			state.App.SetFocus(table)
		})

		detailContainer.AddItem(form, 0, 1, true)
		state.App.SetFocus(form)
	}

	// renderProjForm := func(proj db.Project) {
	// 	projDetailContainer.Clear()
	// 	form := ProjectDetailForm(state, *state.CurrentProject, func() {
	// 		state.App.SetFocus(table)
	// 	})
	// 	projDetailContainer.AddItem(form, 0, 1, true)
	// }

	// 4. Update on Selection
	table.SetSelectionChangedFunc(func(r, c int) {
		if r <= 0 || r > len(tasks) {
			return
		}
		renderForm(tasks[r-1]) // Pass the actual task struct
	})

	mainContentContainer.
		AddItem(table, 0, 1, true).
		AddItem(quickinput, 3, 1, false).
		SetBorder(true).
		SetTitle(message).
		SetFocusFunc(func() {
			detailContainer.Clear()
			detailContainer.AddItem(defaultProjView, 0, 1, true)
		})

	if state.CurrentProject != nil {
		refreshTable() // populate table immediately
	} else {
		table.Clear() // safe fallback if no project selected

	}

	return tview.NewFlex().
		AddItem(mainContentContainer, 0, 1, true).
		AddItem(detailContainer, 0, 1, false)

}
