package initializers

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"example.com/minitializr/utils"
	"github.com/evilsocket/islazy/zip"
)

type MicronautInitializer BaseIntializer

func (micronautInitializer MicronautInitializer) Initialize() {
	log.Printf("Initializing service %s with Micronaut Initializr...", micronautInitializer.ServiceName)
	log.Printf("Initialization config %v", micronautInitializer.Service.Config)
	baseDir := micronautInitializer.ServiceName
	fullURL, err := micronautInitializer.constructUrl()
	if err != nil {
		log.Println("Error:", err)
		return
	}
	log.Println(fullURL)
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("Error:", err)
		return
	}
	filePath := fmt.Sprintf("%s/.minitializer/%s.zip", userHomeDir, baseDir)
	err = utils.DownloadFile(fullURL, filePath)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	_, err = zip.Unzip(filePath, fmt.Sprintf("%s/.minitializer", userHomeDir))
	if err != nil {
		log.Println("Error:", err)
		return
	}
	err = os.RemoveAll(filePath)
	if err != nil {
		log.Println("Error:", err)
		return
	}
}

func (micronautInitializer MicronautInitializer) constructUrl() (string, error) {
	config := micronautInitializer.Service.Config
	versionAlias := "launch"
	switch config["version"] {
	case "latest", "4.2.4":
		versionAlias = "launch"
	case "snapshot", "4.2.5-SNAPSHOT":
		versionAlias = "snapshot"
	case "prev", "3.10.1":
		versionAlias = "prev"
	}
	baseURL := fmt.Sprintf("https://%s.micronaut.io/create/%s/%s.%s", versionAlias, config["type"], config["basePackage"], config["name"])
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
