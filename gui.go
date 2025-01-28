package main

import (
	"crypto/sha256"
	"fmt"
	"image/color"
	"log"
	"strings"
	"unicode/utf8"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	//"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var GUIZOOM = binding.NewFloat()
var GUIDATAUPDATER = binding.NewInt()
var GUICONTAINER *CellWidgetContainer

var FONTSIZE = 14

var IsCreateSelect = false

var SELECTED = []string{}

// var COLORBG = color.NRGBA{R: 0x28, G: 0x2c, B: 0x34, A: 0xff}
var COLORBG = color.NRGBA{R: 40, G: 44, B: 52, A: 0xff}

// var COLORTXT = color.NRGBA{R: 0xff, G: 0xb7, B: 0xce, A: 0xff}
var COLORTXT = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
var COLORLINES = color.NRGBA{R: 250, G: 67, B: 114, A: 0xff}
var COLORBRD = color.NRGBA{R: 40, G: 44, B: 52, A: 0xff}
var COLORSTR = color.NRGBA{R: 0x5f, G: 0x9e, B: 0xa0, A: 0xff}

// cell widget container
type CellWidgetContainer struct {
	widget.BaseWidget
	Container fyne.Container
}

func NewCellWidgetContainer(content []fyne.CanvasObject) *CellWidgetContainer {
	item := &CellWidgetContainer{
		Container: *container.NewWithoutLayout(),
	}
	item.Container.Objects = content
	item.ExtendBaseWidget(item)
	return item
}

func (item *CellWidgetContainer) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(&item.Container)
}

func ZoomRefresh() {
	zoom, _ := GUIZOOM.Get()
	GUIZOOM.Set(zoom - 0.1)
	GUIZOOM.Set(zoom + 0.1)
}

func (item *CellWidgetContainer) Scrolled(d *fyne.ScrollEvent) {
	zoom, _ := GUIZOOM.Get()
	newZoom := zoom
	fmt.Println("container position on zoom", item.Container.Position())
	if d.Scrolled.DY > 0 {
		if zoom < 1 {
			newZoom = zoom + 0.1
			GUIZOOM.Set(newZoom)
		} else {
			GUIZOOM.Set(1)
		}
	} else {
		if zoom > 0.1 {
			newZoom = zoom - 0.1
			GUIZOOM.Set(newZoom)
		} else {
			GUIZOOM.Set(0.1)
		}
	}
	newOffsetX := d.Position.X - (d.Position.X-item.Container.Position().X)*float32(newZoom/zoom)
	newOffsetY := d.Position.Y - (d.Position.Y-item.Container.Position().Y)*float32(newZoom/zoom)
	item.Container.Move(fyne.NewPos(newOffsetX, newOffsetY))
}

func realCoordinates(pos, contpos fyne.Position) (int, int) {
	zoom, _ := GUIZOOM.Get()
	realX := (pos.X - contpos.X) / float32(zoom)
	realY := (pos.Y - contpos.Y) / float32(zoom)

	return int(realX), int(realY)
}

func (item *CellWidgetContainer) Tapped(e *fyne.PointEvent) {
	fmt.Println(e.Position, item.Container.Position())
	u, _ := GUIDATAUPDATER.Get()
	GUIDATAUPDATER.Set(u + 1)
	SELECTED = []string{}
	if IsCreateSelect {
		realX, realY := realCoordinates(e.Position, item.Container.Position())
		key, err := USERMIND.AddCell([2]int{realX, realY})
		checkErr(err)
		myw := NewCellWidget(key, USERMIND.Cells[key])
		item.Container.Objects = append(item.Container.Objects, myw)
		item.Refresh()
		ZoomRefresh()
		IsCreateSelect = false
	}
}

func (item *CellWidgetContainer) Dragged(d *fyne.DragEvent) {
	item.Container.Move(fyne.NewPos(item.Container.Position().X+d.Dragged.DX, item.Container.Position().Y+d.Dragged.DY))
	fmt.Println("container position on move", item.Container.Position())
}

func (item *CellWidgetContainer) DragEnd() {
	fmt.Println("Background drag end")
}

// move icon widget
type CellWidgetMoveIcon struct {
	widget.Icon
	OnDragStart func(d *fyne.DragEvent)
	OnDragEnd   func()
}

func newCellWidgetMoveIcon(res fyne.Resource) *CellWidgetMoveIcon {
	icon := &CellWidgetMoveIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)

	return icon
}

func (icon *CellWidgetMoveIcon) Dragged(d *fyne.DragEvent) {
	icon.OnDragStart(d)
}

func (icon *CellWidgetMoveIcon) DragEnd() {
	icon.OnDragEnd()
}

// cell widget
type CellWidget struct {
	widget.BaseWidget
	Id            string
	Cell          *Cell
	Movebtn       *CellWidgetMoveIcon
	Background    *canvas.Rectangle
	Textcontainer *fyne.Container
	Zoom          float32
}

func NewCellWidget(key string, cell *Cell) *CellWidget {
	movebnt := newCellWidgetMoveIcon(theme.Icon(theme.IconNameViewZoomFit))
	movebnt.Hidden = true

	obj := canvas.NewRectangle(COLORBRD)
	obj.StrokeColor = COLORSTR
	obj.StrokeWidth = 1

	item := &CellWidget{
		Id:            key,
		Cell:          cell,
		Movebtn:       movebnt,
		Background:    obj,
		Zoom:          1,
		Textcontainer: container.NewWithoutLayout(),
	}
	item.ExtendBaseWidget(item)

	list := binding.NewDataListener(func() {
		zoom, _ := GUIZOOM.Get()
		item.Resize(fyne.NewSize(float32(cell.Size[0])*float32(zoom), float32(cell.Size[1])*float32(zoom)))
		item.Move(fyne.NewPos(float32(cell.Position[0])*float32(zoom), float32(cell.Position[1])*float32(zoom)))
		item.Movebtn.Resize(fyne.NewSize(20, 20))
		item.Movebtn.Move(fyne.NewPos(-20, -20))
		item.Refresh()
	})

	GUIZOOM.AddListener(list)

	list2 := binding.NewDataListener(func() {
		item.Movebtn.Hide()
		item.Background.StrokeColor = COLORSTR
		item.Refresh()
	})

	GUIDATAUPDATER.AddListener(list2)

	return item
}

