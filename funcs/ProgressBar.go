package funcs

import (
	"fmt"
	"io"
	"time"

	"github.com/schollz/progressbar/v3"
)

// Progress is a simple progress bar implementation
type Progress struct {
	Total   int64
	Current int64
	StopCh  chan struct{}
}

func (pr *ProgressReader) Read(p []byte) (n int, err error) {
	n, err = pr.Reader.Read(p)
	pr.Progress.Add64(int64(n))
	return
}

// ProgressReader is a custom reader that updates progress
type ProgressReader struct {
	io.Reader
	Progress *progressbar.ProgressBar
}

func (p *Progress) Display() {
	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("\r%6.2f KiB / %6.2f KiB [%-50s] %.2f%% %.2f MiB/s",
				float64(p.Current)/1024, float64(p.Total)/1024, getProgressBar(p), float64(p.Current)/(1024*1024), float64(p.Current)/(1024*1024)/(time.Since(startTime).Seconds()))
		case <-p.StopCh:
			return
		}
	}
}

func (p *Progress) Write(b []byte) (n int, err error) {
	n = len(b)
	p.Current += int64(n)
	return n, nil
}

func (p *Progress) Stop() {
	close(p.StopCh)
}

func getProgressBar(p *Progress) string {
	width := 50
	progress := int((float64(p.Current) / float64(p.Total)) * float64(width))
	bar := make([]rune, width)
	for i := 0; i < width; i++ {
		if i < progress {
			bar[i] = '='
		} else {
			bar[i] = ' '
		}
	}
	return string(bar)
}
