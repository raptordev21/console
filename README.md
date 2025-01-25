# console
Style console/terminal string

# Log Options
```go
package console

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
```

## Examples using LogOptions:
Note: Every `LogOptions` combination may not work. This depends on terminal.
```go
package main

import (
  "github.com/raptordev21/console"
)

func main() {
  console.Log(console.LogOptions{
		Msg:    "This is bold with blue text color",
		Color:  "blue",
		IsBold: true,
	})

  err := "Some Error" 
  console.Log(console.LogOptions{
		Msg:    fmt.Sprint("Red Color Text:", err),
		Color:  "red",
	})

  app := "My App"
  console.Log(console.LogOptions{
		Msg:      fmt.Sprintf("Welcome to %v", app),
		BgColor:  "bgCyan",
		IsBold:   true,
		IsBanner: true,
	})

  console.Log(console.LogOptions{
		Msg:   "Text is underlined & color is magenta",
		Color: "magenta",
    IsUnderline: true,
	})
}
```

## Examples using other functions:
```go
package main

import (
  "github.com/raptordev21/console"
)

func main() {
  console.Info("Blue text color")
  console.Success("Green text color")
  console.Error("Red text color")
  console.Warn("Yellow text color")
}
```
