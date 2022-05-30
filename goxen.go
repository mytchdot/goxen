package goxen

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/term"
)

const (
	NL  = "\n"
	PAD = " "
)

func fillSlice(n int, s string) []string {
	var slice []string
	for i := 0; i < n; i++ {
		slice = append(slice, s)
	}
	return slice
}

func getObject(detail int) map[string]int {
	object := map[string]int{
		"top":    0,
		"right":  0,
		"bottom": 0,
		"left":   0,
	}

	if detail > -1 {
		object["top"] = detail
		object["right"] = detail * 3
		object["bottom"] = detail
		object["left"] = detail * 3
	}

	return object
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
	// if options.backgroundColor {
	// 	getBGColorFn(options.backgroundColor)(content)
	// }
	return content
}

func colorizeBorder(border string, options BoxOptions) string {
	newBorder := options

	if options.BorderColor != "" {
		newBorder.BorderColor = getColorFn(options.BorderColor)
	} else {
		newBorder.BorderColor = border
	}
	if options.DimBorder {
		return Colors["Dim"] + newBorder.BorderColor + Colors["Reset"]
	} else {
		return newBorder.BorderColor
	}
}

func stringWidth(str string) int {

	width := 0
	// ambiguousCharacterWidth := 1

	// strStrippedAnsi := stripAnsi(str)

	// if len(strStrippedAnsi) < 1 {
	// 	return 0
	// }

	// strNoEmojis := EmojiRegexp.ReplaceAll(str, `  `)

	// for _, char := range strNoEmojis {
	// 	codePoint, _ := utf8.DecodeRuneInString(string(char)[0])

	// 	// Ignore control characters
	// 	if codePoint <= 0x1F || (codePoint >= 0x7F && codePoint <= 0x9F) {
	// 		continue
	// 	}
	// 	// Ignore combining characters
	// 	if codePoint >= 0x300 && codePoint <= 0x36F {
	// 		continue
	// 	}

	// code = eastAsianWidth(char) - See TODO below
	// }

	// TODO - east asian width support
	// for _, char := range strStrippedAnsi {}

	// 	const code = eastAsianWidth(character);
	// 	switch (code) {
	// 		case 'F':
	// 		case 'W':
	// 			width += 2;
	// 			break;
	// 		case 'A':
	// 			width += ambiguousCharacterWidth;
	// 			break;
	// 		default:
	// 			width += 1;
	// 	}
	// }

	return width + 1 // default on switch, setting here as to emulate for future addition
}

func widestLine(str string) int {
	lineWidth := 0

	splitStr := strings.Split(str, NL)
	for _, line := range splitStr {
		width := stringWidth(line)
		max := math.Max(float64(lineWidth), float64(width))
		lineWidth = int(max)
	}
	return lineWidth
}

type AnsiAlignOptions struct {
	Align string
	Split string
	Pad   string
}

func halfDiff(maxWidth int, curWidth int) int {
	newNum := (maxWidth - curWidth) / 2
	return int(math.Floor(float64(newNum)))
}
func fullDiff(maxWidth int, curWidth int) int {
	newNum := maxWidth - curWidth
	return int(math.Floor(float64(newNum)))
}

func ansiAlign(text string, align string) string {
	if text == "" {
		return text // noop
	}
	if align == "left" {
		return text // noop
	}

	pad := ` `
	var width int

	maxWidth := 0
	// maps over text
	// loop var = loop var stringified(noop?)
	// width var = stringWidth -- rune width??
	// maxWidth var = math.max(maxWidth, width)
	// append {loopvar, widthvar, maxwidth}
	// map again with that new map/object being the loop item
	// xx = witdhDiffFn(maxWidth, loopitem.width) + 1
	//  create new slice from xx
	// join new slice by opts.pad,
	// new slice + loopitemobj.string
	// if resutls of all that mappin

	var newString string
	var stringAndWidthMap []map[string]string

	for _, char := range text {
		str := string(char)
		width = stringWidth(str)
		maxWidth = int(math.Max(float64(maxWidth), float64(width)))

		stringAndWidthMap = append(stringAndWidthMap, map[string]string{
			"str":   str,
			"width": strconv.Itoa(width),
		})
	}

	for _, item := range stringAndWidthMap {

		var newWidthDiff string

		widthAsInt, _ := strconv.Atoi(item["width"])

		if align != "right" {
			newWidthDiff = strconv.Itoa(halfDiff(maxWidth, widthAsInt) + 1)
		} else {
			newWidthDiff = strconv.Itoa(fullDiff(maxWidth, widthAsInt) + 1)
		}

		asSlice := []string{newWidthDiff}
		sliceAsStrng := strings.Join(asSlice[:], pad)
		newString = sliceAsStrng + item["str"]

	}
	return newString

}

