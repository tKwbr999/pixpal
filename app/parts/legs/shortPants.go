package legs

import "github.com/tKwbr999/pixpal/app/pal/model"

// NewShortPants creates a new short pants with color
func NewShortPants(skinColor, pantsColor string) []*model.Pos {
	leg := []*model.Pos{
		// leftLeg
		{X: 4, Y: 29, Color: skinColor},
		{X: 4, Y: 28, Color: skinColor},
		{X: 4, Y: 27, Color: skinColor},
		{X: 4, Y: 26, Color: skinColor},
		{X: 4, Y: 25, Color: skinColor},
		//rightLeg
		{X: 7, Y: 29, Color: skinColor},
		{X: 7, Y: 28, Color: skinColor},
		{X: 7, Y: 27, Color: skinColor},
		{X: 7, Y: 26, Color: skinColor},
		{X: 7, Y: 25, Color: skinColor},
	}
	shortPants := []*model.Pos{
		// leftShortPants
		{X: 3, Y: 24, Color: pantsColor},
		{X: 3, Y: 23, Color: pantsColor},
		{X: 3, Y: 22, Color: pantsColor},
		{X: 3, Y: 21, Color: pantsColor},
		{X: 3, Y: 20, Color: pantsColor},
		{X: 3, Y: 19, Color: pantsColor},
		{X: 4, Y: 24, Color: pantsColor},
		{X: 4, Y: 23, Color: pantsColor},
		{X: 4, Y: 22, Color: pantsColor},
		{X: 4, Y: 21, Color: pantsColor},
		{X: 4, Y: 20, Color: pantsColor},
		{X: 4, Y: 19, Color: pantsColor},
		{X: 5, Y: 20, Color: pantsColor},
		{X: 5, Y: 19, Color: pantsColor},
		// rightShortPants
		{X: 6, Y: 20, Color: pantsColor},
		{X: 6, Y: 19, Color: pantsColor},
		{X: 7, Y: 24, Color: pantsColor},
		{X: 7, Y: 23, Color: pantsColor},
		{X: 7, Y: 22, Color: pantsColor},
		{X: 7, Y: 21, Color: pantsColor},
		{X: 7, Y: 20, Color: pantsColor},
		{X: 7, Y: 19, Color: pantsColor},
		{X: 8, Y: 24, Color: pantsColor},
		{X: 8, Y: 23, Color: pantsColor},
		{X: 8, Y: 22, Color: pantsColor},
		{X: 8, Y: 21, Color: pantsColor},
		{X: 8, Y: 20, Color: pantsColor},
		{X: 8, Y: 19, Color: pantsColor},
	}
	return append(leg, shortPants...)
}
