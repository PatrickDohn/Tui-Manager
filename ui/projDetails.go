package ui

import (
	"fmt"
	"go-tui/db"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	DraculaBg      = tcell.NewRGBColor(40, 42, 54)    // #282A36
	DraculaCurrent = tcell.NewRGBColor(68, 71, 90)    // #44475A
	DraculaFg      = tcell.NewRGBColor(248, 248, 242) // #F8F8F2
	DraculaComment = tcell.NewRGBColor(98, 114, 164)  // #6272A4
	DraculaCyan    = tcell.NewRGBColor(139, 233, 253) // #8BE9FD
	DraculaGreen   = tcell.NewRGBColor(80, 250, 123)  // #50FA7B
	DraculaOrange  = tcell.NewRGBColor(255, 184, 108) // #FFB86C
	DraculaPink    = tcell.NewRGBColor(255, 121, 198) // #FF79C6
	DraculaPurple  = tcell.NewRGBColor(189, 147, 249) // #BD93F9
	DraculaRed     = tcell.NewRGBColor(255, 85, 85)   // #FF5555
	DraculaYellow  = tcell.NewRGBColor(241, 250, 140) // #F1FA8C
)

func ProjectDetailForm(state *UIState, project db.Project, onComplete func()) tview.Primitive {
	form := tview.NewForm()
	// form.SetLabelColor(tcell.ColorDarkRed).SetFieldBackgroundColor(tcell.ColorWhite).SetFieldTextColor(tcell.ColorBlack)
	form.SetLabelColor(DraculaGreen)
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
