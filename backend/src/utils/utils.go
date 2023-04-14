package utils

import (
	"fmt"
	"os"
	"path"
)

var FrontentAddress string = "localhost:5173"

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

	dir, filename := path.Split(filesystemPath)
	ext := path.Ext(filename)
	base := filename[:len(filename)-len(ext)]
	i := 1

	for {
		newFilename := fmt.Sprintf("%s_%d%s", base, i, ext)
		newFilepath := path.Join(dir, newFilename)
		fmt.Println("Checking filepath:", newFilepath)
		if !FileExists(newFilepath) {
			return newFilepath
		}
		i++
	}
}

func createTemp() string {
	dirs := []string{"", "."}
	for _, dir := range dirs {
		println(f.Name())

	}
	return f.Name()
}
