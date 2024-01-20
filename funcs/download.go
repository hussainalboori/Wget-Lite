package funcs

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

func DownloadFile(url string, wg *sync.WaitGroup) error {
	defer wg.Done()
	// Make the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set a common User-Agent header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	// Follow redirects
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Allow redirects
			return nil
		},
	}

	response, err := client.Do(req)
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
	fileExtension := getFileExtension(contentType, url)

	filename := filepath.Base(url)

	savePath := getUniqueFileName(filename, *SaveDir, *SaveAs, fileExtension)
	file, err := os.Create(savePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	Output(response, file, savePath, url)

	return nil
}
