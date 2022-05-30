import { LiteralUnion } from "type-fest";
// "arrow": {
// 	"topLeft": "↘",
// 	"top": "↓",
// 	"topRight": "↙",
// 	"right": "←",
// 	"bottomRight": "↖",
// 	"bottom": "↑",
// 	"bottomLeft": "↗",
// 	"left": "→"
// }
declare namespace cliBoxes {
  /**
	Style of the box border.
	*/
  interface BoxStyle {
    readonly topLeft: string;
    readonly top: string;
    readonly topRight: string;
    readonly right: string;
    readonly bottomRight: string;
    readonly bottom: string;
    readonly bottomLeft: string;
    readonly left: string;
  }

  /**
	All box styles.
	*/
  interface Boxes {
    /**
		@example
		```
		┌────┐
		│    │
		└────┘
		```
		*/
    readonly single: BoxStyle;

    /**
		@example
		```
		╔════╗
		║    ║
		╚════╝
		```
		*/
    readonly double: BoxStyle;

    /**
		@example
		```
		╭────╮
		│    │
		╰────╯
		```
		*/
    readonly round: BoxStyle;

    /**
		@example
		```
		┏━━━━┓
		┃    ┃
		┗━━━━┛
		```
		*/
    readonly bold: BoxStyle;

    /**
		@example
		```
		╓────╖
		║    ║
		╙────╜
		```
		*/
    readonly singleDouble: BoxStyle;

    /**
		@example
		```
		╒════╕
		│    │
		╘════╛
		```
		*/
    readonly doubleSingle: BoxStyle;

    /**
		@example
		```
		+----+
		|    |
		+----+
		```
		*/
    readonly classic: BoxStyle;

    /**
		@example
		```
		↘↓↓↓↓↙
		→    ←
		↗↑↑↑↑↖
		```
		*/
    readonly arrow: BoxStyle;
  }
}

/**
Boxes for use in the terminal.
@example
```
import cliBoxes = require('cli-boxes');
console.log(cliBoxes.single);
// {
// 	topLeft: '┌',
// 	top: '─',
// 	topRight: '┐',
// 	right: '│',
// 	bottomRight: '┘',
// 	bottom: '─',
// 	bottomLeft: '└',
// 	left: '│'
// }
```
*/
declare const cliBoxes: cliBoxes.Boxes & {
  // TODO: Remove this for the next major release
  default: typeof cliBoxes;
};

declare namespace boxen {
  /**
	Characters used for custom border.

	@example
	```
	// affffb
	// e    e
	// dffffc

	const border: CustomBorderStyle = {
		topLeft: 'a',
		topRight: 'b',
		bottomRight: 'c',
		bottomLeft: 'd',
		vertical: 'e',
		horizontal: 'f'
	};
	```
	*/
  interface CustomBorderStyle extends BoxStyle {}

  /**
	Spacing used for `padding` and `margin`.
	*/
  interface Spacing {
    readonly top: number;
    readonly right: number;
    readonly bottom: number;
    readonly left: number;
  }

  interface Options {
    /**
		Color of the box border.
		*/
    readonly borderColor?: LiteralUnion<
      | "black"
      | "red"
      | "green"
      | "yellow"
      | "blue"
      | "magenta"
      | "cyan"
      | "white"
      | "gray"
      | "grey"
      | "blackBright"
      | "redBright"
      | "greenBright"
      | "yellowBright"
      | "blueBright"
      | "magentaBright"
      | "cyanBright"
      | "whiteBright",
      string
    >;

    /**
		Style of the box border.

		@default BorderStyle.Single
		*/
    readonly borderStyle?: BorderStyle | CustomBorderStyle;

    /**
		Reduce opacity of the border.

		@default false
		*/
    readonly dimBorder?: boolean;

    /**
		Space between the text and box border.

		@default 0
		*/
    readonly padding?: number | Spacing;

    /**
		Space around the box.

		@default 0
		*/
    readonly margin?: number | Spacing;

    /**
		Float the box on the available terminal screen space.

		@default 'left'
		*/
    readonly float?: "left" | "right" | "center";

    /**
		Color of the background.
		*/
    readonly backgroundColor?: LiteralUnion<
      | "black"
      | "red"
      | "green"
      | "yellow"
      | "blue"
      | "magenta"
      | "cyan"
      | "white"
      | "blackBright"
      | "redBright"
      | "greenBright"
      | "yellowBright"
      | "blueBright"
      | "magentaBright"
      | "cyanBright"
      | "whiteBright",
      string
    >;

    /**
		Align the text in the box based on the widest line.

		@default 'left'
		*/
    readonly align?: "left" | "right" | "center";
  }
}

declare const enum BorderStyle {
  Single = "single",
  Double = "double",
  Round = "round",
  Bold = "bold",
  SingleDouble = "singleDouble",
  DoubleSingle = "doubleSingle",
  Classic = "classic",
}

