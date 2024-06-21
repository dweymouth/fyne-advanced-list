package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	fyneadvancedlist "github.com/dweymouth/fyne-advanced-list"
)

func main() {
	a := app.NewWithID("test")
	w := a.NewWindow("win")

	data := []string{}
	for i := 0; i < 1000; i++ {
		data = append(data, fmt.Sprintf("Test list row %d", i))
	}

	l := fyneadvancedlist.NewList(
		func() int { return len(data) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(lii fyneadvancedlist.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(data[lii])
		},
	)
	var selected fyneadvancedlist.ListItemID
	l.OnSelected = func(id fyneadvancedlist.ListItemID) { selected = id }
	l.OnReorderSelectionTo = func(insertAt fyneadvancedlist.ListItemID) {
		newData := make([]string, 0, len(data))
		newData = append(newData, data[:insertAt]...)
		newData = append(newData, data[selected])
		for i := insertAt; i < len(data); i++ {
			if i != selected {
				newData = append(newData, data[i])
			}
		}
		data = newData
		l.UnselectAll()
		l.Refresh()
	}

	w.SetContent(container.NewBorder(
		container.NewStack(
			canvas.NewRectangle(color.RGBA{R: 128, A: 255}), widget.NewLabel("Pad")),
		container.NewStack(
			canvas.NewRectangle(color.RGBA{R: 128, A: 255}), widget.NewLabel("Pad")),
		nil, nil, l),
	)
	w.Resize(fyne.NewSize(300, 400))
	w.ShowAndRun()
}
