package screen

import (
	"errors"
	"math/rand"
)

type Field [][]rune

func NewField(width int, height int) (Field, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("width and height must be positive values")
	}

	result := make(Field, height)

	for i := range result {
		result[i] = make([]rune, width)
	}

	return result, nil
}

func (f Field) SetPlain(char rune) {
	for i := range f {
		for j := range f[i] {
			f[i][j] = char
		}
	}
}

func (f Field) SetRandomWith(chars []rune) {
	if len(chars) == 0 {
		return
	}

	var index int
	
	for i := range f {
		for j := range f[i] {
			index = rand.Intn(len(chars))
			f[i][j] = chars[index]
		}
	}
}