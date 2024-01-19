package main

import (
	"fmt"
	"sync"
	"wget/funcs"

	"github.com/spf13/pflag"
)

func main() {
	// Use pflag instead of flag
	pflag.Parse()

	// Retrieve non-flag arguments
	args := pflag.Args()
	if len(args) < 1 {
		fmt.Println("Error: URL is required")
		return
	}

	url := ""

	// Use a separate wait group for the goroutines
	var wg sync.WaitGroup

	if *funcs.InputFile != "" {
		funcs.MultiDownloads()
	} else {
		for _, arg := range args {
			url = arg
			// Increment the new wait group for each goroutine
			wg.Add(1)
			// go funcs.SendSingleRequest(currentArg, &wg)
			// Call the downloadFile function with the parsed arguments
			go func() {
				err := funcs.DownloadFile(url, &wg)
				if err != nil {
					fmt.Printf("Error downloading file: %v\n", err)
					return
				}
			}()

		}
	}
	// // Call the downloadFile function with the parsed arguments
	// err := funcs.DownloadFile(url)
	// if err != nil {
	// 	fmt.Printf("Error downloading file: %v\n", err)
	// 	return
	// }

	// Wait for all downloads to complete
	wg.Wait()

	fmt.Println("\nDownload completed successfully\n")

	// Debugging prints
	fmt.Println("<-------------------------------->")
	fmt.Println("URL:", url)
	fmt.Println("flagO:", *funcs.SaveAs)
	fmt.Println("flagP:", *funcs.SaveDir)
	fmt.Println("flagRateLimit:", *funcs.RateLimit)
	fmt.Println("<-------------------------------->\n")

}
