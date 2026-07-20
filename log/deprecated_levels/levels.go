package levels

import "github.com/go-kit/log"

type Levels struct {
	logger   log.Logger
	levelKey string

	debugValue string
	infoValue  string
	warnValue  string
	errorValue string
	critValue  string
}

func New(logger log.Logger, options ...Option) Levels {
	_ = "STUB: not implemented"
	return *new(Levels)
}

func (l Levels) With(keyvals ...interface{}) Levels { _ = "STUB: not implemented"; return *new(Levels) }

func (l Levels) Debug() log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }

func (l Levels) Info() log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }

func (l Levels) Warn() log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }

func (l Levels) Error() log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }

func (l Levels) Crit() log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }

type Option func(*Levels)

func Key(key string) Option { _ = "STUB: not implemented"; return *new(Option) }

func DebugValue(value string) Option { _ = "STUB: not implemented"; return *new(Option) }

func InfoValue(value string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WarnValue(value string) Option { _ = "STUB: not implemented"; return *new(Option) }

func ErrorValue(value string) Option { _ = "STUB: not implemented"; return *new(Option) }

func CritValue(value string) Option { _ = "STUB: not implemented"; return *new(Option) }
