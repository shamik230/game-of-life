package game

import (
	"fmt"
	"math/rand"

	"github.com/shamik230/game-of-life/screen"
)

func NewGame(gameScreen *screen.Screen) {

	var (
		width  = gameScreen.GetWidth()
		height = gameScreen.GetHeight()
	)

	initialField, err := screen.NewField(width, height)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	gameScreen.SetField(initialField)

	for {
		w := rand.Intn(width)
		h := rand.Intn(height)
		initialField[h][w] = '🌑'
			gameScreen.SetField(initialField)
		gameScreen.RenderFrame()
	}
}
