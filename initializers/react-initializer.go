package initializers

import (
	"github.com/HamzaBenyazid/minitializr/types"
)

type ReactInitializer struct {
	ServiceName string
	Service     types.MIService
}

func (ReactInitializer) Initialize(*types.MIConfig) {
	// start a nodejs container with ~/.minitializer mounted

	// call `npx create-react-app` to generate the project

	// zip the project and delete the generated folder

	//send the zip file as response
}
