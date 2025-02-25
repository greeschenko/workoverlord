package app

import (
	"greeschenko/workoverlord2/internal/interfaces"
	"greeschenko/workoverlord2/internal/models"
	"sync"
)

type App struct {
	USERMIND *models.MIND
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
