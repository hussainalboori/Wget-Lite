# 
<h1 align="center">Wget-Lite Project</h1>


<p align="center">
    <img src="./assets/Design4.png" alt="Wget-light Logo"  />
</p>

<h2 align="center">About The Project</h2>

This is a lightweight basic version of GNU `wget` written in Go.
It provides a command-line tool to download a file from a given URL. It includes features such as a progress bar to track the download progress and the option to limit the download speed.

## Usage

To use the tool, run the following command:

```shell
go run main.go  <url>
```

- `url`: The URL of the file to download.


#### Directory Structure:
```console
─ Wget/
│
├── funcs/
│   ├── download.go
│   ├── fileName.go
│   ├── globs.go
│   ├── MultiDownloads.go  
│   ├── outPut.go
│   └── ProgressBar.go
|
├── assetes/
│   └── project-logos
│
├── main.go
├── go.mod
├── README.md
└── ...
```

## Implemented Options

### Silent Mode (`-B`)

Enable Silent Mode to redirect all output to a file named `Wget-light-log.txt`.

Example:
```bash
go run main.go -B https://example.com/file.txt
```

### Save As (`-O`)

Specify a different name for the downloaded file.

Example:
```bash
go run main.go -O https://example.com/file.txt
```

### Save As (`-P`)

Specify a different name for the downloaded file.

Example:
```bash
go run main.go -P /path/to/directory https://example.com/file.txt
```


### Rate Limit (`--rate-limit`)
 
Limit the download speed in bytes per second.

Example: 
```bash
go run main.go --rate-limit 100000 https://example.com/file.txt
```


### Multiple Downloads (`-i`)

Download multiple files specified in a file.

Example:
```bash
go run main.go -i urls.txt
```

### Website Mirroring (`--mirror`)

Mirror a website, downloading its entire content.

Example:
```bash
go run main.go --mirror https://example.com
```


## More Information

* The project is written in pure Golang, utilizing the language's features and libraries to achieve its functionality.

* To build the project, use the provided Makefile by running the `make` command. This will compile the project and create an executable.

* To clean up all generated files, including the compiled binary and any downloaded files, use the following command:
  ```bash
  make clean
  ```

## Authors
- amali (Amjad Ali)
- husalboori
- aalsendi
- hnabeel
