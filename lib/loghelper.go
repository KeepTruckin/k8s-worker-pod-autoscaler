package lib

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	logMode  = "log-mode"
	logLevel = "log-level"
)

func NewLogger(cmd *cobra.Command, env string) (zerolog.Logger, error) {
	var lm int
	var slvl string
	var err error

	if viper.IsSet(logMode) {
		lm = viper.GetInt(logMode)
	} else {
		lm, err = cmd.Flags().GetInt(logMode)
		if err != nil {
			return zerolog.Nop(), fmt.Errorf("error getting the value of flag %q :%s", logMode, err)
		}
	}

	if viper.IsSet(logLevel) {
		slvl = viper.GetString(logLevel)
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

	return NewLogger(lm).With().Str("env", env).Logger().Level(lvl), nil
}
