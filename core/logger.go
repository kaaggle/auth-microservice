package core

import "go.uber.org/zap"

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"logs.log",
	}
	return cfg.Build()
}
