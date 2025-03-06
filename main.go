package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dirPath := flag.String("dir", ".", "Directory to rename files in")
	prefix := flag.String("prefix", "", "Prefix to add to file names")
	suffix := flag.String("suffix", "", "Suffix to add to file names")

	flag.Parse()

	if *dirPath == "" {
		exitWithError("Directory is required")
	}

	if *prefix == "" && *suffix == "" {
		exitWithError("Either prefix or suffix is required")
	}

	dir, err := os.Open(*dirPath)
	if err != nil {
		exitWithError(fmt.Sprintf("Error opening directory %s: %v", *dirPath, err))
	}
	defer dir.Close()

	files, err := dir.Readdir(0)
	if err != nil {
		exitWithError(fmt.Sprintf("Error reading directory %s: %v", *dirPath, err))
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		fileExt := filepath.Ext(file.Name())
		oldPath := filepath.Join(*dirPath, fileName)
		newPath := *prefix + strings.TrimSuffix(fileName, fileExt) + *suffix + fileExt

		err = os.Rename(oldPath, newPath)
		if err != nil {
			displayError(fmt.Sprintf("Error renaming file %s: %v", fileName, err))
			continue
		}
	}
}

func exitWithError(message string) {
	displayError(message)
	os.Exit(1)
}

func displayError(message string) {
	fmt.Fprintln(os.Stderr, "Error: "+message)
}
