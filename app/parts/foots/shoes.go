package foots

import "github.com/tKwbr999/pixpal/app/pal/model"

func NewShoes(shoesColor string) []*model.Pos {

	shoes := []*model.Pos{
		// leftLeg
		{X: 2, Y: 30, Color: shoesColor},
		{X: 3, Y: 30, Color: shoesColor},
		{X: 4, Y: 30, Color: shoesColor},
		//rightLeg
		{X: 7, Y: 30, Color: shoesColor},
		{X: 8, Y: 30, Color: shoesColor},
		{X: 9, Y: 30, Color: shoesColor},
	}
	return shoes
}
