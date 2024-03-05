package game

import (
	"fmt"
	"math"
	
	"github.com/shamik230/game-of-life/screen"
)

type gameField struct {
	prev screen.Field
	next screen.Field
}

func NewGame(gameScreen *screen.Screen) {
	var (
		width  = gameScreen.GetWidth()
		height = gameScreen.GetHeight()
	)

	gameField, err := newGameField(width, height)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	gameField.next.SetPlain('1')

	for {
		gameScreen.SetField(gameField.next)
		gameScreen.RenderFrame()
		gameField.evalNextFrame()
	}
}

func newGameField(width int, height int) (*gameField, error) {
	prev, err := screen.NewField(width, height)
	if err != nil {
		return nil, err
	}

	next, err := screen.NewField(width, height)
	if err != nil {
		return nil, err
	}

	return &gameField{
		prev: prev,
		next: next,
	}, nil
}

func (f *gameField) evalNextFrame() {
	f.prev = f.next
	for i := range f.prev {
		for j := range f.prev[i] {
			if f.prev[i][j] == '1' {
				f.next[i][j] = '0'
			} else {
				f.next[i][j] = '1'
			}
		}
	}
}

func modulo(x, base int) int {
	y := math.Floor(float64(x) / float64(base))
	return x - int(y) * base
}