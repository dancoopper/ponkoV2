package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func MainWin() {
	a := app.New()
	w := a.NewWindow("PONKO Part 2")
	w2 := a.NewWindow("Impatient bitch")
	w2.Hide()

	exText := widget.NewLabel("GIRL GIVE ME SOME TIME...I'M FUCKING WORKING ON IT DAMN")
	w2.SetContent(exText)

	/*
		exText := widget.NewLabel(Data1.Explanation)
		exText.Wrapping = fyne.TextWrapWord
		w2.SetContent(exText)

	*/

	title := canvas.NewText(Data1.Title, color.White)
	tCent := container.New(layout.NewCenterLayout(), title)

	explan := widget.NewButton("Info", w2.Show)
	eCent := container.New(layout.NewCenterLayout(), explan)

	//image := canvas.NewImageFromResource(theme.FyneLogo())
	//image := canvas.NewImageFromURI(uri)
	// image := canvas.NewImageFromImage(main)
	// image := canvas.NewImageFromReader(reader, name)
	image := canvas.NewImageFromFile("pic/nasapic.jpg")
	image.FillMode = canvas.ImageFillOriginal

	content := container.New(layout.NewVBoxLayout(), tCent, image, eCent)
	w.SetContent(content)
	w.SetOnClosed(w.Close)
	w.ShowAndRun()

}
