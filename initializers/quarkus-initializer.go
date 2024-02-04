package initializers


import (
	"fmt"
	"log"
	"net/url"
	"example.com/minitializr/utils"
)

type QuarkusInitializer BaseInitializer

func (initializer QuarkusInitializer) Initialize() {
	log.Printf("Initializing service %s with %s...", initializer, initializer.ServiceName)
	log.Printf("Initialization config %v", initializer.Service.Config)
	fullURL, err := initializer.constructUrl()
	if err != nil {
		log.Println("Error:", err)
		return
	}
	err = utils.InitializeWithWebIntializer(initializer.ServiceName, "", fullURL)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	
}

func (quarkusInitializer QuarkusInitializer) constructUrl() (string, error) {
	config := quarkusInitializer.Service.Config
	baseURL := "https://code.quarkus.io/d"
	urlParams := url.Values{}
	urlParams.Add("a", fmt.Sprintf("%v", config["artifact"]))
	urlParams.Add("g", fmt.Sprintf("%v", config["group"]))
	urlParams.Add("j", fmt.Sprintf("%v", config["javaVersion"]))
	urlParams.Add("b", fmt.Sprintf("%v", config["buildTool"]))
	urlParams.Add("v", fmt.Sprintf("%v", config["version"]))
	urlParams.Add("cn", fmt.Sprintf("%v", config["code.quarkus.io"]))
	features, ok := config["extensions"].([]any)
	if ok {
		for _, v := range features {
			urlParams.Add("e", fmt.Sprintf("%v", v))
		}
	}

	// Construct the full URL with parameters
	fullURL := fmt.Sprintf("%s?%s", baseURL, urlParams.Encode())
	return fullURL, nil
}
