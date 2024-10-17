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
var GUIDATAUPDATER = binding.NewInt()

var FONTSIZE = 14

// var COLORBG = color.NRGBA{R: 0x28, G: 0x2c, B: 0x34, A: 0xff}
var COLORBG = color.NRGBA{R: 40, G: 44, B: 52, A: 0xff}
var COLORTXT = color.NRGBA{R: 0xff, G: 0xb7, B: 0xce, A: 0xff}
var COLORLINES = color.NRGBA{R: 250, G: 67, B: 114, A: 0xff}
var COLORBRD = color.NRGBA{R: 40, G: 44, B: 52, A: 0xff}
var COLORSTR = color.NRGBA{R: 0x5f, G: 0x9e, B: 0xa0, A: 0xff}

// cell widget container
type CellWidgetContainer struct {
	widget.BaseWidget
	Content   []fyne.CanvasObject
	Container fyne.CanvasObject
}

func (item *CellWidgetContainer) CreateRenderer() fyne.WidgetRenderer {
	obj := canvas.NewRectangle(COLORBG)
	c := container.NewStack(obj, item.Container)
	return widget.NewSimpleRenderer(c)
}

func (item *CellWidgetContainer) Scrolled(d *fyne.ScrollEvent) {
	zoom, _ := GUIZOOM.Get()
	fmt.Println("container position on zoom", item.Container.Position())
	//item.Container.Move(fyne.NewPos(1000, 1000))
	if d.Scrolled.DY > 0 {
		if zoom < 1 {
			GUIZOOM.Set(zoom + 0.1)
		} else {
			GUIZOOM.Set(1)
		}
	} else {
		if zoom > 0.1 {
			GUIZOOM.Set(zoom - 0.1)
		} else {
			GUIZOOM.Set(0.1)
		}
	}
}

func (item *CellWidgetContainer) Tapped(_ *fyne.PointEvent) {
	u, _ := GUIDATAUPDATER.Get()
	GUIDATAUPDATER.Set(u + 1)
}

func (item *CellWidgetContainer) Dragged(d *fyne.DragEvent) {
	item.Container.Move(fyne.NewPos(item.Container.Position().X+d.Dragged.DX, item.Container.Position().Y+d.Dragged.DY))
	fmt.Println("container position on move", item.Container.Position())
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
	Cell       *Cell
	Movebtn    *CellWidgetMoveIcon
	Resizebtn  *CellWidgetMoveIcon
	Background *canvas.Rectangle
	Zoom       float32
}

func NewCellWidget(cell *Cell) *CellWidget {
	movebnt := newCellWidgetMoveIcon(theme.Icon(theme.IconNameViewZoomFit))
	movebnt.Hidden = true

	resizebnt := newCellWidgetMoveIcon(theme.Icon(theme.IconNameViewFullScreen))
	resizebnt.Hidden = true

	obj := canvas.NewRectangle(COLORBRD)
	obj.StrokeColor = COLORSTR
	obj.StrokeWidth = 1

	item := &CellWidget{
		Cell:       cell,
		Movebtn:    movebnt,
		Resizebtn:  resizebnt,
		Background: obj,
		Zoom:       1,
	}
	item.ExtendBaseWidget(item)

	return item
}

func (item *CellWidget) Tapped(_ *fyne.PointEvent) {
	item.Movebtn.Show()
	item.Resizebtn.Show()
	item.Background.StrokeColor = COLORLINES
	item.Refresh()
	log.Println("I have been tapped")
}

func (item *CellWidget) DoubleTapped(_ *fyne.PointEvent) {
	fmt.Println("Widget double tapped")
}

//func (item *CellWidget) Dragged(d *fyne.DragEvent) {
//	item.Move(d.AbsolutePosition)
//	log.Println("Drag start", d.Position)
//}

//func (item *CellWidget) DragEnd() {
//	log.Println("Drag end")
//}

