package head

import (
	"github.com/tKwbr999/pixpal/app/pal/model"
)

func NewShortHair(skinColor, eyeColor, hairColor, earColor string) []*model.Pos {
	//face
	face := []*model.Pos{
		{X: 3, Y: 9, Color: skinColor},
		{X: 4, Y: 9, Color: skinColor},
		{X: 5, Y: 9, Color: skinColor},
		{X: 6, Y: 9, Color: skinColor},
		{X: 7, Y: 9, Color: skinColor},
		{X: 8, Y: 9, Color: skinColor},
		{X: 3, Y: 8, Color: skinColor},
		{X: 4, Y: 8, Color: skinColor},
		{X: 5, Y: 8, Color: skinColor},
		{X: 6, Y: 8, Color: skinColor},
		{X: 7, Y: 8, Color: skinColor},
		{X: 8, Y: 8, Color: skinColor},
		{X: 3, Y: 7, Color: skinColor},
		{X: 5, Y: 7, Color: skinColor},
		{X: 6, Y: 7, Color: skinColor},
		{X: 8, Y: 7, Color: skinColor},
		{X: 3, Y: 6, Color: skinColor},
		{X: 4, Y: 6, Color: skinColor},
		{X: 5, Y: 6, Color: skinColor},
		{X: 6, Y: 6, Color: skinColor},
		{X: 7, Y: 6, Color: skinColor},
		{X: 8, Y: 6, Color: skinColor},
	}
	//leftEye
	leftEye := []*model.Pos{
		{X: 4, Y: 7, Color: eyeColor},
	}
	//rightEye
	rightEye := []*model.Pos{
		{X: 7, Y: 7, Color: eyeColor},
	}
	//leftEar
	leftEar := []*model.Pos{
		{X: 2, Y: 7, Color: earColor},
	}
	//rightEar
	rightEar := []*model.Pos{
		{X: 9, Y: 7, Color: earColor},
	}
	//hair
	hair := []*model.Pos{
		{X: 3, Y: 5, Color: hairColor},
		{X: 4, Y: 5, Color: hairColor},
		{X: 5, Y: 5, Color: hairColor},
		{X: 6, Y: 5, Color: hairColor},
		{X: 7, Y: 5, Color: hairColor},
		{X: 8, Y: 5, Color: hairColor},
		{X: 3, Y: 4, Color: hairColor},
		{X: 4, Y: 4, Color: hairColor},
		{X: 5, Y: 4, Color: hairColor},
		{X: 6, Y: 4, Color: hairColor},
		{X: 7, Y: 4, Color: hairColor},
		{X: 8, Y: 4, Color: hairColor},
		{X: 4, Y: 3, Color: hairColor},
		{X: 5, Y: 3, Color: hairColor},
		{X: 6, Y: 3, Color: hairColor},
		{X: 7, Y: 3, Color: hairColor},
		{X: 5, Y: 2, Color: hairColor},
		{X: 6, Y: 2, Color: hairColor},
	}
	allParts := append(face, leftEye...)
	allParts = append(allParts, rightEye...)
	allParts = append(allParts, leftEar...)
	allParts = append(allParts, rightEar...)
	allParts = append(allParts, hair...)
	return allParts

}
