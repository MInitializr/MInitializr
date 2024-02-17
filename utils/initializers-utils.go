package utils

import (
	"fmt"
	"net/http"
	"os"
)

func InitializeWithWebIntializer(projectName, serviceName, baseDir, initializerUrl string) error {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	projectPath := fmt.Sprintf("%s/.minitializer/%s", userHomeDir, projectName)
	servicePath := fmt.Sprintf("%s/%s", projectPath, serviceName)
	zipPath := servicePath + ".zip"
	response, err := http.Get(initializerUrl)
	if err != nil {
		return err
	}
	err = DownloadFile(response, zipPath)
	if err != nil {
		return err
	}
	
	targetPath := projectPath + "/" + baseDir

	_, err = Unzip(zipPath, targetPath)
	if err != nil {
		return err
	}
	err = os.RemoveAll(zipPath)
	if err != nil {
		return err
	}
	return nil
}
