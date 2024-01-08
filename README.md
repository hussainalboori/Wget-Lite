# Download File with Progress Bar and Rate Limiting

This Go code provides a command-line tool to download a file from a given URL. It includes features such as a progress bar to track the download progress and the option to limit the download speed.

## Prerequisites

Before running the code, make sure you have Go installed on your system.

## Usage

To use the tool, run the following command:

```shell
go run main.go -url <URL> [-O <SaveAs>] [-P <SaveDir>] [--rate-limit <RateLimit>]
```

- `-url`: The URL of the file to download (required).
- `-O`: (Optional) Save the downloaded file with a different name.
- `-P`: (Optional) Directory to save the downloaded file.
- `--rate-limit`: (Optional) Limit the download speed in bytes per second.

## Example

To download a file from a URL, run the following command:

```shell
go run main.go -url https://example.com/file.zip
```

This will download the file from the specified URL and display a progress bar indicating the download progress. Once the download is complete, it will show a success message.

You can also specify additional options:

```shell
go run main.go -url https://example.com/file.zip -O my_file.zip -P /path/to/save -rate-limit 10240
```

In this example, the downloaded file will be saved as `my_file.zip`, in the `/path/to/save` directory, and the download speed will be limited to 10240 bytes per second.

## Explanation

The code uses the `net/http` package to make an HTTP GET request to the specified URL. It retrieves the response body and writes it to a file while updating a progress bar.

If the `--rate-limit` option is provided, the code wraps the writer with a rate-limited writer, which limits the write speed to the specified rate. It also calculates and displays the throughput (bytes per second) every second.

The progress bar is implemented using the `progressBar` struct, which keeps track of the total file size and the number of bytes downloaded. It updates and prints the progress percentage as data is written to the file.

The code also handles creating the save directory if it doesn't exist and determines the filename to save the downloaded file.

Finally, the code uses the `flag` package to parse command-line arguments and executes the `downloadFile` function with the provided options.

Please note that this code does not handle all possible error cases and does not include extensive error handling. It's intended as a basic example to demonstrate the functionality.