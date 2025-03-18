package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// cell widget container
type CellWidgetContainer struct {
	widget.BaseWidget
	Container fyne.Container
	Gui       *GUI
}

func NewCellWidgetContainer(content []fyne.CanvasObject, gui *GUI) *CellWidgetContainer {
	item := &CellWidgetContainer{
		Container: *container.NewWithoutLayout(),
		Gui:       gui,
	}
	item.Container.Objects = content
	item.ExtendBaseWidget(item)
	GUIZOOM.Set(1)
	return item
}

func (item *CellWidgetContainer) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(&item.Container)
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

func (item *CellWidgetContainer) Tapped(e *fyne.PointEvent) {
	u, _ := GUIDATAUPDATER.Get()
	GUIDATAUPDATER.Set(u + 1)
	SELECTED = []*CellWidget{}

	if IsCreateSelect {
		realX, realY := realCoordinates(e.Position, item.Container.Position())
		key, err := item.Gui.AddCell([2]int{realX, realY})
		if err != nil {
			fmt.Println("filed add new gui cell", err)
		}
		cell, err := item.Gui.Data.GetOne(key)
        if err != nil {
			fmt.Println("filed data cell not exist", err)
        }
		myw := NewCellWidget(key, cell, item.Gui)
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

func ZoomRefresh() {
	zoom, _ := GUIZOOM.Get()
	GUIZOOM.Set(zoom - 0.1)
	GUIZOOM.Set(zoom + 0.1)
}

func realCoordinates(pos, contpos fyne.Position) (int, int) {
	zoom, _ := GUIZOOM.Get()
	realX := (pos.X - contpos.X) / float32(zoom)
	realY := (pos.Y - contpos.Y) / float32(zoom)

	return int(realX), int(realY)
}
