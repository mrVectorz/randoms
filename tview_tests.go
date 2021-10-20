package main

import (
//	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
  app := tview.NewApplication()
	menu := tview.NewFlex()
	menuList := tview.NewList().
		ShowSecondaryText(false)
	menu.AddItem(menuList, 0, 2, true)

	menuList.AddItem("Start", "", 's', nil)
	menuList.AddItem("Stop", "", 't', nil).
		AddItem("Flows", "", 'f', nil).
		AddItem("Logs", "", 'l', nil).
		AddItem("Add/Remove Fields from aggregate", "", 'a', nil).
		AddItem("Sort by ", "", 'S', nil)

	if err := app.SetRoot(menu, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
