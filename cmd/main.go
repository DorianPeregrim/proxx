package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/DorianPeregrim/proxx"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args
	if len(args) > 1 && args[1] == "--gui" {
		ebiten.SetWindowSize(640, 480)
		ebiten.SetWindowTitle("Hello, World!")

		g := proxx.NewGame()
		if err := ebiten.RunGame(g); err != nil {
			log.Fatal(err)
		}
	} else {
		app := proxx.NewApplication()
		if err := app.Start(); err != nil {
			fmt.Println(err)
			return
		}
	}
}
