package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// cell widget container
type CellWidgetContainer struct {
	widget.BaseWidget
	Container      fyne.Container
	zoom           binding.Float
	updater        binding.Int
	guidataupdater binding.Int
}

func NewCellWidgetContainer(content []fyne.CanvasObject) *CellWidgetContainer {
	item := &CellWidgetContainer{
		Container: *container.NewWithoutLayout(),
	}
	item.Container.Objects = content
	item.ExtendBaseWidget(item)
	item.zoom.Set(1)
	return item
}

func (item *CellWidgetContainer) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(&item.Container)
}

func (item *CellWidgetContainer) Scrolled(d *fyne.ScrollEvent) {
	zoom, _ := item.zoom.Get()
	newZoom := zoom
	fmt.Println("container position on zoom", item.Container.Position())
	if d.Scrolled.DY > 0 {
		if zoom < 1 {
			newZoom = zoom + 0.1
			item.zoom.Set(newZoom)
		} else {
			item.zoom.Set(1)
		}
	} else {
		if zoom > 0.1 {
			newZoom = zoom - 0.1
			item.zoom.Set(newZoom)
		} else {
			item.zoom.Set(0.1)
		}
	}
	newOffsetX := d.Position.X - (d.Position.X-item.Container.Position().X)*float32(newZoom/zoom)
	newOffsetY := d.Position.Y - (d.Position.Y-item.Container.Position().Y)*float32(newZoom/zoom)
	item.Container.Move(fyne.NewPos(newOffsetX, newOffsetY))
}

func (item *CellWidgetContainer) Tapped(e *fyne.PointEvent) {
	fmt.Println(e.Position, item.Container.Position())
	//u, _ := item.guidataupdater.Get()
	// item.guidataupdater.Set(u + 1)
	// SELECTED = []string{}
	//
	//	if IsCreateSelect {
	//		realX, realY := realCoordinates(e.Position, item.Container.Position())
	//		key, err := USERMIND.AddCell([2]int{realX, realY})
	//		checkErr(err)
	//		myw := NewCellWidget(key, USERMIND.Cells[key])
	//		item.Container.Objects = append(item.Container.Objects, myw)
	//		item.Refresh()
	//		item.ZoomRefresh()
	//		IsCreateSelect = false
	//	}
}

func (item *CellWidgetContainer) Dragged(d *fyne.DragEvent) {
	item.Container.Move(fyne.NewPos(item.Container.Position().X+d.Dragged.DX, item.Container.Position().Y+d.Dragged.DY))
	fmt.Println("container position on move", item.Container.Position())
}

func (item *CellWidgetContainer) DragEnd() {
	fmt.Println("Background drag end")
}

func (item *CellWidgetContainer) ZoomRefresh() {
	zoom, _ := item.zoom.Get()
	item.zoom.Set(zoom - 0.1)
	item.zoom.Set(zoom + 0.1)
}
