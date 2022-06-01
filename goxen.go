package goxen

import (
	"log"
	"math"
	"regexp"
	"strings"

	"github.com/mattn/go-runewidth"
	"golang.org/x/term"
)

type AlignOptions struct {
	Align string
	Split string
	Pad   string
}

func halfDiff(maxWidth int, curWidth int) int {
	newNum := (maxWidth - curWidth) / 2
	return int(float64(newNum))
}
func fullDiff(maxWidth int, curWidth int) int {
	newNum := maxWidth - curWidth
	return int(float64(newNum))
}
func Align(text string, opts AlignOptions) string {
	if text == "" {
		return text
	}

	if opts.Align != "" {
		opts.Align = "center"
	}
	if opts.Split != "" {
		opts.Split = "\n"
	}
	if opts.Pad != "" {
		opts.Pad = ` `
	}

	// short-circuit `align: 'left'` as no-op
	if opts.Align == "left" {
		return text
	}

	newText := strings.Split(text, opts.Split)

	maxWidth := 0

	var width int
	var widthDiff int
	var strSlice []string

	for _, line := range newText {
		str := string(line)
		width = runewidth.StringWidth(str)
		maxWidth = int(math.Max(float64(width), float64(maxWidth)))
		if opts.Align != "right" {
			widthDiff = halfDiff(maxWidth, width) + 1
		} else {
			widthDiff = fullDiff(maxWidth, width) + 1
		}

		mfrSlice := strings.Join(fillSlice(widthDiff, opts.Pad), "") + str
		strSlice = append(strSlice, mfrSlice)
	}
	return strings.Join(strSlice, opts.Split)
}

func fillSlice(n int, s string) []string {
	var slice []string
	for i := 0; i < n; i++ {
		slice = append(slice, s)
	}
	return slice
}

// TODO - Allow for custom Boxes, pass whole box struct item...??
func getBorderChars(borderStyle string) Box {
	chararacters := Boxes[borderStyle]
	// if chararacters !=  {
	// 	log.Fatal("Invalid border style: " + borderStyle)
	// }
	return chararacters
}

func isHex(char string) bool {
	return regexp.MustCompile("/^#[0-f]{3}(?:[0-f]{3})?$/i").MatchString(char)
}

func isColorValid(color string) bool {
	_, ok := Colors[color]
	return isHex(color) || ok
}

func getColorFn(color string) string {
	if isHex(color) {
		return color
	}
	return Colors[color]
}

func colorizeContent(content string, options BoxOptions) string {
	// TODO - Colorize content
	// if options.backgroundColor {
	// 	getBGColorFn(options.backgroundColor)(content)
	// }
	return content
}

func colorizeBorder(border string, options BoxOptions) string {
	newBorder := options
	if options.BorderColor != "" {
		newBorder.BorderColor = getColorFn(options.BorderColor) + border + Colors["Reset"]
	} else {
		newBorder.BorderColor = border
	}
	if options.DimBorder {
		return Colors["Dim"] + newBorder.BorderColor + Colors["Reset"]
	} else {
		return newBorder.BorderColor
	}

}

func widestLine(str string) int {
	lineWidth := 0

	splitStr := strings.Split(str, NL)
	for _, line := range splitStr {
		width := runewidth.StringWidth(line)

		max := math.Max(float64(lineWidth), float64(width))
		lineWidth = int(max)
	}

	return lineWidth
}

func Goxen(message string, options BoxOptions) string {
	if options.BorderColor != "" && !isColorValid(options.BorderColor) {
		log.Fatal("Invalid border color: " + options.BorderColor)
	}

	chars := getBorderChars(options.BorderStyle)
	paddingTop := options.PaddingTop       // See TODO - Padding [X,[Y]]
	paddingBottom := options.PaddingBottom // See TODO - Padding [X,[Y]]

	message = Align(message, AlignOptions{
		Align: options.Align,
	})

	lines := strings.Split(message, NL)
	if paddingTop > 0 {
		lines = append(fillSlice(paddingTop, ""), lines...)
	}
	if paddingBottom > 0 {
		lines = append(lines, fillSlice(paddingTop, "")...)
	}
	contentWidth := widestLine(message)

	// paddingLeft := strings.Repeat(PAD, options.paddingLeft)
	paddingLeft := strings.Repeat(PAD, 0)

	columns, _, err := term.GetSize(0)
	if err != nil {
		log.Fatal("Failed to get terminal size: ", err)
	}

	marginLeft := strings.Repeat(PAD, 0)
	marginRight := 0 // TODO
	float := "left"  // TODO

	if float == "center" {
		padWidth := math.Max(float64(columns-contentWidth)/2, 0)
		marginLeft = strings.Repeat(PAD, int(padWidth))
	} else if float == "right" {
		padWidth := math.Max(float64(columns-contentWidth-marginRight)-2, 0)
		marginLeft = strings.Repeat(PAD, int(padWidth))
	}

	// horizontal := strings.Repeat(chars.Horizontal, contentWidth+options.paddingLeft+paddingRight)
	horizontal := strings.Repeat(chars.Horizontal, contentWidth+0+0)

	marginTop := 0
	marginBottom := 0
	top := colorizeBorder(strings.Repeat(NL, marginTop)+marginLeft+chars.TopLeft+horizontal+chars.TopRight, options)
	bottom := colorizeBorder(marginLeft+chars.BottomLeft+horizontal+chars.BottomRight+strings.Repeat(NL, marginBottom), options)

	side := colorizeBorder(chars.Vertical, options)

	var middle []string

	for _, line := range lines {
		// paddingRight := strings.Repeat(PAD, contentWidth-stringWidth(line)-paddingLeft)
		padRight := strings.Repeat(PAD, contentWidth-runewidth.StringWidth(line)-0)
		middle = append(middle, marginLeft+side+colorizeContent(paddingLeft+line+padRight, options)+side)
	}

	return top + NL + strings.Join(middle, NL) + NL + bottom
}
