package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

func getFileData(fileUrl string) (string, error) {
	resp, err := http.Get(fileUrl)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func saveOutputToFile(filePath string, output string) error {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(output)

	return nil
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
