package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	woapp "greeschenko/workoverlord2/internal/app"
	"log"
)

// GUI — стандартна реалізація GUI
type GUI struct {
	App fyne.App
}

func NewFyneGUI() *GUI {
	return &GUI{
		App: app.New(),
	}
}

func (g *GUI) Start() {
	woapp := woapp.GetInstance()
	w := g.App.NewWindow("WorkOverlord")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter your password")

	form := widget.NewForm(widget.NewFormItem("Password", passwordEntry))

	form.OnSubmit = func() {
		woapp.Storage.SetSecret(passwordEntry.Text)

		fmt.Println("secret is ", woapp.Storage.GetSecret())

		if initDb() != nil {
			dialog.ShowInformation("Error", "Wrong password", w)
		} else {
			initGui(w)
		}
	}

	passwordEntry.OnSubmitted = func(_ string) {
		form.OnSubmit()
	}

	content := container.NewVBox(
		widget.NewLabel("Please enter your password:"),
		form,
	)

	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl}
	w.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Tab")
	})

	w.SetContent(content)

	w.Canvas().Focus(passwordEntry)

	w.Resize(fyne.NewSize(1200, 600))
	w.ShowAndRun()
}
