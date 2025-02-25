package color

const BLACK = "#000000"

type HairColor int
type BasicHairColor int

const (
	HairColor1   HairColor = iota // #000000 Black
	HairColor2                    // #4B3621 Dark Brown
	HairColor3                    // #3B2F2F Dark Brownish Gray
	HairColor4                    // #5B3A29 Chestnut
	HairColor5                    // #7B4B3A Coffee
	HairColor6                    // #A52A2A Brown
	HairColor7                    // #8B4513 Saddle Brown
	HairColor8                    // #A0522D Sienna
	HairColor9                    // #D2691E Chocolate
	HairColor10                   // #C0C0C0 Silver
	HairColor11                   // #808080 Gray
	HairColor12                   // #B22222 Firebrick
	HairColor13                   // #FF4500 Orange Red
	HairColor14                   // #FF6347 Red
	HairColor15                   // #FF7F50 Coral
	HairColor16                   // #FF8C00 Dark Orange
	HairColor17                   // #F5DEB3 Wheat
	HairColor18                   // #F0E68C Khaki
	HairColor19                   // #D3D3D3 Light Gray
	HairColor20                   // #DCDCDC Gainsboro
	HairColor21                   // #B0C4DE Light Steel Blue
	HairColor22                   // #ADD8E6 Light Blue
	HairColor23                   // #87CEFA Light Sky Blue
	HairColor24                   // #6B4C3A Dark Coffee
	HairColor25                   // #5B3A29 Chestnut Brown
	HairColor26                   // #4E3B31 Dark Brownish
	HairColor27                   // #7B4B3A Medium Coffee
	HairColor28                   // #A0522D Light Sienna
	HairColor29                   // #D2691E Medium Chocolate
	HairColor30                   // #C0C0C0 Light Silver
	HairColor31                   // #A9A9A9 Dark Gray
	HairColor32                   // #B22222 Medium Firebrick
	HairColor33                   // #FF4500 Light Orange Red
	HairColor34                   // #FF6347 Light Red
	HairColor35                   // #FF7F50 Light Coral
	HairColor36                   // #FF8C00 Medium Dark Orange
	HairColor37                   // #F5DEB3 Light Wheat
	HairColor38                   // #F0E68C Light Khaki
	HairColor39                   // #D3D3D3 Dark Light Gray
	HairColor40                   // #DCDCDC Light Gainsboro
	HairColor41                   // #B0C4DE Medium Light Steel Blue
	HairColor42                   // #ADD8E6 Medium Light Blue
	HairColor43                   // #87CEFA Medium Light Sky Blue
	HairColor44                   // #4B3621 Dark Brown
	HairColor45                   // #3B2F2F Dark Grayish Brown
	HairColor46                   // #5B3A29 Chestnut
	HairColor47                   // #7B4B3A Coffee
	HairColor48                   // #A52A2A Brown
	HairColor49                   // #8B4513 Saddle Brown
	HairColor50                   // #A0522D Sienna
	HairColor51                   // #D2691E Chocolate
	HairColor52                   // #C0C0C0 Silver
	HairColor53                   // #808080 Gray
	HairColor54                   // #B22222 Firebrick
	HairColor55                   // #FF4500 Orange Red
	HairColor56                   // #FF6347 Red
	HairColor57                   // #FF7F50 Coral
	HairColor58                   // #FF8C00 Dark Orange
	HairColor59                   // #F5DEB3 Wheat
	HairColor60                   // #F0E68C Khaki
	HairColor61                   // #D3D3D3 Light Gray
	HairColor62                   // #DCDCDC Gainsboro
	HairColor63                   // #B0C4DE Light Steel Blue
	HairColor64                   // #ADD8E6 Light Blue
	HairColor65                   // #87CEFA Light Sky Blue
	HairColor66                   // #6B4C3A Dark Coffee
	HairColor67                   // #5B3A29 Chestnut Brown
	HairColor68                   // #4E3B31 Dark Brownish
	HairColor69                   // #7B4B3A Medium Coffee
	HairColor70                   // #A0522D Light Sienna
	HairColor71                   // #D2691E Medium Chocolate
	HairColor72                   // #C0C0C0 Light Silver
	HairColor73                   // #A9A9A9 Dark Gray
	HairColor74                   // #B22222 Medium Firebrick
	HairColor75                   // #FF4500 Light Orange Red
	HairColor76                   // #FF6347 Light Red
	HairColor77                   // #FF7F50 Light Coral
	HairColor78                   // #FF8C00 Medium Dark Orange
	HairColor79                   // #F5DEB3 Light Wheat
	HairColor80                   // #F0E68C Light Khaki
	HairColor81                   // #D3D3D3 Dark Light Gray
	HairColor82                   // #DCDCDC Light Gainsboro
	HairColor83                   // #B0C4DE Medium Light Steel Blue
	HairColor84                   // #ADD8E6 Medium Light Blue
	HairColor85                   // #87CEFA Medium Light Sky Blue
	HairColor86                   // #4B3621 Dark Brown
	HairColor87                   // #3B2F2F Dark Grayish Brown
	HairColor88                   // #5B3A29 Chestnut
	HairColor89                   // #7B4B3A Coffee
	HairColor90                   // #A52A2A Brown
	HairColor91                   // #8B4513 Saddle Brown
	HairColor92                   // #A0522D Sienna
	HairColor93                   // #D2691E Chocolate
	HairColor94                   // #C0C0C0 Silver
	HairColor95                   // #808080 Gray
	HairColor96                   // #B22222 Firebrick
	HairColor97                   // #FF4500 Orange Red
	HairColor98                   // #FF6347 Red
	HairColor99                   // #FF7F50 Coral
	HairColor100                  // #FF8C00 Dark Orange
	HairColorAmount
)

