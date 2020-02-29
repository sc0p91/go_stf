package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func formApp(app fyne.App) {

	w := app.NewWindow("FW Finder")

	hname := widget.NewEntry()
	hname.SetPlaceHolder("vnixeh")
	rsapin := widget.NewPasswordEntry()
	code := widget.NewEntry()
	code.SetPlaceHolder("mobile token")

	form := &widget.Form{
		OnCancel: func() {
			w.Close()
		},
		OnSubmit: func() {
			fmt.Println("Form submitted")
			fmt.Println("Hostname:", hname.Text)
			fmt.Println("RSA PIN:", rsapin.Text)
			fmt.Println("Token:", code.Text)
		},
	}
	form.Append("Hostname", hname)
	form.Append("RSA PIN", rsapin)
	form.Append("Token", code)
	w.SetContent(form)
	w.Show()
}

func main() {
	app := app.New()

	w := app.NewWindow("Menu")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Firewall Toolkit"),
		widget.NewButton("Search Fw", func() { formApp(app) }),
		widget.NewButton("Open fwehap", func() { app.Quit() }),
		widget.NewButton("Open zfehap", func() { app.Quit() }),
		widget.NewButton("Quit", func() { app.Quit() }),
	))

	w.ShowAndRun()
}
