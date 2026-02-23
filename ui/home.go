package ui

import (
	"go-tui/db"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CreateHomePage(state *UIState) tview.Primitive {
	table := tview.NewTable().
		SetBorders(true).
		SetSelectable(true, false).
		// This ensures the selection stays within existing cells
		SetEvaluateAllRows(true)
	detailContainer := tview.NewFlex().SetDirection(tview.FlexRow)
	mainContentContainer := tview.NewFlex().SetDirection(tview.FlexRow)

	var tasks []db.Task
	// 1. Fetch the tasks
	refreshTable := func() {
		table.Clear()
		// Headers
		table.SetCell(0, 0, tview.NewTableCell("Task").SetAttributes(tcell.AttrBold))
		table.SetCell(0, 1, tview.NewTableCell("Priority").SetAttributes(tcell.AttrBold))
		table.SetCell(0, 2, tview.NewTableCell("Due Date").SetAttributes(tcell.AttrBold))
		table.SetCell(0, 3, tview.NewTableCell("Status").SetAttributes(tcell.AttrDim))

		// SQLite’s date() function converts timestamps to UTC internally unless told otherwise.
		// value in db: 2026-02-12 21:13:53.618046-05:00
		// value if queried like WHERE(date(due_date)) => returns 2026-02-13 02:13:53 UTC
		//
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		tomorrow := today.Add(24 * time.Hour)

		// fmt.Println("todays date:", today)
		// fmt.Println("Tomorows date: ", tomorrow)
		// fmt.Println("DATES: ", today > tomorrow)

		state.DB.
			Where("due_date >= ? AND due_date < ?", today, tomorrow).
			Find(&tasks)

		for i, t := range tasks {
			dateDisplay := ""
			if !t.DueDate.IsZero() {
				dateDisplay = t.DueDate.Format("01-02-2006")
			}

			table.SetCell(i+1, 0, tview.NewTableCell(t.Title))
			table.SetCell(i+1, 1, tview.NewTableCell(t.Priority))
			table.SetCell(i+1, 2, tview.NewTableCell(dateDisplay))
			table.SetCell(i+1, 3, tview.NewTableCell(t.Status))
		}

	}

	// <-- QUICK ADD INPUT -->
	quickinput := tview.NewInputField().
		SetLabel(" [green]+[white] New Task: ").
		SetFieldWidth(0).
		SetPlaceholder("Type task title and press enter...").
		SetPlaceholderTextColor(DraculaComment)

	quickinput.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			title := quickinput.GetText()
			if title == "" {
				return
			}
			now := time.Now()

			newTask := db.Task{
				Title:    title,
				Status:   "Pending",
				DueDate:  now,
				Priority: "",
				Desc:     "",
			}
			state.DB.Create(&newTask)

			quickinput.SetText("")
			refreshTable()
		}
	})

	// 2. Helper to render the Form using the Model
	renderForm := func(task db.Task) {
		detailContainer.Clear()
		form := CreateTaskDetailForm(state, task, func() {
			// Delete logic: remove from slice and refresh table
			refreshTable()
			state.App.SetFocus(table)
		})
		detailContainer.AddItem(form, 0, 1, true)
	}

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
		SetTitle(" Daily Tasks ")

	refreshTable()

	return tview.NewFlex().
		AddItem(mainContentContainer, 0, 1, true).
		AddItem(detailContainer, 0, 1, false)
}
