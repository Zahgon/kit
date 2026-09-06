package term

import (
	"io"

	"github.com/go-kit/log"
	"github.com/go-kit/log/term"
)

type Color = term.Color

const (
	Default = term.Default

	Black       = term.Black
	DarkRed     = term.DarkRed
	DarkGreen   = term.DarkGreen
	Brown       = term.Brown
	DarkBlue    = term.DarkBlue
	DarkMagenta = term.DarkMagenta
	DarkCyan    = term.DarkCyan
	Gray        = term.Gray

	DarkGray = term.DarkGray
	Red      = term.Red
	Green    = term.Green
	Yellow   = term.Yellow
	Blue     = term.Blue
	Magenta  = term.Magenta
	Cyan     = term.Cyan
	White    = term.White
)

type FgBgColor = term.FgBgColor

func NewColorLogger(w io.Writer, newLogger func(io.Writer) log.Logger, color func(keyvals ...interface{}) FgBgColor) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}
