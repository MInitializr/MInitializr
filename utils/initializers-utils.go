package utils

import (
	"fmt"
	"log"
	"os"
)

func InitializeWithWebIntializer(projectName, serviceName, baseDir, initializerUrl string) error {
	log.Println(initializerUrl)
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	servicePath := fmt.Sprintf("%s/.minitializer/%s/%s", userHomeDir, projectName, serviceName)
	zipPath := servicePath + ".zip"
	err = DownloadFile(initializerUrl, zipPath)
	if err != nil {
		return err
	}
	_, err = Unzip(zipPath, servicePath)
	if err != nil {
		return err
	}
	err = os.RemoveAll(zipPath)
	if err != nil {
		return err
	}
	return nil
}
