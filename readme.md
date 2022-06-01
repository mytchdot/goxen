# goxen

> Create boxes in the terminal

![](screenshot.png)

_Port of 'boxen' (npm)_

## Install

```sh
go get github.com/mytchmason/goxen
```

## Usage

```go
import (
        "github.com/mytchmason/goxen"
        ...
       )

fmt.Println(goxen.Goxen("unicorn", goxen.BoxOptions{ Padding: 1 }));
/*
┌─────────────┐
│             │
│   unicorn   │
│             │
└─────────────┘
*/

fmt.Println(goxen.Goxen("unicorn", goxen.BoxOptions{ Padding: 1,  BorderStyle: "double" }));

/*

   ╔═════════════╗
   ║             ║
   ║   unicorn   ║
   ║             ║
   ╚═════════════╝

*/

```

## API

### Goxen(text, options)

#### text

Type: `string`

Text inside the box.

#### options

Type: `struct`

```go
type BoxOptions struct {
	/**
	Color of the box border.
	*/
	BorderColor string

	/**
	Style of the box border.

	@default BorderStyle.Single
	*/
	BorderStyle string

	/**
	Reduce opacity of the border.

	@default false
	*/
	DimBorder bool

	/**
	Space between the text and box border.

	@default 0
	*/
	// TODO - PaddingX int
	// TODO - PaddingY int
	PaddingTop    int
	PaddingBottom int

	/**
	Space around the box.

	@default 0
	*/
	// TODO - MarginX int
	// TODO - MarginY int
	// Margin int

	/**
	Align the text in the box based on the widest line.

	@default 'left'
	*/
	Align string
}
```

##### borderStyle

Type: `string | object`\
Default: `'single'`\
Values:

- `'single'`

```
┌───┐
│foo│
└───┘
```

- `'double'`

```
╔═══╗
║foo║
╚═══╝
```

- `'round'` (`'single'` sides with round corners)

```
╭───╮
│foo│
╰───╯
```

- `'bold'`

```
┏━━━┓
┃foo┃
┗━━━┛
```

- `'singleDouble'` (`'single'` on top and bottom, `'double'` on right and left)

```
╓───╖
║foo║
╙───╜
```

- `'doubleSingle'` (`'double'` on top and bottom, `'single'` on right and left)

```
╒═══╕
│foo│
╘═══╛
```

- `'classic'`

```
+---+
|foo|
+---+
```

##### dimBorder

Type: `boolean`\
Default: `false`

Reduce opacity of the border.

##### padding

Type: `number | object`\
Default: `0`

Space between the text and box border.

Accepts a number or an object with any of the `top`, `right`, `bottom`, `left` properties. When a number is specified, the left/right padding is 3 times the top/bottom to make it look nice.

##### margin

Type: `number | object`\
Default: `0`

Space around the box.

Accepts a number or an object with any of the `top`, `right`, `bottom`, `left` properties. When a number is specified, the left/right margin is 3 times the top/bottom to make it look nice.

##### align

Type: `string`\
Default: `'left'`\
Values: `'right'` `'center'` `'left'`

Float the box on the available terminal screen space.

## Related

- [box-cli](https://github.com/Delta456/box-cli-maker) - Better library!
