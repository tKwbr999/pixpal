package body

import (
	"github.com/tKwbr999/pixpal/app/pal/model"
)

// NewLongSleeve creates a new long sleeve with color
func NewLongSleeve(skinColor, sleeveColor string) []*model.Pos {
	body := []*model.Pos{
		{X: 3, Y: 18, Color: sleeveColor},
		{X: 3, Y: 17, Color: sleeveColor},
		{X: 3, Y: 16, Color: sleeveColor},
		{X: 3, Y: 15, Color: sleeveColor},
		{X: 3, Y: 14, Color: sleeveColor},
		{X: 3, Y: 13, Color: sleeveColor},
		{X: 3, Y: 12, Color: sleeveColor},
		{X: 4, Y: 18, Color: sleeveColor},
		{X: 4, Y: 17, Color: sleeveColor},
		{X: 4, Y: 16, Color: sleeveColor},
		{X: 4, Y: 15, Color: sleeveColor},
		{X: 4, Y: 14, Color: sleeveColor},
		{X: 4, Y: 13, Color: sleeveColor},
		{X: 4, Y: 12, Color: sleeveColor},
		{X: 5, Y: 18, Color: sleeveColor},
		{X: 5, Y: 17, Color: sleeveColor},
		{X: 5, Y: 16, Color: sleeveColor},
		{X: 5, Y: 15, Color: sleeveColor},
		{X: 5, Y: 14, Color: sleeveColor},
		{X: 5, Y: 13, Color: sleeveColor},
		{X: 5, Y: 12, Color: sleeveColor},
		{X: 6, Y: 18, Color: sleeveColor},
		{X: 6, Y: 17, Color: sleeveColor},
		{X: 6, Y: 16, Color: sleeveColor},
		{X: 6, Y: 15, Color: sleeveColor},
		{X: 6, Y: 14, Color: sleeveColor},
		{X: 6, Y: 13, Color: sleeveColor},
		{X: 6, Y: 12, Color: sleeveColor},
		{X: 7, Y: 18, Color: sleeveColor},
		{X: 7, Y: 17, Color: sleeveColor},
		{X: 7, Y: 16, Color: sleeveColor},
		{X: 7, Y: 15, Color: sleeveColor},
		{X: 7, Y: 14, Color: sleeveColor},
		{X: 7, Y: 13, Color: sleeveColor},
		{X: 7, Y: 12, Color: sleeveColor},
		{X: 8, Y: 18, Color: sleeveColor},
		{X: 8, Y: 17, Color: sleeveColor},
		{X: 8, Y: 16, Color: sleeveColor},
		{X: 8, Y: 15, Color: sleeveColor},
		{X: 8, Y: 14, Color: sleeveColor},
		{X: 8, Y: 13, Color: sleeveColor},
		{X: 8, Y: 12, Color: sleeveColor},
	}
	leftLongSleeve := []*model.Pos{
		{X: 2, Y: 18, Color: sleeveColor},
		{X: 2, Y: 17, Color: sleeveColor},
		{X: 2, Y: 16, Color: sleeveColor},
		{X: 2, Y: 15, Color: sleeveColor},
		{X: 2, Y: 14, Color: sleeveColor},
		{X: 2, Y: 13, Color: sleeveColor},
		{X: 2, Y: 12, Color: sleeveColor},
	}
	rightLongSleeve := []*model.Pos{
		{X: 9, Y: 18, Color: sleeveColor},
		{X: 9, Y: 17, Color: sleeveColor},
		{X: 9, Y: 16, Color: sleeveColor},
		{X: 9, Y: 15, Color: sleeveColor},
		{X: 9, Y: 14, Color: sleeveColor},
		{X: 9, Y: 13, Color: sleeveColor},
		{X: 9, Y: 12, Color: sleeveColor},
	}
	leftHand := []*model.Pos{
		{X: 2, Y: 21, Color: skinColor},
		{X: 2, Y: 20, Color: skinColor},
		{X: 2, Y: 19, Color: skinColor},
	}
	rightHand := []*model.Pos{
		{X: 9, Y: 21, Color: skinColor},
		{X: 9, Y: 20, Color: skinColor},
		{X: 9, Y: 19, Color: skinColor},
	}
	shoulder := []*model.Pos{
		{X: 3, Y: 11, Color: sleeveColor},
		{X: 4, Y: 11, Color: sleeveColor},
		{X: 5, Y: 11, Color: sleeveColor},
		{X: 6, Y: 11, Color: sleeveColor},
		{X: 7, Y: 11, Color: sleeveColor},
		{X: 8, Y: 11, Color: sleeveColor},
	}
	neck := []*model.Pos{
		{X: 5, Y: 10, Color: skinColor},
		{X: 6, Y: 10, Color: skinColor},
	}

	// 全てのパーツを結合して返す
	allParts := append(body, leftLongSleeve...)
	allParts = append(allParts, rightLongSleeve...)
	allParts = append(allParts, leftHand...)
	allParts = append(allParts, rightHand...)
	allParts = append(allParts, shoulder...)
	allParts = append(allParts, neck...)
	return allParts
}
