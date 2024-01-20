package funcs

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/schollz/progressbar/v3"
)

// ProgressReader is a custom reader that updates progress
type ProgressReader struct {
	io.Reader
	Progress *progressbar.ProgressBar
	Limiter  <-chan time.Time
}

func (pr *ProgressReader) Read(p []byte) (n int, err error) {
	// If Limiter is present, wait for the rate limit
	if pr.Limiter != nil {
		<-pr.Limiter
	}

	// Check if progress bar is not nil before updating progress
	if pr.Progress != nil {
		n, err = pr.Reader.Read(p)
		pr.Progress.Add64(int64(n))
	} else {
		// If progress bar is nil, read without updating progress
		n, err = pr.Reader.Read(p)
	}
	return
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

// RateLimitedWriter wraps an existing writer and limits the write speed
type RateLimitedWriter struct {
	writer  io.Writer
	limiter <-chan time.Time
	rate    int64
	bytes   int64
}

// NewRateLimitedWriter creates a new RateLimitedWriter with the specified writer and rate limit
func NewRateLimitedWriter(writer io.Writer, rate int64) *RateLimitedWriter {
	duration := time.Second / time.Duration(rate)
	return &RateLimitedWriter{
		writer:  writer,
		limiter: time.Tick(duration),
		rate:    rate,
		bytes:   0,
	}
}

// Write writes data to the writer with rate limiting
func (w *RateLimitedWriter) Write(p []byte) (n int, err error) {
	<-w.limiter
	n, err = w.writer.Write(p)
	w.bytes += int64(n)
	return
}
