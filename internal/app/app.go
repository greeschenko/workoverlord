package app

import (
	"fmt"
	"greeschenko/workoverlord2/internal/interfaces"
	"sync"
)

type App struct {
	GUI     interfaces.GUIInterface
	Storage interfaces.StorageInterface
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
	fmt.Println("app component is running")
	fmt.Println("gui component is running")
    a.GUI.Start()
}
