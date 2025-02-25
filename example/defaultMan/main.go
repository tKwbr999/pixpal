package main

import (
	"fmt"
	"os"

	"github.com/tKwbr999/pixpal/app/pal"
)

func main() {
	content := pal.GenerateDefaultManPal()
	os.WriteFile("default_man.svg", []byte(content), 0644)

	fmt.Println("SVG file generated successfully!")
}
