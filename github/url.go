package github

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

var (
	ErrNoFilePath = errors.New("No path to file in url")
)

type GitHubURL struct {
	Owner      string
	Repo       string
	Head       string
	PathToFile string
	Lines      []uint
}

func parseLine(line string) (uint, error) {
	s := strings.ReplaceAll(line, "L", "")
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return uint(n), nil
}

func parseLinesFrag(frag string) ([]uint, error) {
	var lines []uint

	if !strings.HasPrefix(frag, "L") {
		return lines, nil
	}

	s := strings.Split(frag, "-")
	startLine, err := parseLine(s[0])
	if err != nil {
		return nil, err
	}

	lines = append(lines, startLine)

	if len(s) > 1 {
		endLine, err := parseLine(s[1])
		if err != nil {
			return nil, err
		}

		lines = append(lines, endLine)
	}

	return lines, nil
}

func ParseURL(fileUrl string) (*GitHubURL, error) {
	uri, err := url.Parse(fileUrl)
	if err != nil {
		return nil, err
	}

	s := strings.Split(uri.Path, "/")
	gh := new(GitHubURL)

	s = s[1:] // first element is empty

	if len(s) <= 4 {
		return nil, ErrNoFilePath
	}

	gh.Owner = s[0]
	gh.Repo = s[1]
	gh.Head = s[3]
	gh.Lines, err = parseLinesFrag(uri.Fragment)
	if err != nil {
		return nil, err
	}

	fullPath := s[4:]
	gh.PathToFile = strings.Join(fullPath, "/")

	return gh, nil
}
