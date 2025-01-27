package styles

import (
	"github.com/raptordev21/console/v1/colors"
	"github.com/raptordev21/console/v1/internal/constants"
)

func Get(key string) string {
	styles := make(map[string]string)
	styles[constants.Prop.Reset] = "\033[0m"
	styles[constants.Prop.Bold] = "\033[1m"
	styles[constants.Prop.Dim] = "\033[2m"
	styles[constants.Prop.Italic] = "\033[3m"
	styles[constants.Prop.Underline] = "\033[4m"
	styles[constants.Prop.Blink] = "\033[5m"
	styles[constants.Prop.Reverse] = "\033[7m"
	styles[constants.Prop.Hidden] = "\033[8m"
	styles[constants.Prop.Strikethrough] = "\033[9m"
	styles[colors.Color.Black] = "\033[30m"
	styles[colors.Color.Red] = "\033[31m"
	styles[colors.Color.Green] = "\033[32m"
	styles[colors.Color.Yellow] = "\033[33m"
	styles[colors.Color.Blue] = "\033[34m"
	styles[colors.Color.Magenta] = "\033[35m"
	styles[colors.Color.Cyan] = "\033[36m"
	styles[colors.Color.White] = "\033[37m"
	styles[colors.Color.BgBlack] = "\033[40m"
	styles[colors.Color.BgRed] = "\033[41m"
	styles[colors.Color.BgGreen] = "\033[42m"
	styles[colors.Color.BgYellow] = "\033[43m"
	styles[colors.Color.BgBlue] = "\033[44m"
	styles[colors.Color.BgMagenta] = "\033[45m"
	styles[colors.Color.BgCyan] = "\033[46m"
	styles[colors.Color.BgWhite] = "\033[47m"

	return styles[key]
}
