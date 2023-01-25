package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/DorianPeregrim/proxx"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	g := proxx.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

	//app := proxx.NewApplication()
	//if err := app.Start(); err != nil {
	//	fmt.Println(err)
	//	return
	//}
}
