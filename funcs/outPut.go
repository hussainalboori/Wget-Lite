package funcs

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
	"golang.org/x/time/rate"
)

func Output(response *http.Response, file *os.File, savePath, inputURL string) error {
	startTime := time.Now()

	contentSize := response.ContentLength

	rateLimit, err := parseRateLimit()
	if err != nil {
		return err
	}
	var limiter *rate.Limiter
	if rateLimit > 0 {
		limiter = rate.NewLimiter(rate.Limit(rateLimit), int(rateLimit))
	}

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
		log.Printf("Saving file to: %s\n\n", savePath)

		// Use a custom reader with rate limiting if rateLimit is greater than 0
		// reader := &ProgressReader{Reader: response.Body, Progress: nil}

		// Use a custom reader with rate limiting if rateLimit is greater than 0
		reader := &ProgressReader{
			Reader:   response.Body,
			Progress: nil,
			Limiter:  limiter,
		}

		var finalWriter io.Writer
		if rateLimit > 0 {
			log.Printf("Rate limit: %d bytes/s", rateLimit)
			finalWriter = NewRateLimitedWriter(file, rateLimit)
		} else { // reader := &ProgressReader{Reader: response.Body, Progress: progress}

			finalWriter = file
		}

		_, err = io.Copy(finalWriter, reader)
		if err != nil {
			return err
		}

		if *InputFile == "" {
			// Log the download completion
			log.Printf("Downloaded [%s]", inputURL)
			log.Print("Finished at ", time.Now().Format("2006-01-02 15:04:05"),
				"\n<-------------------------------------------------------------------------------------------->\n")
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

		// Use a custom reader with rate limiting if rateLimit is greater than 0
		// reader := &ProgressReader{Reader: response.Body, Progress: progress}

		// Use a custom reader with rate limiting if rateLimit is greater than 0

		reader := &ProgressReader{
			Reader:   response.Body,
			Progress: progress,
			Limiter:  limiter,
		}

		// Determine the final writer based on rate limiting
		var finalWriter io.Writer
		if rateLimit > 0 {
			// Create a RateLimitedWriter to limit the download speed
			finalWriter = NewRateLimitedWriter(io.MultiWriter(file, progress), rateLimit)
		} else {
			// If rate limiting is not required, use the standard io.MultiWriter
			finalWriter = io.MultiWriter(file, progress)
		}

		// Copy data from the reader to the final writer, applying rate limiting if needed
		_, err := io.Copy(finalWriter, reader)
		if err != nil {
			return err
		}

		// Finish the progress bar once the download is complete
		progress.Finish()

		// If no input file is specified, log the download completion information
		if *InputFile == "" && !*Mirror {
			// Log the download completion
			fmt.Printf("\nDownloaded [%s]\n", inputURL)
			fmt.Printf("finished at %s\n\n", time.Now().Format("2006-01-02 15:04:05"))
		}

	}

	return nil
}
