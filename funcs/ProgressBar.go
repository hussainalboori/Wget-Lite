package funcs

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/schollz/progressbar/v3"
	"golang.org/x/time/rate"
)

// ProgressReader is a custom reader that updates progress
type ProgressReader struct {
	io.Reader
	Progress *progressbar.ProgressBar
	Limiter  *rate.Limiter
}

func (pr *ProgressReader) Read(p []byte) (n int, err error) {
	// If Limiter is present, wait for the rate limit
	if pr.Limiter != nil {
		ctx := context.TODO()
		_ = pr.Limiter.WaitN(ctx, len(p))
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
	limiter *rate.Limiter
	bytes   int64
}

// NewRateLimitedWriter creates a new RateLimitedWriter with the specified writer and rate limit
func NewRateLimitedWriter(writer io.Writer, rateLimit int64) *RateLimitedWriter {
	return &RateLimitedWriter{
		writer:  writer,
		limiter: rate.NewLimiter(rate.Limit(rateLimit), int(rateLimit)),
		bytes:   0,
	}
}

// Write writes data to the writer with rate limiting
func (w *RateLimitedWriter) Write(p []byte) (n int, err error) {
	err = w.limiter.WaitN(context.TODO(), len(p))
	if err != nil {
		return 0, err
	}
	n, err = w.writer.Write(p)
	w.bytes += int64(n)
	return
}