func Goxen(message string, options BoxOptions) string {

	// defaultOptions := BoxOptions{
	// 	BorderColor:   "White",
	// 	BorderStyle:   "Single",
	// 	DimBorder:     false,
	// 	PaddingTop:    0,
	// 	PaddingBottom: 0,
	// 	Align:         "left",
	// }

	if options.BorderColor != "" && !isColorValid(options.BorderColor) {
		log.Fatal("Invalid border color: " + options.BorderColor)
	}

	chars := getBorderChars(options.BorderStyle)
	paddingTop := options.PaddingTop       // See TODO - Padding [X,[Y]]
	paddingBottom := options.PaddingBottom // See TODO - Padding [X,[Y]]
	// margin := options.Margin   // See TODO - Margin[X,[Y]]

	message = ansiAlign(message, options.Align)

	lines := strings.Split(message, NL)
	if paddingTop > 0 {
		lines = append(fillSlice(paddingTop, ""), lines...)
	}
	if paddingBottom > 0 {
		lines = append(lines, fillSlice(paddingTop, "")...)
	}

	// contentWidth = widestLine(text) + padding.left + padding.right;
	contentWidth := widestLine(message)

	// paddingLeft := strings.Repeat(PAD, options.paddingLeft)
	paddingLeft := strings.Repeat(PAD, 0)
	// paddingRight := strings.Repeat(PAD, 0)

	columns, _, err := term.GetSize(0)
	if err != nil {
		log.Fatal("Failed to get terminal size: ", err)
	}

	//   marginLeft := strings.Repeat(PAD, margin.left);
	marginLeft := strings.Repeat(PAD, 0)
	marginRight := 0
	float := "left" // TODO

	if float == "center" {
		// padWidth := math.Max((columns-contentWidth)/2, 0)
		padWidth := math.Max(float64(columns-contentWidth)/2, 0)
		marginLeft = strings.Repeat(PAD, int(padWidth))
	} else if float == "right" {
		padWidth := math.Max(float64(columns-contentWidth-marginRight)-2, 0)
		marginLeft = strings.Repeat(PAD, int(padWidth))
	}

	// TODO - marginRight

	// horizontal := strings.Repeat(chars.Horizontal, contentWidth+options.paddingLeft+paddingRight)
	horizontal := strings.Repeat(chars.Horizontal, contentWidth+0+0)

	marginTop := 0
	marginBottom := 0

	top := colorizeBorder(strings.Repeat(NL, marginTop)+marginLeft+chars.TopLeft+horizontal+chars.TopRight, options)
	bottom := colorizeBorder(marginLeft+chars.BottomLeft+horizontal+chars.BottomRight+strings.Repeat(NL, marginBottom), options)

	side := colorizeBorder(chars.Vertical, options)

	// middle
	var middle string

	// map lines
	// loop var = line
	// paddingright
	// reutnr str = marginleft + side + colorizeContent(paddingLeft + line + paddingRight) + side
	for _, line := range lines {
		// paddingRight := strings.Repeat(PAD, contentWidth-stringWidth(line)-paddingLeft)
		padRight := strings.Repeat(PAD, contentWidth-stringWidth(line)-0)
		middle = marginLeft + side + colorizeContent(paddingLeft+line+padRight, options) + side
	}

	return top + NL + middle + NL + bottom
}
