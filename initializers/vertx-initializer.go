package initializers

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	"example.com/minitializr/utils"
)

type VertxInitializer BaseInitializer

func (initializer VertxInitializer) Initialize() {
	log.Printf("Initializing service %s with SpringBoot Initializr...", initializer.ServiceName)
	log.Printf("Initialization config %v", initializer.Service.Config)
	fullURL, err := initializer.constructUrl()
	if err != nil {
		log.Println("Error:", err)
		return
	}
	err = utils.InitializeWithWebIntializer(initializer.ServiceName, initializer.ServiceName, fullURL)
	if err != nil {
		log.Println("Error:", err)
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
