package game

import (
	"fmt"

	"github.com/shamik230/game-of-life/screen"
)

func NewGame(gameScreen *screen.Screen) {
	gameScreen.SetLatencyMs(250)
	w, h := 100, 100

	initialField, err := screen.NewField(w, h);
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for x := 1; x < 100; x++ {
		for i := 0; i < x; i++ {
			for j := 0; j < x; j++ {
				initialField[i][j] = '👑'
			}
		}
		gameScreen.SetField(initialField)
		gameScreen.RenderFrame()
	}
}