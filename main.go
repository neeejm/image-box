package main

import (
	"fmt"

	ut "github.com/neeejm/image-box/utils"
)

const LINE_WIDTH = 8
const LINE_COLOR = "#0488D0"

func main() {
	fmt.Println("__main__")
	box := ut.Box{
		TopRow:    0.09726668,
		RightRow:  0.7593585,
		BottomRow: 0.84501964,
		LeftRow:   0.33775616,
		LineWidth: 7.5,
		LineColor: "#0488D0", // shade of blue
	}

	ut.DrawBox(box)
}
