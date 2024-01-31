package initializers


import (
	"fmt"
	"log"
	"net/url"
	"os"
	"example.com/minitializr/utils"
	"github.com/evilsocket/islazy/zip"
)

type QuarkusInitializer BaseIntializer

func (quarkusInitializer QuarkusInitializer) Initialize() {
	log.Printf("Initializing service %s with Micronaut Initializr...", quarkusInitializer.ServiceName)
	log.Printf("Initialization config %v", quarkusInitializer.Service.Config)
	baseDir := quarkusInitializer.ServiceName
	fullURL, err := quarkusInitializer.constructUrl()
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
