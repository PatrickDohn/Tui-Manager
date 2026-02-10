package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

func CreateTaskDetailForm(state *UIState, task Task, onDelete func()) tview.Primitive {
	form := tview.NewForm()

	// Use the struct fields directly
	form.AddInputField("Title", task.Title, 30, nil, nil)
	form.AddTextArea("Description", task.Desc, 40, 5, 0, nil)

	// Dropdown for Priority
	priorities := []string{"High", "Medium", "Low"}
	initialIndex := 1 // Default to Medium
	form.AddDropDown("Priority", priorities, initialIndex, nil)

	form.AddButton("Save", func() {
		// Later: state.DB.Save(&task)
	})

	form.AddButton("Delete", onDelete)

	form.SetBorder(true).SetTitle(" Task ID: " + fmt.Sprint(task.ID))
	return form
}
