package funcs

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadHtml(inputURL string) error {
	response, err := http.Get(inputURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Check if the response is successful
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Error2: HTTP status code %d", response.StatusCode)
	}

	// Create the save directory if it doesn't exist
	if *SaveDir != "" {
		err := os.MkdirAll(*SaveDir, 0755)
		if err != nil {
			return err
		}
	}

	// Determine the file extension
	contentType := response.Header.Get("Content-Type")
	fileExtension := getFileExtension(contentType, inputURL)
	// SaveAs := ""
	// if fileExtension == ".htm" {
	// 	SaveAs = "index"
	// } else {
	// 	SaveAs = "" // Reset SaveAs if there is a file extension
	// }

	filename := filepath.Base(inputURL)
	savePath := getUniqueFileName(filename, *SaveDir, "", fileExtension)

	file, err := os.Create(savePath)
	if err != nil && !*Mirror {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	rateLimit, err := parseRateLimit()
	if err != nil {
		return err
	}

	// Use a custom reader with rate limiting if rateLimit is greater than 0
	reader := &ProgressReader{Reader: response.Body, Progress: nil}
	var finalWriter io.Writer
	if rateLimit > 0 {
		log.Printf("Rate limit: %d bytes/s", rateLimit)
		finalWriter = NewRateLimitedWriter(file, rateLimit)
	} else {
		finalWriter = file
	}

	_, err = io.Copy(finalWriter, reader)
	if err != nil {
		return err
	}

	return nil
}
