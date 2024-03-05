package main

import (
	"fmt"

	"github.com/shamik230/game-of-life/game"
	"github.com/shamik230/game-of-life/screen"
)

func main() {
	const (
		Width  = 40
		Height = 20
	)
	
	gameScreen, err := screen.NewScreen(Width, Height)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	game.NewGame(gameScreen)
}

// TODO
// game.setField(someField)
// game.Pause()
// game.Unpause()
// game.Speed(2.5)
// game.Start()