var HairColorShadeStrings = map[HairColor]string{
	HairColor1:   "#000000", // Black
	HairColor2:   "#4B3621", // Dark Brown
	HairColor3:   "#4B3621", // Dark Brown
	HairColor4:   "#3B2F2F", // Dark Grayish Brown
	HairColor5:   "#5B3A29", // Chestnut
	HairColor6:   "#7B4B3A", // Coffee
	HairColor7:   "#A52A2A", // Brown
	HairColor8:   "#8B4513", // Saddle Brown
	HairColor9:   "#A0522D", // Sienna
	HairColor10:  "#D2691E", // Chocolate
	HairColor11:  "#C0C0C0", // Silver
	HairColor12:  "#808080", // Gray
	HairColor13:  "#B22222", // Firebrick
	HairColor14:  "#FF4500", // Orange Red
	HairColor15:  "#FF6347", // Red
	HairColor16:  "#FF7F50", // Coral
	HairColor17:  "#FF8C00", // Dark Orange
	HairColor18:  "#F5DEB3", // Wheat
	HairColor19:  "#F0E68C", // Khaki
	HairColor20:  "#D3D3D3", // Light Gray
	HairColor21:  "#DCDCDC", // Gainsboro
	HairColor22:  "#B0C4DE", // Light Steel Blue
	HairColor23:  "#ADD8E6", // Light Blue
	HairColor24:  "#87CEFA", // Light Sky Blue
	HairColor25:  "#6B4C3A", // Dark Coffee
	HairColor26:  "#5B3A29", // Chestnut Brown
	HairColor27:  "#4E3B31", // Dark Brownish
	HairColor28:  "#7B4B3A", // Medium Coffee
	HairColor29:  "#A0522D", // Light Sienna
	HairColor30:  "#D2691E", // Medium Chocolate
	HairColor31:  "#C0C0C0", // Light Silver
	HairColor32:  "#A9A9A9", // Dark Gray
	HairColor33:  "#B22222", // Medium Firebrick
	HairColor34:  "#FF4500", // Light Orange Red
	HairColor35:  "#FF6347", // Light Red
	HairColor36:  "#FF7F50", // Light Coral
	HairColor37:  "#FF8C00", // Medium Dark Orange
	HairColor38:  "#F5DEB3", // Light Wheat
	HairColor39:  "#F0E68C", // Light Khaki
	HairColor40:  "#D3D3D3", // Dark Light Gray
	HairColor41:  "#DCDCDC", // Light Gainsboro
	HairColor42:  "#B0C4DE", // Medium Light Steel Blue
	HairColor43:  "#ADD8E6", // Medium Light Blue
	HairColor44:  "#87CEFA", // Medium Light Sky Blue
	HairColor45:  "#4B3621", // Dark Brown
	HairColor46:  "#3B2F2F", // Dark Grayish Brown
	HairColor47:  "#5B3A29", // Chestnut
	HairColor48:  "#7B4B3A", // Coffee
	HairColor49:  "#A52A2A", // Brown
	HairColor50:  "#8B4513", // Saddle Brown
	HairColor51:  "#A0522D", // Sienna
	HairColor52:  "#D2691E", // Chocolate
	HairColor53:  "#C0C0C0", // Silver
	HairColor54:  "#808080", // Gray
	HairColor55:  "#B22222", // Firebrick
	HairColor56:  "#FF4500", // Orange Red
	HairColor57:  "#FF6347", // Red
	HairColor58:  "#FF7F50", // Coral
	HairColor59:  "#FF8C00", // Dark Orange
	HairColor60:  "#F5DEB3", // Wheat
	HairColor61:  "#F0E68C", // Khaki
	HairColor62:  "#D3D3D3", // Light Gray
	HairColor63:  "#DCDCDC", // Gainsboro
	HairColor64:  "#B0C4DE", // Light Steel Blue
	HairColor65:  "#ADD8E6", // Light Blue
	HairColor66:  "#87CEFA", // Light Sky Blue
	HairColor67:  "#6B4C3A", // Dark Coffee
	HairColor68:  "#5B3A29", // Chestnut Brown
	HairColor69:  "#4E3B31", // Dark Brownish
	HairColor70:  "#7B4B3A", // Medium Coffee
	HairColor71:  "#A0522D", // Light Sienna
	HairColor72:  "#D2691E", // Medium Chocolate
	HairColor73:  "#C0C0C0", // Light Silver
	HairColor74:  "#A9A9A9", // Dark Gray
	HairColor75:  "#B22222", // Medium Firebrick
	HairColor76:  "#FF4500", // Light Orange Red
	HairColor77:  "#FF6347", // Light Red
	HairColor78:  "#FF7F50", // Light Coral
	HairColor79:  "#FF8C00", // Medium Dark Orange
	HairColor80:  "#F5DEB3", // Light Wheat
	HairColor81:  "#F0E68C", // Light Khaki
	HairColor82:  "#D3D3D3", // Dark Light Gray
	HairColor83:  "#DCDCDC", // Light Gainsboro
	HairColor84:  "#B0C4DE", // Medium Light Steel Blue
	HairColor85:  "#ADD8E6", // Medium Light Blue
	HairColor86:  "#87CEFA", // Medium Light Sky Blue
	HairColor87:  "#4B3621", // Dark Brown
	HairColor88:  "#3B2F2F", // Dark Grayish Brown
	HairColor89:  "#5B3A29", // Chestnut
	HairColor90:  "#7B4B3A", // Coffee
	HairColor91:  "#A52A2A", // Brown
	HairColor92:  "#8B4513", // Saddle Brown
	HairColor93:  "#A0522D", // Sienna
	HairColor94:  "#D2691E", // Chocolate
	HairColor95:  "#C0C0C0", // Silver
	HairColor96:  "#808080", // Gray
	HairColor97:  "#B22222", // Firebrick
	HairColor98:  "#FF4500", // Orange Red
	HairColor99:  "#FF6347", // Red
	HairColor100: "#FF7F50", // Coral
}

var BasicHairColorShadeStrings = map[HairColor]string{
	HairColor1:  "#000000", // Black
	HairColor2:  "#4B3621", // Dark Brown
	HairColor5:  "#7B4B3A", // Coffee
	HairColor6:  "#A52A2A", // Brown
	HairColor7:  "#8B4513", // Saddle Brown
	HairColor11: "#808080", // Gray
	HairColor12: "#B22222", // Firebrick
	HairColor19: "#D3D3D3", // Light Gray
	HairColor20: "#DCDCDC", // Gainsboro
	HairColor31: "#A9A9A9", // Dark Gray
	HairColor44: "#4B3621", // Dark Brown
	HairColor45: "#3B2F2F", // Dark Grayish Brown
	HairColor46: "#5B3A29", // Chestnut
	HairColor47: "#7B4B3A", // Coffee
	HairColor50: "#A0522D", // Sienna
}

func (a HairColor) String() string {
	return HairColorShadeStrings[a]
}

// ランダムにHairColorを抽選する
func RandomHairColor() string {
	no := Random(int(HairColorAmount - 1))
	return HairColor(no).String()
}
