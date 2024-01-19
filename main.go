package main

import (
	"fmt"
	"wget/funcs"

	"github.com/spf13/pflag"
)

var (
	saveAs    = pflag.StringP("O", "O", "", "save the downloaded file with a different name")
	saveDir   = pflag.StringP("P", "P", "", "directory to save the downloaded file")
	rateLimit = pflag.Int64P("rate-limit", "R", 0, "limit the download speed in bytes per second")
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

	url := args[0]

	// Debugging prints
	fmt.Println("URL:", url)
	fmt.Println("flagO:", *saveAs)
	fmt.Println("flagP:", *saveDir)
	fmt.Println("flagRateLimit:", *rateLimit)

	// Call the downloadFile function with the parsed arguments
	err := funcs.DownloadFile(url, *saveAs, *saveDir, *rateLimit)
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
		return
	}

	fmt.Println("\nDownload completed successfully\n")
}
