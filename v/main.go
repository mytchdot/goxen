package main

import (
	"fmt"

	"github.com/mytchmason/goxen"
)

func main() {
	// fmt.Println(goxen.Goxen("Hello World!"))
	fmt.Println(goxen.Goxen("aofaejbfbjael", goxen.BoxOptions{
		BorderColor:   "White",
		BorderStyle:   "Double",
		DimBorder:     true,
		PaddingTop:    2,
		PaddingBottom: 2,
		Align:         "center",
	}))
}
