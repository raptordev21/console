# console
Style console/terminal string

# Log Options
```go
package console

type LogOptions struct {
	msg             string
	color           string
	bgColor         string
	isBold          bool
	isUnderline     bool
	isItalic        bool
	isDim           bool
	isBlink         bool
	isReverse       bool
	isHidden        bool
	isStrikethrough bool
	isBanner        bool
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
		msg:    "This is bold with blue text color",
		color:  "blue",
		isBold: true,
	})

  err := "Some Error" 
  console.Log(console.LogOptions{
		msg:    fmt.Sprint("Red Color Text:", err),
		color:  "red",
	})

  app := "My App"
  console.Log(console.LogOptions{
		msg:      fmt.Sprintf("Welcome to %v", app),
		bgColor:  "bgCyan",
		isBold:   true,
		isBanner: true,
	})

  console.Log(console.LogOptions{
		msg:   "Text is underlined & color is magenta",
		color: "magenta",
    isUnderline: true,
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
