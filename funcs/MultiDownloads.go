package funcs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func MultiDownloads() {
	dfile, err := os.Open(*InputFile)
	var wg sync.WaitGroup
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("[MULTIREQSEND ERROR]: File dont exist")
		} else {
			log.Fatal("MULTIREQSEND ERROR]:", err)
		}
	}
	defer dfile.Close()

	scanner := bufio.NewScanner(dfile)
	urlArray := []string{}

	for scanner.Scan() {
		urlArray = append(urlArray, scanner.Text())
	}

	for _, url := range urlArray {
		wg.Add(1)
		go DownloadFile(url, &wg)
	}
	wg.Wait()

	if *InputFile != "" {
		// Log the download completion
		fmt.Printf("\nDownloaded [%d] files\n", len(urlArray))
		fmt.Printf("finished at %s\n", time.Now().Format("2006-01-02 15:04:05"))
	}
}
