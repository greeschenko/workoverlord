package gui

import (
	"fmt"
	"greeschenko/workoverlord2/internal/interfaces"
	"greeschenko/workoverlord2/internal/models"
	"greeschenko/workoverlord2/pkg/kdtreepositioner"
	"image/color"
	"log"
	"os"
	"os/exec"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var GUIZOOM = binding.NewFloat()
var GUIDATAUPDATER = binding.NewInt()
var GUICONTAINER *CellWidgetContainer

var FONTSIZE = 14

var IsCreateSelect = false

var SELECTED []*CellWidget

// var COLORBG = color.NRGBA{R: 0x28, G: 0x2c, B: 0x34, A: 0xff}
var COLORBG = color.NRGBA{R: 40, G: 44, B: 52, A: 0x00}

// var COLORTXT = color.NRGBA{R: 0xff, G: 0xb7, B: 0xce, A: 0xff}
var COLORTXT = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
var COLORLINES = color.NRGBA{R: 250, G: 67, B: 114, A: 0xff}
var COLORBRD = color.NRGBA{R: 40, G: 44, B: 52, A: 0xff}
var COLORSTR = color.NRGBA{R: 0x5f, G: 0x9e, B: 0xa0, A: 0xff}

// GUI — стандартна реалізація GUI
type GUI struct {
	App        fyne.App
	container  *CellWidgetContainer
	Data       interfaces.DataInterface
	Positioner kdtreepositioner.KDTree
}

func NewFyneGUI(Data interfaces.DataInterface) *GUI {
	return &GUI{
		App:  app.New(),
		Data: Data,
	}
}

func (g *GUI) Start() {
	w := g.App.NewWindow("WorkOverlord")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter your password")

	form := widget.NewForm(widget.NewFormItem("Password", passwordEntry))

	form.OnSubmit = func() {
		g.Data.SetSecret(passwordEntry.Text)

		if g.Data.Load() != nil {
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

	w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		switch ev.Name {
		case fyne.KeyO:
			if len(SELECTED) > 0 {
                item := SELECTED[0]
				err := g.UpdateCell(item.ID())
				if err != nil {
					panic(err)
				} else {
					item.genText()
					item.Refresh()
					ZoomRefresh()
				}
			} else {
				log.Println("no selected element")
			}
		case fyne.KeyK:
			if len(SELECTED) > 0 {
				cur := SELECTED[0]
				SELECTED[0].SetSelected(false)
				g.Positioner.FindNearestInDirection(cur, "up").SetSelected(true)
				SELECTED[0].CenterInWindow()
			} else {
				log.Println("no selected element")
			}
		case fyne.KeyJ:
			if len(SELECTED) > 0 {
				cur := SELECTED[0]
				SELECTED[0].SetSelected(false)
				g.Positioner.FindNearestInDirection(cur, "down").SetSelected(true)
				SELECTED[0].CenterInWindow()
			} else {
				log.Println("no selected element")
			}
		case fyne.KeyH:
			if len(SELECTED) > 0 {
				cur := SELECTED[0]
				SELECTED[0].SetSelected(false)
				g.Positioner.FindNearestInDirection(cur, "left").SetSelected(true)
				SELECTED[0].CenterInWindow()
			} else {
				log.Println("no selected element")
			}
		case fyne.KeyL:
			if len(SELECTED) > 0 {
				cur := SELECTED[0]
				SELECTED[0].SetSelected(false)
				g.Positioner.FindNearestInDirection(cur, "right").SetSelected(true)
				SELECTED[0].CenterInWindow()
			} else {
				log.Println("no selected element")
			}
		case fyne.KeyMinus:
			zoom, _ := GUIZOOM.Get()
			newZoom := zoom
			if zoom > 0.1 {
				newZoom = zoom - 0.1
				GUIZOOM.Set(newZoom)
			} else {
				GUIZOOM.Set(0.1)
			}
		case fyne.KeyEqual:
			zoom, _ := GUIZOOM.Get()
			newZoom := zoom
			if zoom < 1 {
				newZoom = zoom + 0.1
				GUIZOOM.Set(newZoom)
			} else {
				GUIZOOM.Set(1)
			}
		}
	})

	w.SetContent(content)

	w.Canvas().Focus(passwordEntry)

	w.Resize(fyne.NewSize(1200, 600))
	w.ShowAndRun()
}

func (g *GUI) showData(w fyne.Window) {
	guicells := g.RecurceAddGuiCells()

	var canvasObjects []fyne.CanvasObject
	for _, cell := range guicells {
		canvasObjects = append(canvasObjects, cell)
	}
	g.container = NewCellWidgetContainer(canvasObjects, g)

	addbtn := widget.NewButton("ADD", func() {
		IsCreateSelect = true
	})
	deletebtn := widget.NewButton("DELETE", func() {
		fmt.Println("delete btn click")
		if len(SELECTED) == 0 {
			fmt.Println("no cells selected")
		} else {
			for _, v := range SELECTED {
				g.Data.Delete(v.ID())
			}

			guicells := g.RecurceAddGuiCells()

			var canvasObjects []fyne.CanvasObject
			for _, cell := range guicells {
				canvasObjects = append(canvasObjects, cell)
			}
			g.container.Container.Objects = canvasObjects
			g.container.Refresh()
		}
	})
	closebtn := widget.NewButton("CLOSE", func() {
		w.Close()
	})
	mainmenu := container.NewHBox(addbtn, deletebtn, closebtn)

	content := container.NewBorder(mainmenu, nil, nil, nil, g.container)
	w.SetContent(content)

	objects := make([]kdtreepositioner.SpatialObject, len(guicells))
	for i, obj := range guicells {
		objects[i] = obj
	}

	g.Positioner = kdtreepositioner.NewKDTree(objects, 0)
	g.Positioner.NearestNeighbor([2]int{1000, 500}).SetSelected(true)
}

func (g *GUI) RecurceAddGuiCells() []*CellWidget {
	var celllist []*CellWidget
	for i, e := range g.Data.GetAll() {
		if *e.Status == models.CellStatusConfig {
			continue
		}
		myw := NewCellWidget(i, e, g)

		celllist = append(celllist, myw)
	}

	return celllist
}

func (g *GUI) AddCell(point [2]int) (string, error) {
	newkey := time.Now().Format(time.RFC3339)
	return newkey, g.editContent(newkey, "", &point)
}

func (g *GUI) UpdateCell(key string) error {
	text, err := g.Data.GetOne(key)
	if err != nil {
		return fmt.Errorf("text with key '%s' not found: %s", key, err)
	}
	return g.editContent(key, text.Content, nil)
}

// editText handles the editing of a text by key
func (g *GUI) editContent(key string, existingContent string, point *[2]int) error {
	// Create a temporary file to store the input text
	tmpfile, err := os.CreateTemp("", "temp*.txt")
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up after use

	// Write existing content to the temporary file if available
	if existingContent != "" {
		if err := os.WriteFile(tmpfile.Name(), []byte(existingContent), 0644); err != nil {
			return fmt.Errorf("failed to write existing content to temporary file: %v", err)
		}
	}

	// Detect the terminal type using $TERM
	term := os.Getenv("TERM")
	cmd := prepareEditorCommand(term, tmpfile.Name())

	// Start the editor process
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to open editor in new terminal: %v", err)
	}

	// Wait for the Vim process to finish
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("editor did not close properly: %v", err)
	}

	// Read the content from the temporary file after editing
	content, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		return fmt.Errorf("failed to read from temporary file: %v", err)
	}

	if existingContent == "" {
		newstatus := models.CellStatusActive
		_, err := g.Data.Add(
			key,
			models.Cell{
				Content:  string(content),
				Position: point,
				Status:   &newstatus,
			},
		)
		if err != nil {
			return fmt.Errorf("failed to add cell to data: %v", err)
		}
	} else {
		_, err := g.Data.Patch(
			key,
			models.Cell{
				Content: string(content),
			},
		)
		if err != nil {
			return fmt.Errorf("failed to update cell to data: %v", err)
		}
	}
	//TODO change to load
	//saveData()
	return nil
}

// prepareEditorCommand prepares the command to open the editor based on $TERM
func prepareEditorCommand(term string, filePath string) *exec.Cmd {
	var cmd *exec.Cmd
	switch term {
	case "xterm", "xterm-256color", "screen", "st", "konsole":
		cmd = exec.Command(term, "-e", "vim", filePath)
	case "gnome-terminal":
		cmd = exec.Command("gnome-terminal", "--", "vim", filePath)
	default:
		cmd = exec.Command("xterm", "-e", "vim", filePath)
	}
	return cmd
}
