package ui

import (
	"fmt"
	"go-tui/db"

	"github.com/rivo/tview"
)

func ProjectDetailForm(state *UIState, project db.Project, onComplete func()) tview.Primitive {
	form := tview.NewForm()
	// form.SetLabelColor(tcell.ColorDarkRed).SetFieldBackgroundColor(tcell.ColorWhite).SetFieldTextColor(tcell.ColorBlack).SetF
	// Use the struct fields directly
	form.AddInputField("Title", project.Name, 30, nil, func(text string) {
		project.Name = text
	})
	form.AddInputField("Description", project.Description, 30, nil, func(text string) {
		project.Description = text
	})

	form.AddTextArea("Notes", project.Notes, 60, 30, 0, func(text string) {
		project.Notes = text
	})

	form.AddButton("Save", func() {
		state.DB.Save(&project)

		onComplete()
	})

	// form.AddButton("Delete")
	message := fmt.Sprintf(" %s Details", state.CurrentProject.Name)
	form.SetBorder(true).SetTitle(message)
	return form
}
