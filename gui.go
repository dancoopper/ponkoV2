package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strings"
	"time"
)

func MainWin() {
	a := app.New()
	w := a.NewWindow("PONKO Part 2")
	w2 := a.NewWindow("Impatient bitch")
	w2.Hide()

	//text := widget.NewLabel(Data1.Explanation)
	//text.Wrapping = fyne.TextWrap(4)

	r := strings.NewReplacer(".", "\n")
	explanT := r.Replace(Data1.Explanation)

	list := widget.NewLabel(explanT)
	//exText := widget.NewLabel("GIRL GIVE ME SOME TIME...I'M FUCKING WORKING ON IT DAMN")
	con := container.New(layout.NewCenterLayout(), list)
	w2.SetContent(con)

	title := canvas.NewText(Data1.Title, color.White)
	tCent := container.New(layout.NewCenterLayout(), title)

	explan := widget.NewButton("Info", w2.Show)
	eCent := container.New(layout.NewCenterLayout(), explan)

	date := "nasa" + time.DateOnly + ".jpg"
	image := canvas.NewImageFromFile("pic/" + date)
	image.FillMode = canvas.ImageFillOriginal

	content := container.NewBorder(tCent, eCent, nil, nil, image)
	w.SetContent(content)
	w.SetOnClosed(w.Close)
	w.ShowAndRun()

}
