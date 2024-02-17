package initializers

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/HamzaBenyazid/minitializr/logger"
	"github.com/HamzaBenyazid/minitializr/types"
	"github.com/HamzaBenyazid/minitializr/utils"
	"go.uber.org/zap"
)

type VertxInitializer BaseInitializer

func (initializer VertxInitializer) Initialize(miConfig *types.MIConfig) {
	logger.Debugf("Initializing service %s with SpringBoot Initializr...", initializer.ServiceName)
	logger.Debugf("Initialization config %v", initializer.Service.Config)
	fullURL, err := initializer.constructUrl()
	if err != nil {
		logger.Error("Error:", zap.Error(err))
		return
	}
	err = utils.InitializeWithWebIntializer(miConfig.Metadata["name"], initializer.ServiceName, initializer.ServiceName, fullURL)
	if err != nil {
		logger.Error("Error:", zap.Error(err))
		return
	}
}

func (initializer VertxInitializer) constructUrl() (string, error) {
	config := initializer.Service.Config
	baseURL := "https://start.vertx.io/starter.zip"
	urlParams := url.Values{}
	for k, v := range config {
		switch val := v.(type) {
		case string:
			urlParams.Add(k, val)
		case int:
			urlParams.Add(k, strconv.Itoa(val))
		default:
			// no match; here v has the same type as i
		}
	}
	// Construct the full URL with parameters
	fullURL := fmt.Sprintf("%s?%s", baseURL, urlParams.Encode())
	return fullURL, nil
}
