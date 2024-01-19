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


//////////////////////////////////////////////////
# wget-Lite

This is a lightweight basic version of GNU `wget` written in Go.
This Go code provides a command-line tool to download a file from a given URL. It includes features such as a progress bar to track the download progress and the option to limit the download speed.


## Implemented Options

1. `-B` : Enables Silent Mode. All output will be written to a file called `Log.txt`.
2. `-O` and `-P` : rename the file under a different name and under a different path respectively.
3. The project implements a rate limiter (still in works). Basically the program can control the speed of the download by using the flag `--rate-limit`. If you download a huge file you can limit the speed of your download, preventing the program from using the full possible bandwidth of your connection.
4. Downloading different files is possible. For this the program will receive the `-i` flag followed by a file name that will contain all links that are to be downloaded. The downloads will be done in async.
5. Finally, the project is able to mirror a website using the `-mirror` tag (in works).

## More Information

* The project is written in pure golang, with a makefile for creating a build by running the `make` command.
* Multiple different external repositories were used, such as `github.com/progressbar/v3` for progress bar functionality.
* The project tried to use as much of the stdlibs as possible, but had to resort to external packages for some functionality like HTML parsing.

## Authors
- amali01 (Amjad Ali)
- husalboori
- aalsendi
- hnabeel

/////////////////////////////////////////////////

<h1 align="center">My-Ls Project</h1>


<p align="center">
    <img src="Design 1 (1).png" alt="Wget-light Logo" />
</p>

<h2 align="center">About The Project</h2>
<h4 align="center">My-ls is a project that aims to create a custom ls command using Go.</h4>

## Getting Started
You can run the My-Ls project with the following command:
```console
git clone git clone https://github.com/amali01/my-ls-1.git
cd my-ls
```

## Usage
```
go run . [OPTIONS] [FILE|DIR]
```
#### Directory Structure:
```console
─ my-ls-1/
├── get/
│   ├── myInfo.go
│   ├── myStruct.go
│   ├── mySort.go
│   ├── CleanInput.go  
│   ├── myPrint.go
│   ├── myFlags.go
│   └── myColors.go
|
├── main.go
├── go.mod
├── README.md
└── ...
```
## Examples
Here are some examples of how to use My-Ls:

- List in long format (equivalent to ls -l):
```
go run . -l 
```
- List in reverse order (equivalent to ls -r):
```
go run . -r 
```
- List contents of a specific folder (e.g., folder/):
```
go run . folder/

```
- Combine multiple options (e.g., -lraRt):
```
go run . -lraRt
```

## Available options

* <code>-l</code> - list with long format
* <code>-r</code> - list in reverse order
* <code>-a</code> - list all files including hidden file starting with '.'
* <code>-R</code> - list recursively directory tree
* <code>-t</code> - sort by time & date

## Additional information

- Only standard go packages were in use.
- For faster audit use bash audit.sh   

## Authors

- amali01 (Amjad Ali)
- emahfoodh (Eman Mahfoodh)
