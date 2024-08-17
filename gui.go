package main

import (
	"fmt"
	"image/color"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var GUIZOOM = binding.NewFloat()

// cell widget container
type CellWidgetContainer struct {
	widget.BaseWidget
	Content   []fyne.CanvasObject
	Container fyne.CanvasObject
}

func (item *CellWidgetContainer) CreateRenderer() fyne.WidgetRenderer {
	obj := canvas.NewRectangle(color.NRGBA{R: 0x28, G: 0x2c, B: 0x34, A: 0xff})
	c := container.NewStack(obj, item.Container)
	return widget.NewSimpleRenderer(c)
}

func (item *CellWidgetContainer) Scrolled(d *fyne.ScrollEvent) {
	t, _ := GUIZOOM.Get()
	if d.Scrolled.DY > 0 {
		if t < 1 {
			GUIZOOM.Set(t + 0.1)
		}
	} else {
		if t > 0.1 {
			GUIZOOM.Set(t - 0.1)
		}
	}
}

func (item *CellWidgetContainer) Tapped(_ *fyne.PointEvent) {
	log.Println("Background been tapped")
}

func (item *CellWidgetContainer) Dragged(d *fyne.DragEvent) {
	item.Container.Move(fyne.NewPos(item.Container.Position().X+d.Dragged.DX, item.Container.Position().Y+d.Dragged.DY))
}

func (item *CellWidgetContainer) DragEnd() {
	fmt.Println("Background drag end")
}

func NewCellWidgetContainer(content []fyne.CanvasObject) *CellWidgetContainer {
	cont := container.NewWithoutLayout(content...)
	item := &CellWidgetContainer{
		Content:   content,
		Container: cont,
	}
	item.ExtendBaseWidget(item)
	return item
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
	Content    string
	Movebtn    *CellWidgetMoveIcon
	Resizebtn  *CellWidgetMoveIcon
	Background *canvas.Rectangle
}

func NewCellWidget(content string) *CellWidget {
	zoom, _ := GUIZOOM.Get()
	movebnt := newCellWidgetMoveIcon(theme.Icon(theme.IconNameViewZoomFit))
	movebnt.Resize(fyne.NewSize(30, 30))
	movebnt.Move(fyne.NewPos(-30, -30))
	movebnt.Hidden = true

	resizebnt := newCellWidgetMoveIcon(theme.Icon(theme.IconNameViewFullScreen))
	resizebnt.Resize(fyne.NewSize(30, 30))
	resizebnt.Move(fyne.NewPos(400*float32(zoom), 400*float32(zoom)))
	resizebnt.Hidden = true

	obj := canvas.NewRectangle(color.Black)
	obj.StrokeColor = color.White
	obj.StrokeWidth = 1

	item := &CellWidget{
		Content:    content,
		Movebtn:    movebnt,
		Resizebtn:  resizebnt,
		Background: obj,
	}
	item.ExtendBaseWidget(item)

	return item
}

func (item *CellWidget) Tapped(_ *fyne.PointEvent) {
	item.Movebtn.Show()
	item.Resizebtn.Show()
	item.Background.StrokeColor = color.NRGBA{R: 0xff, A: 0xff}
	item.Refresh()
	log.Println("I have been tapped")
}

//func (item *CellWidget) Dragged(d *fyne.DragEvent) {
//	item.Move(d.AbsolutePosition)
//	log.Println("Drag start", d.Position)
//}

func (item *CellWidget) DragEnd() {
	log.Println("Drag end")
}

func (item *CellWidget) CreateRenderer() fyne.WidgetRenderer {
	zoom, _ := GUIZOOM.Get()
	item.Movebtn.OnDragStart = func(d *fyne.DragEvent) {
		item.Move(d.AbsolutePosition)
		log.Println("Icon Drag start")
	}
	item.Movebtn.OnDragEnd = func() {
		log.Println("Icon Drag end")
	}

	item.Resizebtn.OnDragStart = func(d *fyne.DragEvent) {
		log.Println("resize Icon Drag start", d.Dragged)
		item.Resize(fyne.NewSize(item.Size().Width+d.Dragged.DX, item.Size().Height+d.Dragged.DY))
		item.Resizebtn.Move(fyne.NewPos(item.Resizebtn.Position().X+d.Dragged.DX, item.Resizebtn.Position().Y+d.Dragged.DY))
	}
	item.Resizebtn.OnDragEnd = func() {
		log.Println("resize Icon Drag end")
	}

	lineslist := strings.Split(item.Content, "\n")
	var lines []fyne.CanvasObject
	for i := range lineslist {
		e := lineslist[i]
		text := canvas.NewText(e, color.White)
		text.TextSize = 10 * float32(zoom)
		text.TextStyle.Monospace = true
		lines = append(lines, text)
	}
	text := container.NewVBox(lines...)
	text1 := container.NewScroll(text)
	wrap := container.New(
		layout.NewCustomPaddedLayout(
			10*float32(zoom),
			10*float32(zoom),
			10*float32(zoom),
			10*float32(zoom),
		),
		text1,
	)
	c := container.NewStack(item.Background, wrap, container.NewWithoutLayout(item.Movebtn, item.Resizebtn))
	return widget.NewSimpleRenderer(c)
}

func RecurceAddGuiCells(data []Cell, celllist []fyne.CanvasObject) []fyne.CanvasObject {
	for i := range data {
		e := data[i]
		myw := NewCellWidget(e.Data)

		list := binding.NewDataListener(func() {
			zoom, _ := GUIZOOM.Get()
			myw.Resize(fyne.NewSize(float32(e.Size[0])*float32(zoom), float32(e.Size[1])*float32(zoom)))
			myw.Move(fyne.NewPos(float32(e.Position[0])*float32(zoom), float32(e.Position[1])*float32(zoom)))
			myw.Refresh()
		})

		GUIZOOM.AddListener(list)

		celllist = append(celllist, myw)
	}

	return celllist
}

func RunGui() {

	var celllist []fyne.CanvasObject
	GUIZOOM.Set(0.5)

	myApp := app.New()
	w := myApp.NewWindow("WorkOverlord")

	celllist = RecurceAddGuiCells(USERMIND, celllist)

	list := binding.NewDataListener(func() {
		zoom, _ := GUIZOOM.Get()
		fmt.Println("TEST DATA CHANGED", zoom)
	})

	GUIZOOM.AddListener(list)

	cont := NewCellWidgetContainer(celllist)
	w.SetContent(cont)

	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl}
	w.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Tab")
	})

	w.Resize(fyne.NewSize(1200, 600))
	w.ShowAndRun()
}
