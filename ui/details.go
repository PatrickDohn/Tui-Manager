package ui

import (
	"fmt"
	"go-tui/db"
	"time"

	"github.com/rivo/tview"
)

func CreateTaskDetailForm(state *UIState, task db.Task, onComplete func(), onDelete func()) tview.Primitive {

	// var dueDateText string
	dueDateText := task.DueDate.Format("01-02-2006")
	form := tview.NewForm()

	// Use the struct fields directly
	form.AddInputField("Title", task.Title, 30, nil, func(text string) {
		task.Title = text
	})

	form.AddInputField("Due Date", dueDateText, 30, nil, func(text string) {
		dueDateText = text
	})

	form.AddTextArea("Description", task.Desc, 60, 30, 0, func(text string) {
		task.Desc = text
	})

	// Dropdown for Priority
	priorities := []string{"High", "Medium", "Low"}
	initialIndex := 1 // Default to Medium
	form.AddDropDown("Priority", priorities, initialIndex, func(option string, optionIndex int) {
		task.Priority = option
	})

	// Dropdown for status'
	status := []string{"Pending", "In Progress", "Done"}

	form.AddDropDown("Status", status, 0, func(opt string, optIndex int) {
		task.Status = opt
	})

	form.AddButton("Save", func() {

		t, err := time.Parse("01-02-2006", dueDateText)
		if err != nil {
			// show an error to the user however you like

			onComplete()
		}
		task.DueDate = t.UTC()
		state.DB.Save(&task)
		onComplete()
	})

	form.AddButton("Delete", func() {
		state.DB.Delete(&task, task.ID)
		onDelete()
	})
	message := fmt.Sprintf(" %s Details", task.Title)
	form.SetBorder(true).SetTitle(message)
	return form
}
