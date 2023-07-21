package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func listFilesAndDirectories(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	fmt.Println("Listing files and directories in:", path)
	for _, file := range files {
		name := file.Name()
		if file.IsDir() {
			name += "/"
		}
		fmt.Println(name)
	}
	return nil
}
func downloadFile(url, destination string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Println("Downloaded files:", destination)
	return nil

}
func main() {
	var path string
	fmt.Print("Enter the path: ")
	fmt.Scanln(&path)
	err := listFilesAndDirectories(path)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	var fileURL string
	fmt.Print("Enter the file URL to download: ")
	fmt.Scanln(&fileURL)

	var destinationPath string
	fmt.Print("Enter the destination path: ")
	fmt.Scanln(&destinationPath)

	err = downloadFile(fileURL, destinationPath)
	if err != nil {
		fmt.Print("Error downloading file:", err)
		os.Exit(1)
	}
}
