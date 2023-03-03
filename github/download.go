package github

import (
	"fmt"
	"io"
	"net/http"
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

	return data, nil
}
