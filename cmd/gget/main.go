package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/tsivinsky/gget/cli"
	"github.com/tsivinsky/gget/github"
)

var (
	outFileName string
)

func main() {
	flag.StringVar(&outFileName, "o", "", "-o path/to/file")
	flag.Parse()

	urlPath := flag.Arg(0)

	gh, err := github.ParseURL(urlPath)
	if err != nil {
		cli.ExitWithError(err, 1)
	}

	if outFileName == "" {
		fname := path.Base(gh.PathToFile)
		outFileName = fname
	}

	cwd, err := os.Getwd()
	if err != nil {
		cli.ExitWithError(err, 1)
	}

	outFileName = path.Join(cwd, outFileName)

	content, err := github.DownloadFile(*gh)
	if err != nil {
		cli.ExitWithError(err, 1)
	}

	var w io.Writer

	f, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		cli.ExitWithError(err, 1)
	}
	defer f.Close()
	w = f

	_, err = w.Write(content)
	if err != nil {
		cli.ExitWithError(err, 1)
	}

	fmt.Printf("Saved file to %s\n", f.Name())
}
