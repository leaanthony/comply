package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

func main() {

	fmt.Printf("Comply © 2018-Present Lea Anthony. All rights reserved.\n\n")
	// Get current working directory
	cwd, err := os.Getwd()
	CheckError(err)

	// Check vendor exists and is a directory
	vendorDir := filepath.Join(cwd, "vendor")
	stat, err := os.Stat(vendorDir)
	if err != nil {
		Fatal("The vendor directory does not exist")
	}

	if !stat.IsDir() {
		Fatal("vendor is not a directory")
	}

	// Calculate license dir
	licenseDir := filepath.Join(cwd, "licenses")

	// If dir already exists, remove it
	err = os.RemoveAll(licenseDir)

	// Find license files in vendor directory
	licenseFiles := []string{}
	filepath.Walk(vendorDir, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			match, err := regexp.MatchString("(?i)LICENSE*", f.Name())
			if err == nil && match {
				// Strip vendor dir
				localDir := filepath.Dir(path)[len(vendorDir)+1:]

				dirToCreate := filepath.Join(licenseDir, localDir)
				err = os.MkdirAll(dirToCreate, 0644)
				CheckError(err)

				dstFile := filepath.Join(dirToCreate, f.Name())
				err = Copy(path, dstFile)
				CheckError(err)

				licenseFiles = append(licenseFiles, filepath.Join(localDir, f.Name()))
			}
		}
		return nil
	})

	if len(licenseFiles) > 0 {
		fmt.Printf("Copied %d license file(s) into '%s':\n\n", len(licenseFiles), licenseDir)
		for _, name := range licenseFiles {
			fmt.Printf("  %s\n", name)
		}
		fmt.Println()
	} else {
		fmt.Printf("No license files found.\n\n")
	}

	fmt.Println("Thanks for using Comply! I hope it was useful for you.")

}

// Fatal prints the given error message then exists
func Fatal(msg string) {
	fmt.Printf("Fatal: %s\n", msg)
	os.Exit(1)
}

// CheckError handles the case when the given error is not null
func CheckError(err error) {
	if err != nil {
		Fatal(err.Error())
	}
}

// Copy the src file to the dst file
// Credit: https://stackoverflow.com/a/21061062
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
