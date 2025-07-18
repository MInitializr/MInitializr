package initializers

import (
	"fmt"

	"net/url"

	"github.com/HamzaBenyazid/minitializr/logger"
	"github.com/HamzaBenyazid/minitializr/types"
	"github.com/HamzaBenyazid/minitializr/utils"
	"go.uber.org/zap"
)

type MicronautInitializer BaseInitializer

func (initializer MicronautInitializer) Initialize(miConfig *types.MIConfig) {
	logger.Debugf("Initializing service %s with Micronaut Initializr...", initializer.ServiceName)
	logger.Debugf("Initialization config %v", initializer.Service.Config)
	fullURL, err := initializer.constructUrl()
	if err != nil {
		logger.Error("Error:", zap.Error(err))
		return
	}
	err = utils.InitializeWithWebIntializer(miConfig.Metadata["name"], initializer.ServiceName, "", fullURL)
	if err != nil {
		logger.Error("Error:", zap.Error(err))
		return
	}
}

func (initializer MicronautInitializer) constructUrl() (string, error) {
	config := initializer.Service.Config
	baseURL := fmt.Sprintf("https://%s.micronaut.io/create/%s/%s.%s", config["version"], config["type"], config["basePackage"], config["name"])
	urlParams := url.Values{}
	urlParams.Add("lang", fmt.Sprintf("%v", config["lang"]))
	urlParams.Add("build", fmt.Sprintf("%v", config["build"]))
	urlParams.Add("javaVersion", fmt.Sprintf("%v", config["javaVersion"]))
	urlParams.Add("test", fmt.Sprintf("%v", config["test"]))
	features, ok := config["features"].([]any)
	if ok {
		for _, v := range features {
			urlParams.Add("features", fmt.Sprintf("%v", v))
		}
	}

	// Construct the full URL with parameters
	fullURL := fmt.Sprintf("%s?%s", baseURL, urlParams.Encode())
	return fullURL, nil
}
