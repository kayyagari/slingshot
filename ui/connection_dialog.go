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
	urlEntry.SetText("http://localhost:8080/webstart.jnlp")

	usernameEntry := widget.NewEntry()
	usernameEntry.SetText("admin")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetText("admin")

	fiUrl := &widget.FormItem{"URL:", urlEntry}
	fiUsername := &widget.FormItem{"Username:", usernameEntry}
	fiPassword := &widget.FormItem{"Password:", passwordEntry}

	form := widget.NewForm(fiUrl, fiUsername, fiPassword)
	w.SetContent(widget.NewVBox(
		form,
		widget.NewButton("Run", func() {
			dirPath, j, err := jnlp.ParseAndDownload(urlEntry.Text)
			fmt.Println(dirPath, err)
			args := j.ApplicationDesc.Argument
			if usernameEntry.Text != "" {
				args = append(args, usernameEntry.Text)
				args = append(args, passwordEntry.Text)
			}
			jnlp.RunJar(dirPath, j.ApplicationDesc.MainClass, nil, args)
		})))

	return cd
}

func (cd *ConnectionDialog) Show() {
	cd.w.ShowAndRun()
}
