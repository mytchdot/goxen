package main

import (
	"fmt"

	"github.com/mytchmason/goxen"
)

func main() {
	// fmt.Println(goxen.Goxen("Hello World!"))
	fmt.Println(goxen.Goxen("aeklfna;lfjbea", goxen.BoxOptions{
		BorderColor:   "Cyan",
		BorderStyle:   "Single",
		DimBorder:     false,
		PaddingTop:    2,
		PaddingBottom: 2,
		Align:         "center",
	}))
}
