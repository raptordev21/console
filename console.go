package console

import (
	"fmt"

	"github.com/raptordev21/console/v1/colors"
	"github.com/raptordev21/console/v1/internal/constants"
	"github.com/raptordev21/console/v1/internal/helpers"
	"github.com/raptordev21/console/v1/internal/styles"
)

type LogOptions struct {
	Msg             string
	Color           string
	BgColor         string
	IsBold          bool
	IsUnderline     bool
	IsItalic        bool
	IsDim           bool
	IsBlink         bool
	IsReverse       bool
	IsHidden        bool
	IsStrikethrough bool
	IsBanner        bool
}

func Log(options LogOptions) {
	var spaces string
	if options.IsBanner {
		spaces = styles.Get(options.BgColor) + `      ` + helpers.GenerateSpaces(options.Msg) + `      ` + styles.Get(constants.Prop.Reset)
		options.Msg = `      ` + options.Msg + `      ` + styles.Get(constants.Prop.Reset)
	} else {
		options.Msg = options.Msg + styles.Get(constants.Prop.Reset)
	}

	switch options.Color {
	case colors.Color.Black:
		options.Msg = styles.Get(options.Color) + options.Msg
	case colors.Color.Red:
		options.Msg = styles.Get(options.Color) + options.Msg
	case colors.Color.Green:
		options.Msg = styles.Get(options.Color) + options.Msg
	case colors.Color.Yellow:
		options.Msg = styles.Get(options.Color) + options.Msg
	case colors.Color.Blue:
		options.Msg = styles.Get(options.Color) + options.Msg
	case colors.Color.Magenta:
		options.Msg = styles.Get(options.Color) + options.Msg
	case colors.Color.Cyan:
		options.Msg = styles.Get(options.Color) + options.Msg
	case colors.Color.White:
		options.Msg = styles.Get(options.Color) + options.Msg
	default:
	}

	switch options.BgColor {
	case colors.Color.BgBlack:
		options.Msg = styles.Get(options.BgColor) + options.Msg
	case colors.Color.BgRed:
		options.Msg = styles.Get(options.BgColor) + options.Msg
	case colors.Color.BgGreen:
		options.Msg = styles.Get(options.BgColor) + options.Msg
	case colors.Color.BgYellow:
		options.Msg = styles.Get(options.BgColor) + options.Msg
	case colors.Color.BgBlue:
		options.Msg = styles.Get(options.BgColor) + options.Msg
	case colors.Color.BgMagenta:
		options.Msg = styles.Get(options.BgColor) + options.Msg
	case colors.Color.BgCyan:
		options.Msg = styles.Get(options.BgColor) + options.Msg
	case colors.Color.BgWhite:
		options.Msg = styles.Get(options.BgColor) + options.Msg
	default:
	}

	if options.IsBold {
		options.Msg = styles.Get(constants.Prop.Bold) + options.Msg
	}
	if options.IsUnderline {
		options.Msg = styles.Get(constants.Prop.Underline) + options.Msg
	}
	if options.IsItalic {
		options.Msg = styles.Get(constants.Prop.Italic) + options.Msg
	}
	if options.IsDim {
		options.Msg = styles.Get(constants.Prop.Dim) + options.Msg
	}
	if options.IsBlink {
		options.Msg = styles.Get(constants.Prop.Blink) + options.Msg
	}
	if options.IsReverse {
		options.Msg = styles.Get(constants.Prop.Reverse) + options.Msg
	}
	if options.IsHidden {
		options.Msg = styles.Get(constants.Prop.Hidden) + options.Msg
	}
	if options.IsStrikethrough {
		options.Msg = styles.Get(constants.Prop.Strikethrough) + options.Msg
	}

	if options.IsBanner {
		fmt.Println(spaces)
		fmt.Println(options.Msg)
		fmt.Print(spaces)
	} else {
		fmt.Print(options.Msg)
	}
	fmt.Println()
}

func Info(str string) {
	str = styles.Get(colors.Color.Blue) + str + styles.Get(constants.Prop.Reset)
	fmt.Print(str)
	fmt.Println()
}
func Success(str string) {
	str = styles.Get(colors.Color.Green) + str + styles.Get(constants.Prop.Reset)
	fmt.Print(str)
	fmt.Println()
}
func Error(str string) {
	str = styles.Get(colors.Color.Red) + str + styles.Get(constants.Prop.Reset)
	fmt.Print(str)
	fmt.Println()
}
func Warn(str string) {
	str = styles.Get(colors.Color.Yellow) + str + styles.Get(constants.Prop.Reset)
	fmt.Print(str)
	fmt.Println()
}
