package github

import (
	"errors"
	"net/url"
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

	fullPath := s[4:]
	gh.PathToFile = strings.Join(fullPath, "/")

	return gh, nil
}
