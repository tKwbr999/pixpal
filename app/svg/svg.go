package svg

import (
	"fmt"
)

const (
	UNIT         int = 5
	HEIGHT_UNITS int = 30
	WIDTH_UNITS  int = 10
)

const SVG_HEADER = `<svg xmlns="http://www.w3.org/2000/svg" width="%[1]d" height="%[2]d" viewBox="0 0 %[1]d %[2]d">
  <metadata>
    <rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
             xmlns:dc="http://purl.org/dc/elements/1.1/">
      <rdf:Description>
        <dc:title>PixPal Character</dc:title>
        <dc:creator>PixPal Generator</dc:creator>
        <dc:rights>Open Source - MIT License</dc:rights>
        <dc:date>2023</dc:date>
      </rdf:Description>
    </rdf:RDF>
  </metadata>
  `
const RECT_TAG_FORMAT string = `<rect x="%d" y="%d" width="%d" height="%d" fill="%s"/>`

func Header() string {
	width := WIDTH_UNITS * UNIT
	height := HEIGHT_UNITS * UNIT
	header := fmt.Sprintf(SVG_HEADER, width, height)
	return header
}

func Rect(x, y int, color string) string {
	rect := fmt.Sprintf(RECT_TAG_FORMAT, x, y, UNIT, UNIT, color)
	return rect
}

func Pos(p int) int {
	pos := (p - 1) * UNIT
	return pos
}
