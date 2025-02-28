package app

import (
	"greeschenko/workoverlord2/internal/interfaces"
	"greeschenko/workoverlord2/internal/mind"
	"sync"
)

type App struct {
	USERMIND *mind.MIND
	GUI      interfaces.GUIInterface
	Storage  interfaces.StorageInterface
}

var instance *App
var once sync.Once

func GetInstance() *App {
	once.Do(func() {
		instance = &App{}
	})
	return instance
}

func (a App) Run() {
	a.GUI.Start()
}
