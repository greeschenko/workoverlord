package main

import (
	"fmt"
	"greeschenko/workoverlord2/internal/app"
	"greeschenko/workoverlord2/internal/gui"
	"greeschenko/workoverlord2/internal/models"
	"greeschenko/workoverlord2/internal/storage"
)

const logo = `
 __        __    _          ___                 _               _
 \ \      / /__ | | ___ __ / _ \__   _____ _ __| | ___  _ __ __| |
  \ \ /\ / / _ \| |/ / '__| | | \ \ / / _ \ '__| |/ _ \| '__/ _' |
   \ V  V / (_) |   <| |  | |_| |\ V /  __/ |  | | (_) | | | (_| |
    \_/\_/ \___/|_|\_\_|   \___/  \_/ \___|_|  |_|\___/|_|  \__,_|
`

func main() {
	fmt.Print(logo)
	App := app.GetInstance()
	App.Storage = storage.NewStorage()
	App.USERMIND = models.NewMIND()
	App.GUI = gui.NewFyneGUI(App.Storage, App.USERMIND)
	App.Run()
}
