package ui

import "github.com/rivo/tview"

func CreateLoginPage(state *UIState, rootPages *tview.Pages, target tview.Primitive) tview.Primitive {
	form := tview.NewForm()
	var u, p string

	form.AddInputField("Username", "", 20, nil, func(text string) { u = text })
	form.AddPasswordField("Password", "", 20, '*', func(text string) { p = text })

	form.AddButton("Login", func() {
		if u == "admin" && p == "pass" {
			state.UserName = u
			rootPages.SwitchToPage("main_app") // Switch the top-level view
			state.App.SetFocus(target)         // Give focus to the sidebar
		}
	})

	form.SetBorder(true).SetTitle(" Login ")

	// Centering the login form
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(form, 12, 1, true).
			AddItem(nil, 0, 1, false), 40, 1, true).
		AddItem(nil, 0, 1, false)
}