func (item *CellWidget) CreateRenderer() fyne.WidgetRenderer {
	item.Movebtn.OnDragStart = func(d *fyne.DragEvent) {
		item.Move(fyne.NewPos(item.Position().X+d.Dragged.DX, item.Position().Y+d.Dragged.DY))
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

	maxcharacters := int(item.Cell.Size[0] / (FONTSIZE / 3))
	maxlines := int(item.Cell.Size[1] / FONTSIZE)

	lineslist := strings.Split(item.Cell.Content, "\n")
	var lines []fyne.CanvasObject
	for i := range lineslist {
		if i < maxlines {
			e := lineslist[i]
			if len(e) > maxcharacters {
				e = e[0:maxcharacters]
			}
			text := canvas.NewText(e, COLORTXT)
			text.TextStyle.Monospace = true
			list := binding.NewDataListener(func() {
				zoom, _ := GUIZOOM.Get()
				//text.TextSize = float32(FONTSIZE) * float32(zoom)
				//text.Refresh()
				fmt.Println("zoom changed ", zoom)

			})
			GUIZOOM.AddListener(list)
			lines = append(lines, text)
		} else {
			break
		}
	}
	text := container.NewVBox(lines...)
	text.Resize(fyne.NewSize(100, 100))
	text1 := container.NewScroll(text)
	text1.Direction = 3
	wrap := container.New(
		layout.NewCustomPaddedLayout(
			10*item.Zoom,
			10*item.Zoom,
			10*item.Zoom,
			10*item.Zoom,
		),
		text1,
	)
	c := container.NewStack(item.Background, wrap, container.NewWithoutLayout(item.Movebtn, item.Resizebtn))
	return widget.NewSimpleRenderer(c)
}

func RecurceAddGuiCells(mind *MIND, celllist []fyne.CanvasObject) []fyne.CanvasObject {
	for i := range mind.Cells {
		e := mind.Cells[i]
		if e.Status == CellStatusConfig {
			continue
		}
		myw := NewCellWidget(e)
		list := binding.NewDataListener(func() {
			zoom, _ := GUIZOOM.Get()
			myw.Resize(fyne.NewSize(float32(e.Size[0])*float32(zoom), float32(e.Size[1])*float32(zoom)))
			myw.Move(fyne.NewPos(float32(e.Position[0])*float32(zoom), float32(e.Position[1])*float32(zoom)))
			myw.Movebtn.Resize(fyne.NewSize(20, 20))
			myw.Movebtn.Move(fyne.NewPos(-20, -20))
			myw.Resizebtn.Resize(fyne.NewSize(20, 20))
			myw.Resizebtn.Move(fyne.NewPos(float32(e.Size[0])*float32(zoom), float32(e.Size[1])*float32(zoom)))
			myw.Refresh()
		})

		GUIZOOM.AddListener(list)

		list2 := binding.NewDataListener(func() {
			myw.Movebtn.Hide()
			myw.Resizebtn.Hide()
			myw.Background.StrokeColor = COLORSTR
			myw.Refresh()
		})

		GUIDATAUPDATER.AddListener(list2)

		celllist = append(celllist, myw)
	}

	return celllist
}

func initGui() {

	var celllist []fyne.CanvasObject
	GUIZOOM.Set(1)

	myApp := app.New()
	w := myApp.NewWindow("WorkOverlord")

	celllist = RecurceAddGuiCells(USERMIND, celllist)

	cont := NewCellWidgetContainer(celllist)
	addbtn := widget.NewButton("add new", func() {
		err := USERMIND.AddCell()
		checkErr(err)
	})
	closebtn := widget.NewButton("close", func() {
        w.Close()
	})
	mainmenu := container.NewHBox(addbtn, closebtn)

	content := container.NewBorder(mainmenu, nil, nil, nil, cont)
	w.SetContent(content)

	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl}
	w.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Tab")
	})

	w.Resize(fyne.NewSize(1200, 600))
	w.ShowAndRun()
}
