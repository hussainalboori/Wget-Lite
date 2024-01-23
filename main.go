package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"wget/funcs"

	"github.com/spf13/pflag"
)

func main() {
	// Use pflag instead of flag
	pflag.Parse()

	// Retrieve non-flag arguments
	args := pflag.Args()

	//fixing the path
	*funcs.SaveDir, _ = funcs.ExpandTilde(*funcs.SaveDir)

	inputURL := ""

	// Use a separate wait group for the goroutines
	var wg sync.WaitGroup

	// Mirror the website if the Mirror flag is set
	if *funcs.Mirror {

		inputURL = args[0]
		if *funcs.SaveDir == "" {
			FolderName := filepath.Base(inputURL)
			*funcs.SaveDir = FolderName
		}
		mirrorOptions := funcs.MirrorOptions{
			RejectFileTypes: strings.Split(*funcs.Reject, ","),
			ExcludeDirs:     strings.Split(*funcs.Exclude, ","),
		}

		err := funcs.Mirroring(inputURL, mirrorOptions)
		if err != nil {
			fmt.Println("Error mirroring website:", err)
			return
		}

		// Log the download completion
		fmt.Printf("\nDownloaded [%s]\n", inputURL)
		fmt.Printf("finished at %s\n\n", time.Now().Format("2006-01-02 15:04:05"))

	} else if *funcs.InputFile != "" {
		funcs.MultiDownloads()
	} else {
		for _, arg := range args {
			inputURL = arg
			// Increment the new wait group for each goroutine
			wg.Add(1)
			// Call the downloadFile function with the parsed arguments
			go func() {
				err := funcs.DownloadFile(inputURL, &wg)
				if err != nil {
					fmt.Printf("Error downloading file: %v\n", err)
					return
				}
			}()
			wg.Wait()

		}
	}

	// // Wait for all downloads to complete
	// wg.Wait()

	// fmt.Println("\nDownload completed successfully\n")

	// // Debugging prints
	// fmt.Println("<-------------------------------->")
	// fmt.Println("URL:", inputURL)
	// fmt.Println("flagO:", *funcs.SaveAs)
	// fmt.Println("flagP:", *funcs.SaveDir)
	// fmt.Println("flagRateLimit:", *funcs.RateLimit)
	// fmt.Println("<-------------------------------->\n")
	// wg.Wait()
	os.Exit(0)
}
