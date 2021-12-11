package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	outFilePath, err := getArgumentValue("-o")
	if err != nil {
		// print to stdout
		fmt.Print(string(data))
	} else {
		// save to file
		if outFilePath == "" {
			// if user does not specify the file name for output, set it the same as file name on GitHub
			urlSlice := strings.Split(url, "/")
			outFilePath = urlSlice[len(urlSlice)-1]
		}

		log.Printf("outFilePath: %s\n", outFilePath)

		f, err := os.OpenFile(outFilePath, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		f.Write(data)
	}
}

func getArgumentValue(arg string) (value string, err error) {
	args := os.Args[1:]

	for i, val := range args {
		if val == arg {
			defer func() {
				if err := recover(); err != nil {
					value = ""
					err = nil
				}
			}()

			value = args[i+1]

			return
		}
	}

	return "", errors.New("No " + arg + " provided")
}
