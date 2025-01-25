package constants

type stylesProp struct {
	Reset         string
	Bold          string
	Dim           string
	Italic        string
	Underline     string
	Blink         string
	Reverse       string
	Hidden        string
	Strikethrough string
}

var Prop = &stylesProp{
	Reset:         "reset",
	Bold:          "bold",
	Dim:           "dim",
	Italic:        "italic",
	Underline:     "underline",
	Blink:         "blink",
	Reverse:       "reverse",
	Hidden:        "hidden",
	Strikethrough: "strikethrough",
}
