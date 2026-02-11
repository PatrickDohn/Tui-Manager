package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

func CreateProjectDetailPage(state *UIState) tview.Primitive {

	mainContentContainer := tview.NewFlex().
		SetDirection(tview.FlexRow)

	quickinput := tview.NewInputField().
		SetLabel(" [green]+[white] New Task: ").
		SetFieldWidth(0).
		SetPlaceholder("Type task title and press enter...")

	var message string

	if state.CurrentProject == nil {
		message = " No Project Selected "
	} else {
		message = fmt.Sprintf(" %s ", state.CurrentProject.Name)
	}

	mainContentContainer.
		AddItem(quickinput, 3, 1, false).
		SetBorder(true).
		SetTitle(message)

	return tview.NewFlex().
		AddItem(mainContentContainer, 0, 1, true)
}
