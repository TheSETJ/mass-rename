package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nExamples:")
		fmt.Fprintln(os.Stderr, "  mass-rename --prefix new_\t# Prepends \"new_\" to all file names")
		fmt.Fprintln(os.Stderr, "  mass-rename --suffix _backup\t# Appends \"_backup\" to all file names")
	}

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

		oldFileName := file.Name()
		oldPath := filepath.Join(*dirPath, oldFileName)

		fileExt := filepath.Ext(file.Name())

		newFileName := *prefix + strings.TrimSuffix(oldFileName, fileExt) + *suffix + fileExt
		newPath := filepath.Join(*dirPath, newFileName)

		err = os.Rename(oldPath, newPath)
		if err != nil {
			displayError(fmt.Sprintf("Error renaming file %s: %v", oldPath, err))
			continue
		}

		fmt.Printf("Renamed %s to %s\n", oldPath, newPath)
	}
}

func exitWithError(message string) {
	displayError(message)
	os.Exit(1)
}

func displayError(message string) {
	fmt.Fprintln(os.Stderr, "Error: "+message)
}
