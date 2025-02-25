package pal

import (
	"strings"

	"github.com/tKwbr999/pixpal/app/color"
	"github.com/tKwbr999/pixpal/app/pal/model"
	"github.com/tKwbr999/pixpal/app/parts/body"
	"github.com/tKwbr999/pixpal/app/parts/foots"
	"github.com/tKwbr999/pixpal/app/parts/head"
	"github.com/tKwbr999/pixpal/app/parts/legs"
	"github.com/tKwbr999/pixpal/app/svg"
)

func GenerateDefaultManPal() string {
	var sb strings.Builder
	sb.WriteString(svg.Header())

	skinColor := color.RandomSkinColor()
	cp := &model.ColorPallet{
		Skin:  skinColor,
		Tops:  color.RandomFashionColor(),
		Pants: color.RandomFashionColor(),
		Eye:   color.BLACK,
		Ear:   color.DarkenColor(skinColor),
		Hair:  color.BLACK,
		Foot:  color.RandomFashionColor(),
	}
	h := head.NewShortHair(cp.Skin, cp.Eye, cp.Hair, cp.Ear)
	b := body.NewLongSleeve(cp.Skin, cp.Tops)
	l := legs.NewShortPants(cp.Skin, cp.Pants)
	f := foots.NewShoes(cp.Foot)

	var posList []*model.Pos
	posList = append(posList, h...)
	posList = append(posList, b...)
	posList = append(posList, l...)
	posList = append(posList, f...)

	for _, pos := range posList {
		x := svg.Pos(pos.X)
		y := svg.Pos(pos.Y)

		rect := svg.Rect(x, y, pos.Color)
		sb.WriteString(rect)
	}
	// SVG footer
	sb.WriteString("</svg>")
	return sb.String()
}
