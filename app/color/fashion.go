package color

type FashionColor int

const (
	FashionColor1   FashionColor = iota // #000000 Black
	FashionColor2                       // #0000FF Blue
	FashionColor3                       // #008000 Green
	FashionColor4                       // #FF0000 Red
	FashionColor5                       // #FFFF00 Yellow
	FashionColor6                       // #00FFFF Cyan
	FashionColor7                       // #FF00FF Magenta
	FashionColor8                       // #FFFFFF White
	FashionColor9                       // #808080 Gray
	FashionColor10                      // #C0C0C0 Silver
	FashionColor11                      // #FFA500 Orange
	FashionColor12                      // #800080 Purple
	FashionColor13                      // #A52A2A Brown
	FashionColor14                      // #FF6347 Tomato
	FashionColor15                      // #FF4500 Orange Red
	FashionColor16                      // #FF69B4 Hot Pink
	FashionColor17                      // #FFC0CB Pink
	FashionColor18                      // #FFD700 Gold
	FashionColor19                      // #ADFF2F Green Yellow
	FashionColor20                      // #32CD32 Lime Green
	FashionColor21                      // #00FF00 Lime
	FashionColor22                      // #00BFFF Deep Sky Blue
	FashionColor23                      // #1E90FF Dodger Blue
	FashionColor24                      // #4682B4 Steel Blue
	FashionColor25                      // #5F9EA0 Cadet Blue
	FashionColor26                      // #6495ED Cornflower Blue
	FashionColor27                      // #7B68EE Medium Slate Blue
	FashionColor28                      // #9370DB Medium Purple
	FashionColor29                      // #8A2BE2 Blue Violet
	FashionColor30                      // #6A5ACD Slate Blue
	FashionColor31                      // #483D8B Dark Slate Blue
	FashionColor32                      // #2F4F4F Dark Slate Gray
	FashionColor33                      // #708090 Slate Gray
	FashionColor34                      // #778899 Light Slate Gray
	FashionColor35                      // #B0C4DE Light Steel Blue
	FashionColor36                      // #D3D3D3 Light Gray
	FashionColor37                      // #F0E68C Khaki
	FashionColor38                      // #F5F5DC Beige
	FashionColor39                      // #F5FFFA Mint Cream
	FashionColor40                      // #F0FFF0 Honeydew
	FashionColor41                      // #F0FFFF Azure
	FashionColor42                      // #F8F8FF Ghost White
	FashionColor43                      // #F0F0F0 Gainsboro
	FashionColor44                      // #F5F5F5 White Smoke
	FashionColor45                      // #FAF0E6 Linen
	FashionColor46                      // #FAFAD2 Light Goldenrod Yellow
	FashionColor47                      // #FFE4B5 Moccasin
	FashionColor48                      // #FFE4C4 Bisque
	FashionColor49                      // #FFEBCD Blanched Almond
	FashionColor50                      // #FFF0F5 Lavender Blush
	FashionColor51                      // #FFF5EE Seashell
	FashionColor52                      // #FFF8DC Cornsilk
	FashionColor53                      // #FFFFE0 Light Yellow
	FashionColor54                      // #FFFFF0 Ivory
	FashionColor55                      // #FFFFF0 Ivory
	FashionColor56                      // #FFFAFA Snow
	FashionColor57                      // #FFFAF0 Floral White
	FashionColor58                      // #FFDEAD Navajo White
	FashionColor59                      // #FFEFD5 Papaya Whip
	FashionColor60                      // #FFEFD5 Papaya Whip
	FashionColor61                      // #FFE4E1 Misty Rose
	FashionColor62                      // #FFB6C1 Light Pink
	FashionColor63                      // #FF69B4 Hot Pink
	FashionColor64                      // #FF1493 Deep Pink
	FashionColor65                      // #FF6347 Tomato
	FashionColor66                      // #FF4500 Orange Red
	FashionColor67                      // #FF8C00 Dark Orange
	FashionColor68                      // #FFA07A Light Salmon
	FashionColor69                      // #FA8072 Salmon
	FashionColor70                      // #E9967A Dark Salmon
	FashionColor71                      // #F08080 Light Coral
	FashionColor72                      // #CD5C5C Indian Red
	FashionColor73                      // #B22222 Firebrick
	FashionColor74                      // #A52A2A Brown
	FashionColor75                      // #8B4513 Saddle Brown
	FashionColor76                      // #D2691E Chocolate
	FashionColor77                      // #FF7F50 Coral
	FashionColor78                      // #FF4500 Orange Red
	FashionColor79                      // #FF6347 Tomato
	FashionColor80                      // #FF69B4 Hot Pink
	FashionColor81                      // #FFC0CB Pink
	FashionColor82                      // #FFD700 Gold
	FashionColor83                      // #ADFF2F Green Yellow
	FashionColor84                      // #32CD32 Lime Green
	FashionColor85                      // #00FF00 Lime
	FashionColor86                      // #00BFFF Deep Sky Blue
	FashionColor87                      // #1E90FF Dodger Blue
	FashionColor88                      // #4682B4 Steel Blue
	FashionColor89                      // #5F9EA0 Cadet Blue
	FashionColor90                      // #6495ED Cornflower Blue
	FashionColor91                      // #7B68EE Medium Slate Blue
	FashionColor92                      // #9370DB Medium Purple
	FashionColor93                      // #8A2BE2 Blue Violet
	FashionColor94                      // #6A5ACD Slate Blue
	FashionColor95                      // #483D8B Dark Slate Blue
	FashionColor96                      // #2F4F4F Dark Slate Gray
	FashionColor97                      // #708090 Slate Gray
	FashionColor98                      // #778899 Light Slate Gray
	FashionColor99                      // #B0C4DE Light Steel Blue
	FashionColor100                     // #D3D3D3 Light Gray
	FashionColor101                     // #F0E68C Khaki
	FashionColor102                     // #F5F5DC Beige
	FashionColor103                     // #F5FFFA Mint Cream
	FashionColor104                     // #F0FFF0 Honeydew
	FashionColor105                     // #F0FFFF Azure
	FashionColor106                     // #F8F8FF Ghost White
	FashionColor107                     // #F0F0F0 Gainsboro
	FashionColor108                     // #F5F5F5 White Smoke
	FashionColor109                     // #FAF0E6 Linen
	FashionColor110                     // #FAFAD2 Light Goldenrod Yellow
	FashionColor111                     // #FFE4B5 Moccasin
	FashionColor112                     // #FFE4C4 Bisque
	FashionColor113                     // #FFEBCD Blanched Almond
	FashionColor114                     // #FFF0F5 Lavender Blush
	FashionColor115                     // #FFF5EE Seashell
	FashionColor116                     // #FFF8DC Cornsilk
	FashionColor117                     // #FFFFE0 Light Yellow
	FashionColor118                     // #FFFFF0 Ivory
	FashionColor119                     // #FFFFF0 Ivory
	FashionColor120                     // #FFFAFA Snow
	FashionColor121                     // #FFFAF0 Floral White
	FashionColor122                     // #FFDEAD Navajo White
	FashionColor123                     // #FFEFD5 Papaya Whip
	FashionColor124                     // #FFEFD5 Papaya Whip
	FashionColor125                     // #FFE4E1 Misty Rose
	FashionColor126                     // #FFB6C1 Light Pink
	FashionColor127                     // #FF69B4 Hot Pink
	FashionColor128                     // #FF1493 Deep Pink
	FashionColor129                     // #FF6347 Tomato
	FashionColor130                     // #FF4500 Orange Red
	FashionColor131                     // #FF8C00 Dark Orange
	FashionColor132                     // #FFA07A Light Salmon
	FashionColor133                     // #FA8072 Salmon
	FashionColor134                     // #E9967A Dark Salmon
	FashionColor135                     // #F08080 Light Coral
	FashionColor136                     // #CD5C5C Indian Red
	FashionColor137                     // #B22222 Firebrick
	FashionColor138                     // #A52A2A Brown
	FashionColor139                     // #8B4513 Saddle Brown
	FashionColor140                     // #D2691E Chocolate
	FashionColor141                     // #FF7F50 Coral
	FashionColor142                     // #FF4500 Orange Red
	FashionColor143                     // #FF6347 Tomato
	FashionColor144                     // #FF69B4 Hot Pink
	FashionColor145                     // #FFC0CB Pink
	FashionColor146                     // #FFD700 Gold
	FashionColor147                     // #ADFF2F Green Yellow
	FashionColor148                     // #32CD32 Lime Green
	FashionColor149                     // #00FF00 Lime
	FashionColor150                     // #00BFFF Deep Sky Blue
	FashionColor151                     // #1E90FF Dodger Blue
	FashionColor152                     // #4682B4 Steel Blue
	FashionColor153                     // #5F9EA0 Cadet Blue
	FashionColor154                     // #6495ED Cornflower Blue
	FashionColor155                     // #7B68EE Medium Slate Blue
	FashionColor156                     // #9370DB Medium Purple
	FashionColor157                     // #8A2BE2 Blue Violet
	FashionColor158                     // #6A5ACD Slate Blue
	FashionColor159                     // #483D8B Dark Slate Blue
	FashionColor160                     // #2F4F4F Dark Slate Gray
	FashionColor161                     // #708090 Slate Gray
	FashionColor162                     // #778899 Light Slate Gray
	FashionColor163                     // #B0C4DE Light Steel Blue
	FashionColor164                     // #D3D3D3 Light Gray
	FashionColor165                     // #F0E68C Khaki
	FashionColor166                     // #F5F5DC Beige
	FashionColor167                     // #F5FFFA Mint Cream
	FashionColor168                     // #F0FFF0 Honeydew
	FashionColor169                     // #F0FFFF Azure
	FashionColor170                     // #F8F8FF Ghost White
	FashionColor171                     // #F0F0F0 Gainsboro
	FashionColor172                     // #F5F5F5 White Smoke
	FashionColor173                     // #FAF0E6 Linen
	FashionColor174                     // #FAFAD2 Light Goldenrod Yellow
	FashionColor175                     // #FFE4B5 Moccasin
	FashionColor176                     // #FFE4C4 Bisque
	FashionColor177                     // #FFEBCD Blanched Almond
	FashionColor178                     // #FFF0F5 Lavender Blush
	FashionColor179                     // #FFF5EE Seashell
	FashionColor180                     // #FFF8DC Cornsilk
	FashionColor181                     // #FFFFE0 Light Yellow
	FashionColor182                     // #FFFFF0 Ivory
	FashionColor183                     // #FFFFF0 Ivory
	FashionColor184                     // #FFFAFA Snow
	FashionColor185                     // #FFFAF0 Floral White
	FashionColor186                     // #FFDEAD Navajo White
	FashionColor187                     // #FFEFD5 Papaya Whip
	FashionColor188                     // #FFEFD5 Papaya Whip
	FashionColor189                     // #FFE4E1 Misty Rose
	FashionColor190                     // #FFB6C1 Light Pink
	FashionColor191                     // #FF69B4 Hot Pink
	FashionColor192                     // #FF1493 Deep Pink
	FashionColor193                     // #FF6347 Tomato
	FashionColor194                     // #FF4500 Orange Red
	FashionColor195                     // #FF8C00 Dark Orange
	FashionColor196                     // #FFA07A Light Salmon
	FashionColor197                     // #FA8072 Salmon
	FashionColor198                     // #E9967A Dark Salmon
	FashionColor199                     // #F08080 Light Coral
	FashionColor200                     // #CD5C5C Indian Red
	FashionColor201                     // #B22222 Firebrick
	FashionColor202                     // #A52A2A Brown
	FashionColor203                     // #8B4513 Saddle Brown
	FashionColor204                     // #D2691E Chocolate
	FashionColor205                     // #FF7F50 Coral
	FashionColor206                     // #FF4500 Orange Red
	FashionColor207                     // #FF6347 Tomato
	FashionColor208                     // #FF69B4 Hot Pink
	FashionColor209                     // #FFC0CB Pink
	FashionColor210                     // #FFD700 Gold
	FashionColor211                     // #ADFF2F Green Yellow
	FashionColor212                     // #32CD32 Lime Green
	FashionColor213                     // #00FF00 Lime
	FashionColor214                     // #00BFFF Deep Sky Blue
	FashionColor215                     // #1E90FF Dodger Blue
	FashionColor216                     // #4682B4 Steel Blue
	FashionColorAmount
)

