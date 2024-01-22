package funcs

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
)

func DownloadFile(inputURL string, wg *sync.WaitGroup) error {
	defer wg.Done()

	// Make the HTTP request
	req, err := http.NewRequest("GET", inputURL, nil)
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
	fileExtension := getFileExtension(contentType, inputURL)

	var filename string
	var savePath string

	if *Mirror {
		// Create the relative path for saving the file
		relativePath := getRelativePath(inputURL)
		savePath = filepath.Join(*SaveDir, relativePath)

		// Create directories if they don't exist
		if err := os.MkdirAll(filepath.Dir(savePath), 0755); err != nil {
			return fmt.Errorf("error creating directories: %v", err)
		}

		// Check if the URL has an extension, indicating it's a file
		if filepath.Ext(inputURL) == "" {
			// fmt.Printf("Skipping directory: %s\n", inputURL)
			// return nil
			DownloadHtml(inputURL)
		}

	} else {
		filename = filepath.Base(inputURL)
		// filename = getUniqueFileName(filename, *SaveDir, *SaveAs, fileExtension)
		savePath = getUniqueFileName(filename, *SaveDir, *SaveAs, fileExtension)
	}

	file, err := os.Create(savePath)
	if err != nil && !*Mirror {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	Output(response, file, savePath, inputURL)

	return nil
}

func getRelativePath(inputURL string) string {
	// Parse the base URL
	base, err := url.Parse(BaseURL)
	if err != nil {
		log.Printf("Error parsing baseURL: %v", err)
		return ""
	}

	// Parse the resource URL
	resourceURL, err := url.Parse(inputURL)
	if err != nil {
		log.Printf("Error parsing resource URL: %v", err)
		return ""
	}

	// Compute the relative path
	relativePath, err := filepath.Rel(base.Path, resourceURL.Path)
	if err != nil {
		log.Printf("Error computing relative path: %v", err)
		return ""
	}

	return relativePath
}
