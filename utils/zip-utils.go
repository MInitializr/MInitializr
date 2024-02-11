package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(src string, dest string) ([]string, error) {
	var outFile *os.File
	var zipFile io.ReadCloser
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	clean := func() {
		if outFile != nil {
			outFile.Close()
			outFile = nil
		}

		if zipFile != nil {
			zipFile.Close()
			zipFile = nil
		}
	}

	for _, f := range r.File {
		zipFile, err = f.Open()
		if err != nil {
			return filenames, err
		}

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: https://snyk.io/research/zip-slip-vulnerability#go
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			clean()
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			clean()
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			clean()
			return filenames, err
		}

		outFile, err = os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			clean()
			return filenames, err
		}

		_, err = io.Copy(outFile, zipFile)
		clean()
		if err != nil {
			return filenames, err
		}
	}

	return filenames, nil
}


func ZipDirectory(destFile, srcDir string ) error {
    file, err := os.Create(destFile)
    if err != nil {
        return err
    }
    defer file.Close()

    w := zip.NewWriter(file)
    defer w.Close()

    walker := func(path string, info os.FileInfo, err error) error {
        fmt.Printf("Crawling: %#v\n", path)
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }
        file, err := os.Open(path)
        if err != nil {
            return err
        }
        defer file.Close()

        // Ensure that `path` is not absolute; it should not start with "/".
        // This snippet happens to work because I don't use 
        // absolute paths, but ensure your real-world code 
        // transforms path into a zip-root relative path.
        f, err := w.Create(path)
        if err != nil {
            return err
        }

        _, err = io.Copy(f, file)
        if err != nil {
            return err
        }

        return nil
    }
	err = os.Chdir(srcDir)
	if err != nil {
		return err
	}
    err = filepath.Walk(".", walker)
    if err != nil {
        return err
    }
	return nil
}

