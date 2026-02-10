package ui

import "github.com/rivo/tview"

func CreateSettingsPage(state *UIState) tview.Primitive {
	form := tview.NewForm()

	form.AddInputField("Display Name", state.UserName, 20, nil, func(text string) {
		state.UserName = text // Updates state in real-time
	})

	form.AddButton("Back to Home", func() {
		state.MainPages.SwitchToPage("home")
	})

	form.SetBorder(true).SetTitle(" System Settings ⛭ ")
	return form
}
