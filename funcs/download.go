package funcs

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/schollz/progressbar/v3"
)

var startTime time.Time

func DownloadFile(url string, saveAs string, saveDir string, rateLimit int64) error {
	startTime := time.Now()
	fmt.Printf("start at %s\n", startTime.Format("2006-01-02 15:04:05"))

	// Make the HTTP request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error: HTTP status code %d\n", response.StatusCode)
		os.Exit(1)
	}

	// Create the save directory if it doesn't exist
	if saveDir != "" {
		err := os.MkdirAll(saveDir, 0755)
		if err != nil {
			return err
		}
	}

	filename := filepath.Base(url)

	savePath := getUniqueFileName(filename, saveDir, saveAs)
	file, err := os.Create(savePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()
	//////////////////////////////////////////////////////////////////////////////
	contentSize := response.ContentLength
	fmt.Printf("sending request, awaiting response... status %d OK\n", response.StatusCode)
	fmt.Printf("content size: %d [~%.2fMB]\n", contentSize, float64(contentSize)/(1024*1024))
	fmt.Printf("saving file to: %s\n", savePath)

	//////////////////////////////////////////////////////////////////////////////

	// Create a progress progress
	progress := progressbar.DefaultBytes(
		contentSize,
		"[downloading]",
	)

	// Use a custom reader to update the progress
	reader := &ProgressReader{Reader: response.Body, Progress: progress}
	_, err = io.Copy(io.MultiWriter(file, progress), reader)
	if err != nil {
		return err
	}

	progress.Finish()
	fmt.Printf("\nDownloaded [%s]\n", url)
	fmt.Printf("finished at %s\n", time.Now().Format("2006-01-02 15:04:05"))

	return nil
}

// // rateLimitedWriter wraps an existing writer and limits the write speed
// type rateLimitedWriter struct {
// 	writer     io.Writer
// 	limiter    <-chan time.Time
// 	rate       int64
// 	bytes      int64
// 	lastUpdate time.Time
// }

// // NewRateLimitedWriter creates a new rateLimitedWriter with the specified writer and rate limit
// func NewRateLimitedWriter(writer io.Writer, rate int64) *rateLimitedWriter {
// 	duration := time.Second / time.Duration(rate)
// 	return &rateLimitedWriter{
// 		writer:     writer,
// 		limiter:    time.Tick(duration),
// 		rate:       rate,
// 		bytes:      0,
// 		lastUpdate: time.Now(),
// 	}
// }

// // Write writes data to the writer with rate limiting and throughput monitoring
// func (w *rateLimitedWriter) Write(p []byte) (n int, err error) {
// 	<-w.limiter
// 	n, err = w.writer.Write(p)
// 	if n > 0 {
// 		w.bytes += int64(n)
// 		currTime := time.Now()
// 		elapsed := currTime.Sub(w.lastUpdate).Seconds()
// 		if elapsed >= 1.0 {
// 			throughput := float64(w.bytes) / elapsed
// 			w.lastUpdate = currTime
// 			w.bytes = 0
// 			fmt.Printf("Throughput: %.2f bytes/s\n", throughput)
// 		}
// 	}
// 	return
// }
