package ui

import (
	"fmt"
	"go-tui/db"

	"github.com/rivo/tview"
)

func CreateTaskDetailForm(state *UIState, task db.Task, onComplete func()) tview.Primitive {
	form := tview.NewForm()

	// Use the struct fields directly
	form.AddInputField("Title", task.Title, 30, nil, func(text string) {
		task.Title = text
	})
	form.AddTextArea("Description", task.Desc, 40, 5, 0, func(text string) {
		task.Desc = text
	})

	// Dropdown for Priority
	priorities := []string{"High", "Medium", "Low"}
	initialIndex := 1 // Default to Medium
	form.AddDropDown("Priority", priorities, initialIndex, func(option string, optionIndex int) {
		task.Priority = option
	})

	form.AddButton("Save", func() {
		state.DB.Save(&task)

		onComplete()
	})

	// form.AddButton("Delete")

	form.SetBorder(true).SetTitle(" Task ID: " + fmt.Sprint(task.ID))
	return form
}
