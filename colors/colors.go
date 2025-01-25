package colors

type textColor struct {
	Black     string
	Red       string
	Green     string
	Yellow    string
	Blue      string
	Magenta   string
	Cyan      string
	White     string
	BgBlack   string
	BgRed     string
	BgGreen   string
	BgYellow  string
	BgBlue    string
	BgMagenta string
	BgCyan    string
	BgWhite   string
}

var Color = &textColor{
	Black:     "black",
	Red:       "red",
	Green:     "green",
	Yellow:    "yellow",
	Blue:      "blue",
	Magenta:   "magenta",
	Cyan:      "cyan",
	White:     "white",
	BgBlack:   "bgBlack",
	BgRed:     "bgRed",
	BgGreen:   "bgGreen",
	BgYellow:  "bgYellow",
	BgBlue:    "bgBlue",
	BgMagenta: "bgMagenta",
	BgCyan:    "bgCyan",
	BgWhite:   "bgWhite",
}
