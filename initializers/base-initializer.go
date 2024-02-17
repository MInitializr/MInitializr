package initializers

import (
	"github.com/HamzaBenyazid/minitializr/logger"
	"github.com/HamzaBenyazid/minitializr/types"
)

type BaseInitializer struct {
	ServiceName string
	Service     types.MIService
}

func (BaseInitializer) Initialize(*types.MIConfig) {
	logger.Debug("BaseIntializer")
}
