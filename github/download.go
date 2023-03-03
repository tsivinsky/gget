package github

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func buildFileUrl(gh GitHubURL) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/%s", gh.Owner, gh.Repo, gh.Head, gh.PathToFile)
}

func DownloadFile(gh GitHubURL) ([]byte, error) {
	url := buildFileUrl(gh)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(gh.Lines) == 1 {
		line := gh.Lines[0]
		s := strings.Split(string(data), "\n")
		for i, l := range s {
			if uint(i+1) == line {
				data = []byte(l + "\n")
			}
		}
	}

	if len(gh.Lines) == 2 {
		startLine, endLine := gh.Lines[0], gh.Lines[1]
		s := strings.Split(string(data), "\n")
		s = s[startLine-1 : endLine]

		data = []byte(strings.Join(s, "\n") + "\n")
	}

	return data, nil
}