var FashionColorShadeStrings = map[FashionColor]string{
	FashionColor1:   "#000000", // Black
	FashionColor2:   "#0000FF", // Blue
	FashionColor3:   "#008000", // Green
	FashionColor4:   "#FF0000", // Red
	FashionColor5:   "#FFFF00", // Yellow
	FashionColor6:   "#00FFFF", // Cyan
	FashionColor7:   "#FF00FF", // Magenta
	FashionColor8:   "#FFFFFF", // White
	FashionColor9:   "#808080", // Gray
	FashionColor10:  "#C0C0C0", // Silver
	FashionColor11:  "#FFA500", // Orange
	FashionColor12:  "#800080", // Purple
	FashionColor13:  "#A52A2A", // Brown
	FashionColor14:  "#FF6347", // Tomato
	FashionColor15:  "#FF4500", // Orange Red
	FashionColor16:  "#FF69B4", // Hot Pink
	FashionColor17:  "#FFC0CB", // Pink
	FashionColor18:  "#FFD700", // Gold
	FashionColor19:  "#ADFF2F", // Green Yellow
	FashionColor20:  "#32CD32", // Lime Green
	FashionColor21:  "#00FF00", // Lime
	FashionColor22:  "#00BFFF", // Deep Sky Blue
	FashionColor23:  "#1E90FF", // Dodger Blue
	FashionColor24:  "#4682B4", // Steel Blue
	FashionColor25:  "#5F9EA0", // Cadet Blue
	FashionColor26:  "#6495ED", // Cornflower Blue
	FashionColor27:  "#7B68EE", // Medium Slate Blue
	FashionColor28:  "#9370DB", // Medium Purple
	FashionColor29:  "#8A2BE2", // Blue Violet
	FashionColor30:  "#6A5ACD", // Slate Blue
	FashionColor31:  "#483D8B", // Dark Slate Blue
	FashionColor32:  "#2F4F4F", // Dark Slate Gray
	FashionColor33:  "#708090", // Slate Gray
	FashionColor34:  "#778899", // Light Slate Gray
	FashionColor35:  "#B0C4DE", // Light Steel Blue
	FashionColor36:  "#D3D3D3", // Light Gray
	FashionColor37:  "#F0E68C", // Khaki
	FashionColor38:  "#F5F5DC", // Beige
	FashionColor39:  "#F5FFFA", // Mint Cream
	FashionColor40:  "#F0FFF0", // Honeydew
	FashionColor41:  "#F0FFFF", // Azure
	FashionColor42:  "#F8F8FF", // Ghost White
	FashionColor43:  "#F0F0F0", // Gainsboro
	FashionColor44:  "#F5F5F5", // White Smoke
	FashionColor45:  "#FAF0E6", // Linen
	FashionColor46:  "#FAFAD2", // Light Goldenrod Yellow
	FashionColor47:  "#FFE4B5", // Moccasin
	FashionColor48:  "#FFE4C4", // Bisque
	FashionColor49:  "#FFEBCD", // Blanched Almond
	FashionColor50:  "#FFF0F5", // Lavender Blush
	FashionColor51:  "#FFF5EE", // Seashell
	FashionColor52:  "#FFF8DC", // Cornsilk
	FashionColor53:  "#FFFFE0", // Light Yellow
	FashionColor54:  "#FFFFF0", // Ivory
	FashionColor55:  "#FFFFF0", // Ivory
	FashionColor56:  "#FFFAFA", // Snow
	FashionColor57:  "#FFFAF0", // Floral White
	FashionColor58:  "#FFDEAD", // Navajo White
	FashionColor59:  "#FFEFD5", // Papaya Whip
	FashionColor60:  "#FFEFD5", // Papaya Whip
	FashionColor61:  "#FFE4E1", // Misty Rose
	FashionColor62:  "#FFB6C1", // Light Pink
	FashionColor63:  "#FF69B4", // Hot Pink
	FashionColor64:  "#FF1493", // Deep Pink
	FashionColor65:  "#FF6347", // Tomato
	FashionColor66:  "#FF4500", // Orange Red
	FashionColor67:  "#FF8C00", // Dark Orange
	FashionColor68:  "#FFA07A", // Light Salmon
	FashionColor69:  "#FA8072", // Salmon
	FashionColor70:  "#E9967A", // Dark Salmon
	FashionColor71:  "#F08080", // Light Coral
	FashionColor72:  "#CD5C5C", // Indian Red
	FashionColor73:  "#B22222", // Firebrick
	FashionColor74:  "#A52A2A", // Brown
	FashionColor75:  "#8B4513", // Saddle Brown
	FashionColor76:  "#D2691E", // Chocolate
	FashionColor77:  "#FF7F50", // Coral
	FashionColor78:  "#FF4500", // Orange Red
	FashionColor79:  "#FF6347", // Tomato
	FashionColor80:  "#FF69B4", // Hot Pink
	FashionColor81:  "#FFC0CB", // Pink
	FashionColor82:  "#FFD700", // Gold
	FashionColor83:  "#ADFF2F", // Green Yellow
	FashionColor84:  "#32CD32", // Lime Green
	FashionColor85:  "#00FF00", // Lime
	FashionColor86:  "#00BFFF", // Deep Sky Blue
	FashionColor87:  "#1E90FF", // Dodger Blue
	FashionColor88:  "#4682B4", // Steel Blue
	FashionColor89:  "#5F9EA0", // Cadet Blue
	FashionColor90:  "#6495ED", // Cornflower Blue
	FashionColor91:  "#7B68EE", // Medium Slate Blue
	FashionColor92:  "#9370DB", // Medium Purple
	FashionColor93:  "#8A2BE2", // Blue Violet
	FashionColor94:  "#6A5ACD", // Slate Blue
	FashionColor95:  "#483D8B", // Dark Slate Blue
	FashionColor96:  "#2F4F4F", // Dark Slate Gray
	FashionColor97:  "#708090", // Slate Gray
	FashionColor98:  "#778899", // Light Slate Gray
	FashionColor99:  "#B0C4DE", // Light Steel Blue
	FashionColor100: "#D3D3D3", // Light Gray
	FashionColor101: "#F0E68C", // Khaki
	FashionColor102: "#F5F5DC", // Beige
	FashionColor103: "#F5FFFA", // Mint Cream
	FashionColor104: "#F0FFF0", // Honeydew
	FashionColor105: "#F0FFFF", // Azure
	FashionColor106: "#F8F8FF", // Ghost White
	FashionColor107: "#F0F0F0", // Gainsboro
	FashionColor108: "#F5F5F5", // White Smoke
	FashionColor109: "#FAF0E6", // Linen
	FashionColor110: "#FAFAD2", // Light Goldenrod Yellow
	FashionColor111: "#FFE4B5", // Moccasin
	FashionColor112: "#FFE4C4", // Bisque
	FashionColor113: "#FFEBCD", // Blanched Almond
	FashionColor114: "#FFF0F5", // Lavender Blush
	FashionColor115: "#FFF5EE", // Seashell
	FashionColor116: "#FFF8DC", // Cornsilk
	FashionColor117: "#FFFFE0", // Light Yellow
	FashionColor118: "#FFFFF0", // Ivory
	FashionColor119: "#FFFFF0", // Ivory

}

func (a FashionColor) String() string {
	return FashionColorShadeStrings[a]
}

// ランダムにFashionColorを抽選する
func RandomFashionColor() string {
	no := Random(int(SkinColorAmount))
	return FashionColor(no).String()
}
