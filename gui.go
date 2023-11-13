package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strings"
)

func MainWin() {
	a := app.New()
	w := a.NewWindow("PONKO Part 2")
	w2 := a.NewWindow("Impatient bitch")
	w2.Hide()
	changeLogWin := a.NewWindow("Change Log")
	changeLogWin.Hide()

	/*
		r := strings.NewReplacer(".", "\n")
		explanT := r.Replace(Data1.Explanation)

		list := widget.NewLabel(explanT)
	*/
	//list := widget.NewLabel("MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM") //40 chars long
	exText := widget.NewLabel("Still working on it but also have a few other Ideas so idk when this is going to get done")
	con := container.New(layout.NewCenterLayout(), exText)
	w2.SetContent(con)

	changeLogText := widget.NewLabel("- Added Change log button\n- Added Credit to the photographer\n- Added Save button\n- Added Update on how info window is coming along\n- Fixed bug where date would be wrong")
	changeLogCon := container.New(layout.NewCenterLayout(), changeLogText)
	changeLogWin.SetContent(changeLogCon)

	toolBar := widget.NewToolbar(widget.NewToolbarAction(theme.DocumentIcon(), changeLogWin.Show))

	cleanedCopyright := strings.Replace(Data1.CopyRight, "\n", "", -1)
	titleContens := Data1.Title + " (" + cleanedCopyright + ")"
	title := canvas.NewText(titleContens, color.White)
	TitleCon := container.New(layout.NewCenterLayout(), title)
	tCent := container.NewBorder(nil, TitleCon, toolBar, nil)

	explan := widget.NewButton("Info", w2.Show)
	saveBtn := widget.NewButton("Save", SavePic)
	contanerB := container.NewBorder(nil, nil, saveBtn, explan)
	eCent := container.New(layout.NewCenterLayout(), contanerB)

	date := "nasa" + Data1.Date + ".jpg"
	image := canvas.NewImageFromFile("pic/" + date)
	image.FillMode = canvas.ImageFillOriginal

	content := container.NewBorder(tCent, eCent, nil, nil, image)
	w.SetContent(content)
	w2.SetOnClosed(w2.Close)
	changeLogWin.SetOnClosed(changeLogWin.Close)
	w.SetOnClosed(w.Close)
	w.SetOnClosed(CleanUp)
	w.ShowAndRun()

}
