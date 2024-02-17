package service

import (
	"bytes"
	"fmt"
	"os"

	"github.com/HamzaBenyazid/minitializr/initializers"
	"github.com/HamzaBenyazid/minitializr/types"
	"github.com/HamzaBenyazid/minitializr/utils"
)

func Generate(miConfig *types.MIConfig) (*bytes.Buffer, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	miDataDir := fmt.Sprintf("%s/.minitializer", userHomeDir)
	projectPath := fmt.Sprintf("%s/%s", miDataDir, miConfig.Metadata["name"])

	defer cleanUp(projectPath)

	for serviceName, service := range miConfig.Services {
		getInitializer(serviceName, service).Initialize(miConfig)
	}
	
	zipFileBuf, err := utils.ZipDirectory(projectPath)
	if err != nil {
		return nil, err
	}

	return zipFileBuf, nil
}

func cleanUp(projectPath string) error {
	err := os.RemoveAll(projectPath)
	if err != nil {
		return err
	}
	return nil
}

func getInitializer(serviceName string, service types.MIService) types.Initializer {
	baseInitializer := initializers.BaseInitializer{
		ServiceName: serviceName,
		Service:     service,
	}
	switch service.Technology {
	case "spring-boot":
		return initializers.SpringBootInitializer(baseInitializer)
	case "micronaut":
		return initializers.MicronautInitializer(baseInitializer)
	case "quarkus":
		return initializers.QuarkusInitializer(baseInitializer)
	case "grails":
		return initializers.GrailsInitializer(baseInitializer)
	case "vertx":
		return initializers.VertxInitializer(baseInitializer)
	}
	return baseInitializer
}
