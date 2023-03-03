package main

import (
	"errors"
	"flag"
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

	if urlPath == "" {
		cli.ExitWithError(errors.New("Usage: gget https://github.com/$user/$repo/blob/$head/file"), 1)
	}

	gh, err := github.ParseURL(urlPath)
	if err != nil {
		cli.ExitWithError(err, 1)
	}

	content, err := github.DownloadFile(*gh)
	if err != nil {
		cli.ExitWithError(err, 1)
	}

	var w io.Writer

	if outFileName != "" {
		cwd, err := os.Getwd()
		if err != nil {
			cli.ExitWithError(err, 1)
		}
		outFileName = path.Join(cwd, outFileName)

		f, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			cli.ExitWithError(err, 1)
		}
		defer f.Close()
		w = f
	} else {
		w = os.Stdout
	}

	_, err = w.Write(content)
	if err != nil {
		cli.ExitWithError(err, 1)
	}
}
