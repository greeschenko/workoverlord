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
	"greeschenko/workoverlord2/internal/models"
	"log"
)

// GUI — стандартна реалізація GUI
type GUI struct {
	App       fyne.App
	container *CellWidgetContainer
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

		if woapp.Storage.Load() != nil {
			dialog.ShowInformation("Error", "Wrong password", w)
		} else {
			g.showData(w)
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

func (g *GUI) showData(w fyne.Window) {
	g.container = NewCellWidgetContainer(g.RecurceAddGuiCells())

	addbtn := widget.NewButton("ADD", func() {
		fmt.Println("add btn click")
		//IsCreateSelect = true
	})
	deletebtn := widget.NewButton("DELETE", func() {
		fmt.Println("delete btn click")
		//		if len(SELECTED) == 0 {
		//			fmt.Println("no cells selected")
		//		} else {
		//			for _, v := range SELECTED {
		//				delete(USERMIND.Cells, v)
		//			}
		//			saveData()
		//            woapp.Storage.Save()
		//			g.container.Container.Objects = RecurceAddGuiCells()
		//			g.container.Refresh()
		//		}
	})
	closebtn := widget.NewButton("CLOSE", func() {
		w.Close()
	})
	mainmenu := container.NewHBox(addbtn, deletebtn, closebtn)

	content := container.NewBorder(mainmenu, nil, nil, nil, g.container)
	w.SetContent(content)
}

func (g *GUI) RecurceAddGuiCells() []fyne.CanvasObject {
	var celllist []fyne.CanvasObject
	woapp := woapp.GetInstance()
	for i,e := range woapp.Storage.GetData().Cells {
		if e.Status == models.CellStatusConfig {
			continue
		}
		myw := NewCellWidget(i, e)

		celllist = append(celllist, myw)
	}

	return celllist
}
