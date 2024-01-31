package initializers

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"github.com/evilsocket/islazy/zip"
	"example.com/minitializr/utils"
)

type SpringBootIntializer BaseIntializer;

func (springBootIntializer SpringBootIntializer) Initialize(){
	log.Printf("Initializing service %s with SpringBoot Initializr...", springBootIntializer.ServiceName)
	log.Printf("Initialization config %v", springBootIntializer.Service.Config)
	baseURL := "https://start.spring.io/starter.zip"
	baseDir := springBootIntializer.ServiceName
	fullURL, err:= springBootIntializer.constructUrl(baseURL, springBootIntializer.Service.Config)
	if(err != nil){
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
	_,err = zip.Unzip(filePath, fmt.Sprintf("%s/.minitializer/%s", userHomeDir, baseDir))
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

func (SpringBootIntializer) constructUrl(baseUrl string, params map[string]any) (string, error) {
	urlParams := url.Values{}
	for k,v := range params {
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
	fullURL := fmt.Sprintf("%s?%s", baseUrl, urlParams.Encode())
	return fullURL, nil
}