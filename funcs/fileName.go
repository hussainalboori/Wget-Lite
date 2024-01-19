package funcs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getUniqueFileName(fileName string, path string, saveAs string) string {

	ext := getFileExtension(fileName)
	baseName := strings.TrimSuffix(fileName, ext)
	// Determine the filename
	if saveAs != "" {
		baseName = saveAs
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

func getFileExtension(fileName string) string {
	dotIndex := strings.LastIndex(fileName, ".")
	if dotIndex == -1 {
		return ""
	}
	return fileName[dotIndex:]
}
