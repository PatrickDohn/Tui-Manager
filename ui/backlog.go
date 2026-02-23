package ui

import (
	"go-tui/db"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CreateBacklogPage(state *UIState) tview.Primitive {
	table := tview.NewTable().
		SetBorders(true).
		SetFixed(1, 0).
		SetSelectable(true, false).
		SetEvaluateAllRows(true)

	var tasks []db.Task
	// 1. Fetch the tasks

	table.Clear()
	// Headers
	table.SetCell(0, 0, tview.NewTableCell("Task").SetAttributes(tcell.AttrBold))
	table.SetCell(0, 1, tview.NewTableCell("Priority").SetAttributes(tcell.AttrBold))
	table.SetCell(0, 2, tview.NewTableCell("Due Date").SetAttributes(tcell.AttrBold))
	table.SetCell(0, 3, tview.NewTableCell("Status").SetAttributes(tcell.AttrDim))

	state.DB.
		Find(&tasks)

	for i, t := range tasks {
		dateDisplay := ""
		if !t.DueDate.IsZero() {
			dateDisplay = t.DueDate.Format("01-02-2006")
		}

		table.SetCell(i+1, 0, tview.NewTableCell(t.Title))
		table.SetCell(i+1, 1, tview.NewTableCell(t.Priority).SetTextColor(DraculaGreen))
		table.SetCell(i+1, 2, tview.NewTableCell(dateDisplay))
		table.SetCell(i+1, 3, tview.NewTableCell(t.Status))
	}

	mainContentContainer := tview.NewFlex().SetDirection(tview.FlexRow)

	mainContentContainer.
		AddItem(table, 0, 1, true).
		SetBorder(true).
		SetTitle("Hello")

	return tview.NewFlex().
		AddItem(mainContentContainer, 0, 1, false)

}