declare const boxen: {
  /**
	Creates a box in the terminal.

	@param text - The text inside the box.
	@returns The box.

	@example
	```
	import boxen = require('boxen');

	console.log(boxen('unicorn', {padding: 1}));
	// ┌─────────────┐
	// │             │
	// │   unicorn   │
	// │             │
	// └─────────────┘

	console.log(boxen('unicorn', {padding: 1, margin: 1, borderStyle: 'double'}));
	//
	// ╔═════════════╗
	// ║             ║
	// ║   unicorn   ║
	// ║             ║
	// ╚═════════════╝
	//
	```
	*/
  (text: string, options?: boxen.Options): string;

  /**
	Border styles from [`cli-boxes`](https://github.com/sindresorhus/cli-boxes).
	*/
  BorderStyle: typeof BorderStyle;
};

export = boxen;
("use strict");
const stringWidth = require("string-width");
const chalk = require("chalk");
const widestLine = require("widest-line");
const cliBoxes = require("cli-boxes");
const camelCase = require("camelcase");
const ansiAlign = require("ansi-align");
const termSize = require("term-size");

const getObject = (detail) => {
  let object;

  if (typeof detail === "number") {
    object = {
      top: detail,
      right: detail * 3,
      bottom: detail,
      left: detail * 3,
    };
  } else {
    object = {
      top: 0,
      right: 0,
      bottom: 0,
      left: 0,
      ...detail,
    };
  }

  return object;
};

const getBorderChars = (borderStyle) => {
  const sides = [
    "topLeft",
    "topRight",
    "bottomRight",
    "bottomLeft",
    "vertical",
    "horizontal",
  ];

  let chararacters;

  if (typeof borderStyle === "string") {
    chararacters = cliBoxes[borderStyle];

    if (!chararacters) {
      throw new TypeError(`Invalid border style: ${borderStyle}`);
    }
  } else {
    for (const side of sides) {
      if (!borderStyle[side] || typeof borderStyle[side] !== "string") {
        throw new TypeError(`Invalid border style: ${side}`);
      }
    }

    chararacters = borderStyle;
  }

  return chararacters;
};

const isHex = (color) => color.match(/^#[0-f]{3}(?:[0-f]{3})?$/i);
const isColorValid = (color) =>
  typeof color === "string" && (chalk[color] || isHex(color));
const getColorFn = (color) => (isHex(color) ? chalk.hex(color) : chalk[color]);
const getBGColorFn = (color) =>
  isHex(color) ? chalk.bgHex(color) : chalk[camelCase(["bg", color])];

module.exports = (text, options) => {
  options = {
    padding: 0,
    borderStyle: "single",
    dimBorder: false,
    align: "left",
    float: "left",
    ...options,
  };

  if (options.borderColor && !isColorValid(options.borderColor)) {
    throw new Error(`${options.borderColor} is not a valid borderColor`);
  }

  if (options.backgroundColor && !isColorValid(options.backgroundColor)) {
    throw new Error(
      `${options.backgroundColor} is not a valid backgroundColor`
    );
  }

  const chars = getBorderChars(options.borderStyle);
  const padding = getObject(options.padding);
  const margin = getObject(options.margin);

  const colorizeBorder = (border) => {
    const newBorder = options.borderColor
      ? getColorFn(options.borderColor)(border)
      : border;
    return options.dimBorder ? chalk.dim(newBorder) : newBorder;
  };

  const colorizeContent = (content) =>
    options.backgroundColor
      ? getBGColorFn(options.backgroundColor)(content)
      : content;

  text = ansiAlign(text, { align: options.align });

  const NL = "\n";
  const PAD = " ";

  let lines = text.split(NL);

  if (padding.top > 0) {
    lines = new Array(padding.top).fill("").concat(lines);
  }

  if (padding.bottom > 0) {
    lines = lines.concat(new Array(padding.bottom).fill(""));
  }

  const contentWidth = widestLine(text) + padding.left + padding.right;
  const paddingLeft = PAD.repeat(padding.left);
  const { columns } = termSize();
  let marginLeft = PAD.repeat(margin.left);

  if (options.float === "center") {
    const padWidth = Math.max((columns - contentWidth) / 2, 0);
    marginLeft = PAD.repeat(padWidth);
  } else if (options.float === "right") {
    const padWidth = Math.max(columns - contentWidth - margin.right - 2, 0);
    marginLeft = PAD.repeat(padWidth);
  }

  const horizontal = chars.horizontal.repeat(contentWidth);
  const top = colorizeBorder(
    NL.repeat(margin.top) +
      marginLeft +
      chars.topLeft +
      horizontal +
      chars.topRight
  );
  const bottom = colorizeBorder(
    marginLeft +
      chars.bottomLeft +
      horizontal +
      chars.bottomRight +
      NL.repeat(margin.bottom)
  );
  const side = colorizeBorder(chars.vertical);

  const middle = lines
    .map((line) => {
      const paddingRight = PAD.repeat(
        contentWidth - stringWidth(line) - padding.left
      );
      return (
        marginLeft +
        side +
        colorizeContent(paddingLeft + line + paddingRight) +
        side
      );
    })
    .join(NL);

  return top + NL + middle + NL + bottom;
};

module.exports._borderStyles = cliBoxes;
