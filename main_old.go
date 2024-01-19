package main

import (
	"flag"
	"fmt"
	"os"
)

// func downloadFile(url string, saveAs string, saveDir string, rateLimit int64) error {
// 	// Create the save directory if it doesn't exist
// 	if saveDir != "" {
// 		err := os.MkdirAll(saveDir, 0755)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	// Determine the filename
// 	var filename string
// 	if saveAs != "" {
// 		filename = saveAs
// 	} else {
// 		filename = filepath.Base(url)
// 	}

// 	// Create the file to save
// 	savePath := filepath.Join(saveDir, filename)
// 	file, err := os.Create(savePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Make the HTTP request
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Get the content length
// 	contentLength := resp.ContentLength

// 	// Create the progress bar
// 	fmt.Printf("Downloading %s...\n", filename)
// 	bar := &progressBar{total: contentLength}
// 	defer bar.finish()

// 	// Create a multi-writer to write to file and update progress bar
// 	writer := io.MultiWriter(file, bar)

// 	// Wrap the writer with a rate limiter if rateLimit is specified
// 	if rateLimit > 0 {
// 		writer = NewRateLimitedWriter(writer, rateLimit)
// 	}

// 	// Copy the response body to file with progress update
// 	startTime := time.Now()
// 	_, err = io.Copy(writer, resp.Body)
// 	if err != nil {
// 		return err
// 	}
// 	elapsedTime := time.Since(startTime)

// 	fmt.Printf("Downloaded %s in %s\n", filename, elapsedTime)

// 	return nil
// }

// // progressBar represents a basic progress bar for tracking download progress
// type progressBar struct {
// 	total      int64
// 	downloaded int64
// }

// // Write updates the progress bar when data is written
// func (p *progressBar) Write(b []byte) (int, error) {
// 	n := len(b)
// 	p.downloaded += int64(n)
// 	p.printProgress()
// 	return n, nil
// }

// // finish completes the progress bar when the download is finished
// func (p *progressBar) finish() {
// 	p.downloaded = p.total
// 	p.printProgress()
// 	fmt.Println()
// }

// // printProgress prints the current progress of the download
// func (p *progressBar) printProgress() {
// 	progress := float64(p.downloaded) / float64(p.total) * 100
// 	fmt.Printf("\r%.2f%%", progress)
// }

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

func main_old() {
	if len(os.Args) < 2 {
		fmt.Println("usge go run main.go <url> -o <save as> -p <path to save> -rate-limit <rate limt>")
		return
	}
	// url := os.Args[1]
	// saveAs := flag.String("O", "", "save the downloaded file with a different name")
	// saveDir := flag.String("P", "", "directory to save the downloaded file")
	// rateLimit := flag.Int64("rate-limit", 0, "limit the download speed in bytes per second")
	flag.Parse()

	// err := downloadFile(url, *saveAs, *saveDir, *rateLimit)
	// if err != nil {
	// 	fmt.Printf("Error downloading file: %v\n", err)
	// 	return
	// }

	fmt.Println("Download completed successfully")
}
