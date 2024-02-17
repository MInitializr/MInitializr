package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/HamzaBenyazid/minitializr/logger"
)

func DownloadFile(response *http.Response, filePath string) error {
	logger.Debug("Downloading file from response ...")
	defer response.Body.Close()

	// Create directories if they don't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the response body to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	logger.Debug("Download is done successfully...")
	return nil
}
