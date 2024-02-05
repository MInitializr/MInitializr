package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(url string, filePath string) error {
	log.Printf("Downloading from url %s to file %s ...", url, filePath)
	// Create directories if they don't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// Make the HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

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
	log.Printf("Download is done successfully...")
	return nil
}
