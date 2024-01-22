package funcs

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// MirrorOptions represents options for mirroring a website
type MirrorOptions struct {
	RejectFileTypes []string // File types to reject (e.g., ".jpg", ".png")
	ExcludeDirs     []string // Directories to exclude from mirroring
}

var BaseURL string

// Mirror mirrors a website's frontend by parsing HTML and downloading linked resources
func Mirroring(inputURL string, options MirrorOptions) error {
	BaseURL = inputURL

	// Check if the URL is empty
	if inputURL == "" {
		return fmt.Errorf("empty URL provided")
	}
	fmt.Println(inputURL)

	// Check if the URL has a valid scheme
	parsedURL, err := url.Parse(inputURL)
	if err != nil || parsedURL.Scheme == "" {
		return fmt.Errorf("invalid URL: %s", inputURL)
	}

	// Fetch HTML content from the specified URL
	htmlContent, err := FetchHTML(inputURL)
	if err != nil {
		return err
	}

	// Extract resources from HTML content
	resources, err := ExtractResources(htmlContent, BaseURL, options)
	if err != nil && len(resources) == 0 {
		return err
	}

	// Debugging prints
	// fmt.Println("------------------------------")
	// fmt.Println("Resources len||", len(resources))
	// fmt.Println(resources)

	// Use a WaitGroup to wait for all downloads to finish
	var Mwg sync.WaitGroup
	// Download resources
	for _, resourceURL := range resources {
		// Download the resource only if it's not rejected and not in an excluded directory
		if !rejectResource(resourceURL, options.RejectFileTypes) && !isExcludedDir(resourceURL, options.ExcludeDirs) {
			// Increment the WaitGroup counter
			// fmt.Println("URL:sssssssss")

			Mwg.Add(1)
			go func(resourceURL string) {
				// Decrement the counter when the goroutine completes
				err := DownloadFile(resourceURL, &Mwg)
				if err != nil {
					log.Printf("Error downloading %s: %v", resourceURL, err)
				}
			}(resourceURL)
		}
	}

	// download the html index
	err = DownloadHtml(inputURL)
	if err != nil {
		fmt.Printf("Error downloading html index: %v\n", err)
		return err
	}

	// Wait for all downloads to finish
	Mwg.Wait()

	return nil
}

// isExcludedDir checks if the URL is in an excluded directory
func isExcludedDir(url string, ExcludeDirs []string) bool {
	// Join the array elements into a single string
	joinedFileTypes := strings.Join(ExcludeDirs, "")

	// Check if the joined string is empty
	if joinedFileTypes == "" {
		return false
	}

	for _, excludedDir := range ExcludeDirs {
		if strings.Contains(url, excludedDir) {
			return true
		}
	}

	return false
}

// rejectResource checks if the resource should be rejected based on file types
func rejectResource(inputURL string, rejectFileTypes []string) bool {
	// Join the array elements into a single string
	joinedFileTypes := strings.Join(rejectFileTypes, "")

	// Check if the joined string is empty
	if joinedFileTypes == "" {
		return false
	}

	for _, fileType := range rejectFileTypes {
		if strings.HasSuffix(inputURL, fileType) {
			return true
		}
	}

	return false
}

// this file contains all there is to it about mirroring
// we tokenize html, and recursively follow any links in it

func FetchHTML(inputURL string) (string, error) {
	response, err := http.Get(inputURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	htmlContent, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(htmlContent), nil
}

func ExtractResources(htmlContent, baseURL string, options MirrorOptions) ([]string, error) {
	var resources []string

	// Parse the base URL
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	// tokenize html
	tokenizer := html.NewTokenizer(strings.NewReader(htmlContent))
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return resources, nil
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "link" || token.Data == "script" || token.Data == "img" {
				for _, attr := range token.Attr {
					if attr.Key == "href" || attr.Key == "src" {
						resourceURL := attr.Val
						// Parse the resource URL
						resourceURLParsed, err := url.Parse(resourceURL)
						if err != nil {
							return nil, err
						}

						// Check if the host matches the base URL's host

						// Resolve the URL
						absoluteURL := base.ResolveReference(resourceURLParsed).String()
						if !rejectResource(resourceURL, options.RejectFileTypes) && !isExcludedDir(resourceURL, options.ExcludeDirs) {

							// Check if the link is an HTML page and append ".html"
							if strings.HasSuffix(resourceURLParsed.Path, ".html") {
								absoluteURL += ".html"
							}

							resources = append(resources, absoluteURL)
						}
					}
				}
			} else if token.Data == "a" { // Check for anchor tags
				for _, attr := range token.Attr {

					if attr.Key == "href" {

						resourceURL := attr.Val
						// Parse the resource URL
						resourceURLParsed, err := url.Parse(resourceURL)
						if err != nil {
							return nil, err
						}

						// Resolve the URL
						absoluteURL := base.ResolveReference(resourceURLParsed).String()

						// Check if the link is an HTML page and append ".html"
						if strings.HasSuffix(resourceURLParsed.Path, ".html") {
							absoluteURL += ".html"
						}
					}
				}
			}
		}

	}

}
