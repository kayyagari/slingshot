package ui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"slingshot/jnlp"
)

type ConnectionDialog struct {
	a fyne.App
	w fyne.Window
}

func NewConnectionDialog() *ConnectionDialog {
	cd := &ConnectionDialog{}
	a := app.New()
	cd.a = a
	w := a.NewWindow("Slingshot")
	cd.w = w

	urlEntry := widget.NewEntry()
	urlBox := widget.NewHBox(widget.NewLabel("URL:"), urlEntry)
	w.SetContent(widget.NewVBox(
		urlBox,
		widget.NewButton("Run", func() {
			dirPath, err := jnlp.ParseAndDownload(urlEntry.Text)
			fmt.Println(dirPath, err)
		})))

	return cd
}

func (cd *ConnectionDialog) Show() {
	cd.w.ShowAndRun()
}
