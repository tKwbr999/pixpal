package main

import (
	"fmt"
	"os"

	"github.com/tKwbr999/pixpal/app/pal"
)

func main() {
	content := pal.GenerateDefaultManPal()
	err := os.WriteFile("default_man.svg", []byte(content), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("SVG file generated successfully!")
}
