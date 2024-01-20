package funcs

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
)

func Output(response *http.Response, file *os.File, savePath, url string) error {
	startTime := time.Now()

	contentSize := response.ContentLength

	if *SilentMode {
		fmt.Println("Logs will be written to Wget-light-log.txt")

		// Open or create a log file in append mode
		logFile, err := os.OpenFile("Wget-light-log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer logFile.Close()

		// Redirect standard output to the log file
		log.SetOutput(logFile)

		// Log the output
		log.Printf("start at %s", startTime.Format("2006-01-02 15:04:05"))
		log.Printf("Sending request, awaiting response... %s", responseStatus(response))
		log.Printf("Content size: %d [~%.2fMB]", contentSize, float64(contentSize)/(1024*1024))
		log.Printf("Saving file to: %s", savePath)

		// Use a custom reader to update the progress
		reader := &ProgressReader{Reader: response.Body, Progress: nil}
		_, err = io.Copy(file, reader)
		if err != nil {
			return err
		}
		if *InputFile == "" {
			// Log the download completion
			log.Printf("Downloaded [%s]", url)
			log.Print("Finished at ", time.Now().Format("2006-01-02 15:04:05"), "\n\n")
		}
	} else {

		fmt.Printf("start at %s\n", startTime.Format("2006-01-02 15:04:05"))
		fmt.Printf("sending request, awaiting response... %s\n", responseStatus(response))
		fmt.Printf("content size: %d [~%.2fMB]\n", contentSize, float64(contentSize)/(1024*1024))
		fmt.Printf("saving file to: %s\n", savePath)

		// Create a progress progress
		progress := progressbar.DefaultBytes(
			contentSize,
			"[Downloading]...",
		)
		// Use a custom reader to update the progress
		reader := &ProgressReader{Reader: response.Body, Progress: progress}
		_, err := io.Copy(io.MultiWriter(file, progress), reader)
		if err != nil {
			return err
		}

		progress.Finish()
		if *InputFile == "" {
			// Log the download completion
			fmt.Printf("\nDownloaded [%s]\n", url)
			fmt.Printf("finished at %s\n\n", time.Now().Format("2006-01-02 15:04:05"))
		}
	}

	return nil
}

// ResponseStatus returns the formatted status message
func responseStatus(response *http.Response) string {
	switch response.StatusCode {
	case http.StatusOK:
		return "Status 200 OK."
	case http.StatusForbidden:
		return "Access denied. Status 403 Forbidden."
	default:
		return fmt.Sprintf("Unexpected status code: %d %s", response.StatusCode, http.StatusText(response.StatusCode))
	}
}
