package main

import (
	"fmt"
)

const logo = `
 __        __    _          ___                 _               _
 \ \      / /__ | | ___ __ / _ \__   _____ _ __| | ___  _ __ __| |
  \ \ /\ / / _ \| |/ / '__| | | \ \ / / _ \ '__| |/ _ \| '__/ _' |
   \ V  V / (_) |   <| |  | |_| |\ V /  __/ |  | | (_) | | | (_| |
    \_/\_/ \___/|_|\_\_|   \___/  \_/ \___|_|  |_|\___/|_|  \__,_|
`

var (
	USERMIND = NewMIND()
)

func main() {
    initDb()
	fmt.Println(logo)
    initGui()
}
