package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/DorianPeregrim/proxx"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	app := proxx.NewApplication()
	if err := app.Start(); err != nil {
		fmt.Println(err)
		return
	}
}
