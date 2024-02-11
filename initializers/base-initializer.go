package initializers

import (
	"log"

	"example.com/minitializr/types"
)

type BaseInitializer struct {
	ServiceName string
	Service     types.MIService
}

func (BaseInitializer) Initialize(*types.MIConfig) {
	log.Println("BaseIntializer")
}
