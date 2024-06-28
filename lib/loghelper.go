package lib

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	DiscardLog int = iota
	ConsoleLog
	JSONLog

	logMode  = "log-mode"
	logLevel = "log-level"
)

func getOutput(logMode int) io.Writer {
	switch logMode {
	case DiscardLog:
		return ioutil.Discard
	case ConsoleLog:
		return zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	case JSONLog:
		return os.Stdout
	default:
		panic(fmt.Sprintf("%d is not a known log-mode", logMode))
	}
}

// LogPrinter is used to satisfy the Printer interface on FX
type LogPrinter struct {
	Logger zerolog.Logger
}

func NewLogPrinter(logger zerolog.Logger) *LogPrinter {
	return &LogPrinter{Logger: logger}
}

func (lp *LogPrinter) Printf(format string, args ...interface{}) {
	lp.Logger.Info().Msgf(format, args...)
}

func NewLogger(cmd *cobra.Command, v *viper.Viper, env string) (zerolog.Logger, error) {
	var lm int
	var slvl string
	var err error

	if v.IsSet(logMode) {
		lm = v.GetInt(logMode)
	} else {
		lm, err = cmd.Flags().GetInt(logMode)
		if err != nil {
			return zerolog.Nop(), fmt.Errorf("error getting the value of flag %q :%s", logMode, err)
		}
	}

	if v.IsSet(logLevel) {
		slvl = v.GetString(logLevel)
	} else {
		slvl, err = cmd.Flags().GetString(logLevel)
		if err != nil {
			return zerolog.Nop(), fmt.Errorf("error getting the value of flag %q :%s", logLevel, err)
		}
	}

	lvl, err := zerolog.ParseLevel(slvl)
	if err != nil {
		return zerolog.Nop(), fmt.Errorf("error parsing the level %q: %s", slvl, err)
	}

	return zerolog.New(getOutput(lm)).With().Timestamp().Logger().With().Str("env", env).Logger().Level(lvl), nil
}