func (item *CellWidget) genText() {
	linesList := strings.Split(item.Cell.Content, "\n") // Use camelCase consistently
	var maxLineLength int
	var y float32
	zoom, _ := GUIZOOM.Get()
	fontSize := float32(FONTSIZE)                       // Convert FONTSIZE to float32
	lineSpacing := fontSize / 2
	textSize := fontSize * float32(zoom)
	yIncrement := (fontSize + lineSpacing) * float32(zoom)      // Correct type handling
	lines := make([]fyne.CanvasObject, 0, len(linesList)) // Preallocate slice

	for _, line := range linesList {
		lineLength := utf8.RuneCountInString(line)
		if lineLength > maxLineLength {
			maxLineLength = lineLength
		}

		text := canvas.NewText(line, COLORTXT)
		text.TextStyle.Monospace = true
		text.TextSize = textSize
		text.Move(fyne.NewPos(0, y))
		lines = append(lines, text)

		y += yIncrement
	}

	// Avoid recalculating the size multiple times
	item.Cell.Size = [2]int{
		maxLineLength * FONTSIZE * 2 / 3,
		len(linesList) * FONTSIZE * 6 / 4,
	}

	// Update container
	item.Textcontainer.Objects = lines
}


func (item *CellWidget) Tapped(_ *fyne.PointEvent) {
	item.Movebtn.Show()
	item.Background.StrokeColor = COLORLINES
	item.Refresh()
	SELECTED = append(SELECTED, item.Id)
}

func (item *CellWidget) DoubleTapped(_ *fyne.PointEvent) {
	err := USERMIND.UpdateCell(item.Id)
	if err != nil {
		panic(err)
	} else {
		item.genText()
		item.Refresh()
		ZoomRefresh()
	}
}

func (item *CellWidget) CreateRenderer() fyne.WidgetRenderer {
	item.Movebtn.OnDragStart = func(d *fyne.DragEvent) {
		item.Move(fyne.NewPos(item.Position().X+d.Dragged.DX, item.Position().Y+d.Dragged.DY))
	}
	item.Movebtn.OnDragEnd = func() {
		zoom, _ := GUIZOOM.Get()
		USERMIND.Cells[item.Id].Position = [2]int{int(item.Position().X / float32(zoom)), int(item.Position().Y / float32(zoom))}
		saveData()
	}

	list := binding.NewDataListener(func() {
		go item.genText()
	})
	GUIZOOM.AddListener(list)

	//c := container.NewStack(item.Background, item.Textcontainer, container.NewWithoutLayout(item.Movebtn))
	c := container.NewStack(item.Textcontainer, container.NewWithoutLayout(item.Movebtn))
	return widget.NewSimpleRenderer(c)
}

func RecurceAddGuiCells() []fyne.CanvasObject {
	var celllist []fyne.CanvasObject
	for i := range USERMIND.Cells {
		e := USERMIND.Cells[i]
		if e.Status == CellStatusConfig {
			continue
		}
		myw := NewCellWidget(i, e)

		celllist = append(celllist, myw)
	}

	return celllist
}

func start() {
	myApp := app.New()
	myWindow := myApp.NewWindow("WorkOverlord")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter your password")

	form := widget.NewForm(widget.NewFormItem("Password", passwordEntry))

	form.OnSubmit = func() {
		SECRETKEY = sha256.Sum256([]byte(passwordEntry.Text))

		fmt.Println("secret is ", SECRETKEY)

		if initDb() != nil {
			dialog.ShowInformation("Error", "Wrong password", myWindow)
		} else {
			initGui(myWindow)
		}
	}

	passwordEntry.OnSubmitted = func(_ string) {
		form.OnSubmit()
	}

	content := container.NewVBox(
		widget.NewLabel("Please enter your password:"),
		form,
	)

	myWindow.SetContent(content)

	myWindow.Canvas().Focus(passwordEntry)

	myWindow.Resize(fyne.NewSize(1200, 600))
	myWindow.ShowAndRun()
}

func initGui(w fyne.Window) {

	GUIZOOM.Set(1)

	GUICONTAINER = NewCellWidgetContainer(RecurceAddGuiCells())

	addbtn := widget.NewButton("ADD", func() {
		IsCreateSelect = true
	})
	deletebtn := widget.NewButton("DELETE", func() {
		if len(SELECTED) == 0 {
			fmt.Println("no cells selected")
		} else {
			for _, v := range SELECTED {
				delete(USERMIND.Cells, v)
			}
			saveData()
			GUICONTAINER.Container.Objects = RecurceAddGuiCells()
			GUICONTAINER.Refresh()
		}
	})
	closebtn := widget.NewButton("CLOSE", func() {
		w.Close()
	})
	mainmenu := container.NewHBox(addbtn, deletebtn, closebtn)

	content := container.NewBorder(mainmenu, nil, nil, nil, GUICONTAINER)
	w.SetContent(content)

	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl}
	w.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Tab")
	})
}
