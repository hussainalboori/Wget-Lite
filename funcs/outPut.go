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

		// Open or create a log file
		logFile, err := os.Create("Wget-light-log.txt")
		if err != nil {
			return err
		}
		defer logFile.Close()

		// Redirect standard output to the log file
		log.SetOutput(logFile)

		// Log the output
		log.Printf("start at %s", startTime.Format("2006-01-02 15:04:05"))
		log.Printf("Sending request, awaiting response... status %d OK", response.StatusCode)
		log.Printf("Content size: %d [~%.2fMB]", response.ContentLength, float64(response.ContentLength)/(1024*1024))
		log.Printf("Saving file to: %s", savePath)

		// Use a custom reader to update the progress
		reader := &ProgressReader{Reader: response.Body, Progress: nil}
		_, err = io.Copy(file, reader)
		if err != nil {
			return err
		}

		// Log the download completion
		log.Printf("Downloaded [%s]", url)
		log.Printf("Finished at %s", time.Now().Format("2006-01-02 15:04:05"))
	} else {

		//////////////////////////////////////////////////////////////////////////////
		fmt.Printf("start at %s\n", startTime.Format("2006-01-02 15:04:05"))
		fmt.Printf("sending request, awaiting response... status %d OK\n", response.StatusCode)
		fmt.Printf("content size: %d [~%.2fMB]\n", contentSize, float64(contentSize)/(1024*1024))
		fmt.Printf("saving file to: %s\n", savePath)
		//////////////////////////////////////////////////////////////////////////////
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
		fmt.Printf("\nDownloaded [%s]\n", url)
		fmt.Printf("finished at %s\n", time.Now().Format("2006-01-02 15:04:05"))
	}

	return nil
}
