package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Debug(key string, message string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Debug().Str(key, message).Send()
}

func Info(key string, message string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Str(key, message).Send()
}

func Log(message string) {

}

func Warn(message string) {

}

func Error(message string) {

}
