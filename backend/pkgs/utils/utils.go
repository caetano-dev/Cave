package utils

import (
	"os"
	"path/filepath"
	"strconv"
)

// FileExists is the function that determines if a file already exists with a given name.
func FileExists(filesystemPath string) bool {
	if _, err := os.Stat(filesystemPath); os.IsNotExist(err) {
		return false
	}
	return true
}

// UniqueFilesystemPath is the function that creates a unique filepath by changing the filename.
func UniqueFilesystemPath(filesystemPath string) string {
	if !FileExists(filesystemPath) {
		return filesystemPath
	}

	base := filesystemPath[:len(filesystemPath)-len(filepath.Ext(filesystemPath))] // remove extension
	ext := filepath.Ext(filesystemPath)
	i := 1

	for {
		newfilesystemPath := base + "_" + strconv.Itoa(i) + ext
		if !FileExists(newfilesystemPath) {
			return newfilesystemPath
		}
		i++
	}
}

var FrontentAddress string = "localhost:5173"
