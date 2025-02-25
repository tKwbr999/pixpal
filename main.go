package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tKwbr999/pixpal/app/pal"
)

func main() {
	var output string
	flag.StringVar(&output, "o", "default_man.svg", "output file name")
	flag.Parse()

	content := pal.GenerateDefaultManPal()
	os.WriteFile(output, []byte(content), 0644)

	fmt.Println("SVG file generated successfully!")
}
