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
	// startTime := time.Now()
	// 	fmt.Printf("start at %s\n", startTime.Format("2006-01-02 15:04:05"))

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
	if *SaveDir != "" {
		err := os.MkdirAll(*SaveDir, 0755)
		if err != nil {
			return err
		}
	}

	filename := filepath.Base(url)

	savePath := getUniqueFileName(filename, *SaveDir, *SaveAs)
	file, err := os.Create(savePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	Output(response, file, savePath, url)
	return nil
}

// // RateLimitedWriter wraps an existing writer and limits the write speed
// type RateLimitedWriter struct {
// 	writer     io.Writer
// 	limiter    <-chan time.Time
// 	rate       int64
// 	bytes      int64
// 	lastUpdate time.Time
// }

// // NewRateLimitedWriter creates a new RateLimitedWriter with the specified writer and rate limit
// func NewRateLimitedWriter(writer io.Writer, rate int64) *RateLimitedWriter {
// 	duration := time.Second / time.Duration(rate)
// 	return &RateLimitedWriter{
// 		writer:     writer,
// 		limiter:    time.Tick(duration),
// 		rate:       rate,
// 		bytes:      0,
// 		lastUpdate: time.Now(),
// 	}
// }

// // Write writes data to the writer with rate limiting and throughput monitoring
// func (w *RateLimitedWriter) Write(p []byte) (n int, err error) {
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
