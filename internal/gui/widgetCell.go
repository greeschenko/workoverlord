package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"greeschenko/workoverlord2/internal/models"
	"strings"
	"unicode/utf8"
)

// cell widget
type CellWidget struct {
	widget.BaseWidget
	Id            string
	Cell          *models.Cell
	Movebtn       *CellWidgetMoveIcon
	Background    *canvas.Rectangle
	Textcontainer *fyne.Container
	Gui           *GUI
}

func NewCellWidget(key string, cell *models.Cell, gui *GUI) *CellWidget {
	fmt.Println(key, cell.Size, cell.Position, gui)
	//return nil

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
		Textcontainer: container.NewWithoutLayout(),
		Gui:           gui,
	}
	item.ExtendBaseWidget(item)

	list := binding.NewDataListener(func() {
		zoom, _ := GUIZOOM.Get()
        if cell.Size != nil {
		    item.Resize(fyne.NewSize(float32(cell.Size[0])*float32(zoom), float32(cell.Size[1])*float32(zoom)))
        }
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

func (item *CellWidget) ID() string {
	return item.Id
}

func (item *CellWidget) Coordinates() [2]int {
	return *item.Cell.Position
}

func (item *CellWidget) Tapped(_ *fyne.PointEvent) {
	item.Movebtn.Show()
	item.Background.StrokeColor = COLORLINES
	item.Refresh()
	SELECTED = append(SELECTED, item.Id)
}

func (item *CellWidget) DoubleTapped(_ *fyne.PointEvent) {
	err := item.Gui.UpdateCell(item.Id)
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
		newpos := [2]int{int(item.Position().X / float32(zoom)), int(item.Position().Y / float32(zoom))}
		item.Gui.Data.Patch(item.Id, models.Cell{
			Position: &newpos,
		})
	}

	list := binding.NewDataListener(func() {
		go item.genText()
	})
	GUIZOOM.AddListener(list)

	//c := container.NewStack(item.Background, item.Textcontainer, container.NewWithoutLayout(item.Movebtn))
	c := container.NewStack(item.Textcontainer, container.NewWithoutLayout(item.Movebtn))
	return widget.NewSimpleRenderer(c)
}

func newCellWidgetMoveIcon(res fyne.Resource) *CellWidgetMoveIcon {
	icon := &CellWidgetMoveIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)

	return icon
}

// move icon widget
type CellWidgetMoveIcon struct {
	widget.Icon
	OnDragStart func(d *fyne.DragEvent)
	OnDragEnd   func()
}

func (icon *CellWidgetMoveIcon) Dragged(d *fyne.DragEvent) {
	icon.OnDragStart(d)
}

func (icon *CellWidgetMoveIcon) DragEnd() {
	icon.OnDragEnd()
}

func (item *CellWidget) genText() {
	linesList := strings.Split(item.Cell.Content, "\n") // Use camelCase consistently
	var maxLineLength int
	var y float32
	zoom, _ := GUIZOOM.Get()
	fontSize := float32(FONTSIZE) // Convert FONTSIZE to float32
	lineSpacing := fontSize / 2
	textSize := fontSize * float32(zoom)
	yIncrement := (fontSize + lineSpacing) * float32(zoom) // Correct type handling
	lines := make([]fyne.CanvasObject, 0, len(linesList))  // Preallocate slice

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
	item.Cell.Size = &[2]int{
		maxLineLength * FONTSIZE * 2 / 3,
		len(linesList) * FONTSIZE * 6 / 4,
	}

	// Update container
	item.Textcontainer.Objects = lines
}
