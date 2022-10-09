package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("You need to provide url to file on GitHub")
		os.Exit(1)
	}

	uri, err := url.Parse(args[0])
	if err != nil {
		log.Fatal(err)
	}

	url := strings.Replace(uri.Path, "blob/", "", 1)
	url = fmt.Sprintf("https://raw.githubusercontent.com%s", url)

	data, err := getFileData(url)
	if err != nil {
		log.Fatal(err)
	}

	outFilePath, err := getArgumentValue("-o")
	if err != nil {
		// print to stdout
		fmt.Print(data)
	} else {
		// save to file
		if outFilePath == "" {
			// if user does not specify the file name for output, set it the same as file name on GitHub
			urlSlice := strings.Split(url, "/")
			outFilePath = urlSlice[len(urlSlice)-1]
		}

		fmt.Printf("Saving file to %s\n", outFilePath)

		err := saveOutputToFile(outFilePath, data)
		if err != nil {
			log.Fatal(err)
		}
	}
}
