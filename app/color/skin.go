package color

import "fmt"

type SkinColor int

const (
	SkinColor1  SkinColor = iota // #fcf8f0 Light Skin
	SkinColor2                   // #f7eddd Fair Skin
	SkinColor3                   // #fce6d2 Light Peach
	SkinColor4                   // #f2c6b6 Peach
	SkinColor5                   // #f0ddc3 Medium Light Skin
	SkinColor6                   // #e8b0a0 Light Rose
	SkinColor7                   // #e3bd8e Medium Skin
	SkinColor8                   // #d69b61 Medium Tan Skin
	SkinColor9                   // #cd8142 Tan Skin
	SkinColor10                  // #bf6b37 Olive Skin
	SkinColor11                  // #9f552f Dark Tan Skin
	SkinColor12                  // #80452c Dark Skin
	SkinColor13                  // #683a26 Deep Dark Skin
	SkinColor14                  // #371d13 Very Dark Skin
	SkinColor15                  // #d99a8e Rose
	SkinColor16                  // #c88b7a Medium Rose
	SkinColor17                  // #b77c66 Dark Rose
	SkinColor18                  // #a66e54 Light Brown
	SkinColor19                  // #9b5e4a Brown
	SkinColor20                  // #8f4e3e Dark Brown
	SkinColor21                  // #7b3e32 Deep Brown
	SkinColor22                  // #6f2e26 Very Deep Brown
	SkinColor23                  // #f4d6c1 Light Almond
	SkinColor24                  // #e8c1a0 Almond
	SkinColor25                  // #d6a78a Medium Almond
	SkinColor26                  // #c18e6f Dark Almond
	SkinColor27                  // #b07b5c Deep Almond
	SkinColor28                  // #a06a4b Very Deep Almond
	SkinColor29                  // #f2e1d5 Light Beige
	SkinColor30                  // #e6cfc0 Beige
	SkinColor31                  // #d1b3a0 Medium Beige
	SkinColor32                  // #c19a80 Dark Beige
	SkinColorAmount
)

var SkinColorShadeStrings = map[SkinColor]string{
	SkinColor1:  "#fcf8f0", // Light Skin
	SkinColor2:  "#f7eddd", // Fair Skin
	SkinColor3:  "#fce6d2", // Light Peach
	SkinColor4:  "#f2c6b6", // Peach
	SkinColor5:  "#f0ddc3", // Medium Light Skin
	SkinColor6:  "#e8b0a0", // Light Rose
	SkinColor7:  "#e3bd8e", // Medium Skin
	SkinColor8:  "#d69b61", // Medium Tan Skin
	SkinColor9:  "#cd8142", // Tan Skin
	SkinColor10: "#bf6b37", // Olive Skin
	SkinColor11: "#9f552f", // Dark Tan Skin
	SkinColor12: "#80452c", // Dark Skin
	SkinColor13: "#683a26", // Deep Dark Skin
	SkinColor14: "#371d13", // Very Dark Skin
	SkinColor15: "#d99a8e", // Rose
	SkinColor16: "#c88b7a", // Medium Rose
	SkinColor17: "#b77c66", // Dark Rose
	SkinColor18: "#a66e54", // Light Brown
	SkinColor19: "#9b5e4a", // Brown
	SkinColor20: "#8f4e3e", // Dark Brown
	SkinColor21: "#7b3e32", // Deep Brown
	SkinColor22: "#6f2e26", // Very Deep Brown
	SkinColor23: "#f4d6c1", // Light Almond
	SkinColor24: "#e8c1a0", // Almond
	SkinColor25: "#d6a78a", // Medium Almond
	SkinColor26: "#c18e6f", // Dark Almond
	SkinColor27: "#b07b5c", // Deep Almond
	SkinColor28: "#a06a4b", // Very Deep Almond
	SkinColor29: "#f2e1d5", // Light Beige
	SkinColor30: "#e6cfc0", // Beige
	SkinColor31: "#d1b3a0", // Medium Beige
	SkinColor32: "#c19a80", // Dark Beige
}

func (a SkinColor) String() string {
	return SkinColorShadeStrings[a]
}

// ランダムにSkinColorを抽選する
func RandomSkinColor() string {
	no := Random(int(SkinColorAmount - 1))
	return SkinColor(no).String()
}

// Increase the shade of the specified hex color
func DarkenColor(hex string) string {
	const factor = 0.95 // Fixed value to slightly darken the color
	// Convert hex to RGB
	r, g, b := hexToRGB(hex)

	// Apply the factor to darken the color
	r = clamp(int(float64(r) * factor))
	g = clamp(int(float64(g) * factor))
	b = clamp(int(float64(b) * factor))

	// Convert back to hex
	return rgbToHex(r, g, b)
}

// Helper function to convert hex to RGB
func hexToRGB(hex string) (int, int, int) {
	var r, g, b int
	fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	return r, g, b
}

// Helper function to convert RGB to hex
func rgbToHex(r, g, b int) string {
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

// Helper function to clamp values between 0 and 255
func clamp(value int) int {
	if value < 0 {
		return 0
	} else if value > 255 {
		return 255
	}
	return value
}
