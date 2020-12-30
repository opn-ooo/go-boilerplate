package logger

import (
	"github.com/blendle/zapdriver"
	"github.com/omiselabs/gin-boilerplate/config"
	"go.uber.org/zap"
)

// NewLogger ... generate new Zap logger with the passed config
func NewLogger(config *config.Config) (*zap.Logger, error) {
	logLevel := zap.NewAtomicLevel()

	err := logLevel.UnmarshalText([]byte(config.Log.Level))
	if err != nil {
		return nil, err
	}

	zapConfig := zapdriver.NewProductionConfig()
	zapConfig.Level = logLevel
	zapConfig.Encoding = config.Log.Format

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	return logger, err

}
