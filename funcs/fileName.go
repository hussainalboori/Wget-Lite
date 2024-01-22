package funcs

import (
	"fmt"
	"mime"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func getUniqueFileName(fileName, path, SaveAs, ext string) string {

	baseName := strings.TrimSuffix(fileName, ext)
	// Determine the filename
	if SaveAs != "" {
		baseName = SaveAs
	}
	newFileName := fmt.Sprintf("%s%s", baseName, ext)
	fullPath := filepath.Join(path, newFileName)

	// Check if the file already exists
	_, err := os.Stat(fullPath)
	if err == nil {
		// File already exists, add a number to the filename until a unique filename is found
		i := 1
		for {
			newFileName = fmt.Sprintf("%s_%d%s", baseName, i, ext)
			newFullPath := filepath.Join(path, newFileName)

			_, err := os.Stat(newFullPath)
			if err != nil {
				return newFullPath
			}
			i++
		}
	}

	return fullPath
}

// getFileExtension returns the file extension based on the Content-Type header or URL file extension
func getFileExtension(contentType, urlPath string) string {
	// Try to extract the file extension from the Content-Type header
	ext, _ := mime.ExtensionsByType(contentType)
	if len(ext) > 0 {
		return ext[0]
	}

	// If the Content-Type header doesn't provide a file extension, try to extract it from the URL
	urlExt := filepath.Ext(urlPath)
	if urlExt != "" {
		return urlExt
	}

	// Default to no extension or determine based on other criteria
	return ""
}

func ExpandTilde(path string) (string, error) {
	if len(path) > 0 && path[0] == '~' {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		return filepath.Join(usr.HomeDir, path[1:]), nil
	}
	return path, nil
}
