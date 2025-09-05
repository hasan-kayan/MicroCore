package logger

import "go.uber.org/zap"

func New(env string, level string) (*zap.Logger, error) {
	if env == "prod" {
		return zap.NewProduction()
	}
	return zap.NewDevelopment()
}
