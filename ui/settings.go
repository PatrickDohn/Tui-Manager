package ui

import (
	"go-tui/db"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CreateSettingsPage(state *UIState) tview.Primitive {
	var projects []db.Project
	state.DB.Find(&projects)
	form := tview.NewForm()

	projList := tview.NewList().SetCurrentItem(0)

	projList.Clear()

	state.DB.Find(&projects)

	if len(projects) == 0 {
		projList.AddItem(
			"[gray]No projects[-]",
			"",
			0,
			nil,
		)
		projList.SetCurrentItem(0)
	} else {
		for i := range projects {
			proj := &projects[i] // safe pointer to slice element

			form.AddInputField("Project Name", proj.Name, 40, nil, func(text string) {})
			form.AddInputField("Project Desc.", proj.Description, 60, nil, func(text string) {})
			form.AddInputField("Github Project Name", proj.GithubProjName, 60, nil, func(text string) {})
			form.AddInputField("Github User Name", proj.GHUserName, 60, nil, func(text string) {})
			form.AddInputField("Token", proj.GHUserName, 60, nil, func(text string) {})
			spacer := tview.NewTextView().SetText("--- --- --- --- ---").SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorYellow)
			form.AddFormItem(spacer)
		}

	}

	form.AddButton("Add Project", func() {
		form.AddInputField("Project Name", "", 40, nil, func(text string) {})
		form.AddInputField("Project Desc.", "", 60, nil, func(text string) {})
		spacer := tview.NewTextView().SetText("--- --- --- --- ---").SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorYellow)
		form.AddFormItem(spacer)
	})

	form.SetBorder(true).SetTitle(" System Settings ⛭ ")

	return form
}
