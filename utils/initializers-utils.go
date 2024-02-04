package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/evilsocket/islazy/zip"
)

func InitializeWithWebIntializer(serviceName, baseDir, initializerUrl string) error {
	log.Println(initializerUrl)
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	filePath := fmt.Sprintf("%s/.minitializer/%s.zip", userHomeDir, serviceName)
	err = DownloadFile(initializerUrl, filePath)
	if err != nil {
		return err
	}
	_, err = zip.Unzip(filePath, fmt.Sprintf("%s/.minitializer/%s", userHomeDir, baseDir))
	if err != nil {
		return err
	}
	err = os.RemoveAll(filePath)
	if err != nil {		
		return err
	}
	return nil
}