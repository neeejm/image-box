package main

import (
	"fmt"
	"io/ioutil"
)

const LINE_WIDTH = 8
const LINE_COLOR = "#0488D0"

func main() {
	fmt.Println("__main__")
	// box := ut.Box{
	// 	TopRow:    0.09726668,
	// 	RightCol:  0.7593585,
	// 	BottomRow: 0.84501964,
	// 	LeftCol:   0.33775616,
	// }

	// ut.DrawBox("face.png", box)

	data, err := ioutil.ReadFile("out.png")
	if err != nil {
		panic(err)
	}
	fmt.Println("byte slice data", data)
}
