package funcs

import (
	"io"

	"github.com/schollz/progressbar/v3"
)

// ProgressReader is a custom reader that updates progress
type ProgressReader struct {
	io.Reader
	Progress *progressbar.ProgressBar
}

func (pr *ProgressReader) Read(p []byte) (n int, err error) {
	n, err = pr.Reader.Read(p)
	pr.Progress.Add64(int64(n))
	return
}

// // Progress is a simple progress bar implementation
// type Progress struct {
// 	Total   int64
// 	Current int64
// 	StopCh  chan struct{}
// }

// func (p *Progress) Display() {
// 	ticker := time.NewTicker(time.Millisecond * 100)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			fmt.Printf("\r%6.2f KiB / %6.2f KiB [%-50s] %.2f%% %.2f MiB/s",
// 				float64(p.Current)/1024, float64(p.Total)/1024, getProgressBar(p), float64(p.Current)/(1024*1024), float64(p.Current)/(1024*1024)/(time.Since(startTime).Seconds()))
// 		case <-p.StopCh:
// 			return
// 		}
// 	}
// }

// func (p *Progress) Write(b []byte) (n int, err error) {
// 	n = len(b)
// 	p.Current += int64(n)
// 	return n, nil
// }

// func (p *Progress) Stop() {
// 	close(p.StopCh)
// }

// func getProgressBar(p *Progress) string {
// 	width := 50
// 	progress := int((float64(p.Current) / float64(p.Total)) * float64(width))
// 	bar := make([]rune, width)
// 	for i := 0; i < width; i++ {
// 		if i < progress {
// 			bar[i] = '='
// 		} else {
// 			bar[i] = ' '
// 		}
// 	}
// 	return string(bar)
// }

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// 	"sync"
// 	"time"

// 	"github.com/schollz/progressbar/v3"
// )

// // DownloadFile initiates the download and handles progress and logging
// func DownloadFile(url string, file *os.File, savePath string, silentMode bool, wg *sync.WaitGroup) {
// 	defer wg.Done() // Decrement the wait group counter when the download is complete

// 	startTime := time.Now()

// 	// Make the HTTP request
// 	response, err := http.Get(url)
// 	if err != nil {
// 		log.Printf("Error: %v\n", err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	// Create the save directory if it doesn't exist
// 	err = os.MkdirAll(filepath.Dir(savePath), 0755)
// 	if err != nil {
// 		log.Printf("Error creating directory: %v\n", err)
// 		return
// 	}

// 	// Open or create the file for writing
// 	file, err := os.Create(savePath)
// 	if err != nil {
// 		log.Printf("Error creating file: %v\n", err)
// 		return
// 	}
// 	defer file.Close()

// 	if silentMode {
// 		// Open or create a log file in append mode
// 		logFile, err := os.OpenFile("Wget-light-log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 		if err != nil {
// 			log.Printf("Error opening log file: %v\n", err)
// 			return
// 		}
// 		defer logFile.Close()

// 		// Redirect standard output to the log file
// 		log.SetOutput(logFile)
// 	}

// 	// Output start information
// 	outputStartInfo(startTime, response, savePath, url, silentMode)

// 	// Download progress
// 	if !silentMode {
// 		contentSize := response.ContentLength
// 		progress := progressbar.DefaultBytes(
// 			contentSize,
// 			"[Downloading]...",
// 		)
// 		reader := &ProgressReader{Reader: response.Body, Progress: progress}
// 		_, err := io.Copy(io.MultiWriter(file, progress), reader)
// 		if err != nil {
// 			log.Printf("Error copying content to file: %v\n", err)
// 			return
// 		}
// 		progress.Finish()
// 	}

// 	// Output completion information
// 	outputCompletionInfo(url, startTime, silentMode)
// }

// // OutputStartInfo logs or prints the start information
// func outputStartInfo(startTime time.Time, response *http.Response, savePath, url string, silentMode bool) {
// 	if silentMode {
// 		log.Printf("start at %s", startTime.Format("2006-01-02 15:04:05"))
// 		log.Printf("Sending request, awaiting response... %s", responseStatus(response))
// 		log.Printf("Content size: %d [~%.2fMB]", response.ContentLength, float64(response.ContentLength)/(1024*1024))
// 		log.Printf("Saving file to: %s", savePath)
// 	} else {
// 		fmt.Printf("start at %s\n", startTime.Format("2006-01-02 15:04:05"))
// 		fmt.Printf("sending request, awaiting response... %s\n", responseStatus(response))
// 		fmt.Printf("content size: %d [~%.2fMB]\n", response.ContentLength, float64(response.ContentLength)/(1024*1024))
// 		fmt.Printf("saving file to: %s\n", savePath)
// 	}
// }

// // OutputCompletionInfo logs or prints the completion information
// func outputCompletionInfo(url string, startTime time.Time, silentMode bool) {
// 	if silentMode {
// 		log.Printf("Downloaded [%s]", url)
// 		log.Print("Finished at ", time.Now().Format("2006-01-02 15:04:05"), "\n\n")
// 	} else {
// 		fmt.Printf("\nDownloaded [%s]\n", url)
// 		fmt.Printf("finished at %s\n", time.Now().Format("2006-01-02 15:04:05"))
// 	}
// }

// // ProgressReader is a custom reader that updates progress
// type ProgressReader struct {
// 	io.Reader
// 	Progress *progressbar.ProgressBar
// }

// // Read reads data from the reader and updates the progress bar
// func (pr *ProgressReader) Read(p []byte) (n int, err error) {
// 	n, err = pr.Reader.Read(p)
// 	pr.Progress.Add(n)
// 	return
// }
