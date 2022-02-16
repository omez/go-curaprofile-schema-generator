package cura

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

type releasesResponse = []releaseType

type releaseType struct {
	Name    string `json:"name,omitempty"`
	TagName string `json:"tag_name,omitempty"`
	Body    string `json:"body,omitempty"`
}

var releaseFilter = regexp.MustCompile("^(\\d+)\\.(\\d+)(\\.\\d+)?$")

func GetReleases() ([]string, error) {

	// fetch releases
	var releases releasesResponse
	err := fetchJson("https://api.github.com/repos/Ultimaker/Cura/releases", &releases)
	if err != nil {
		return nil, err
	}

	// filter releases by regex
	var releaseNames []string
	for _, release := range releases {
		if releaseFilter.MatchString(release.Name) {
			releaseNames = append(releaseNames, release.TagName)
		}
	}
	return releaseNames, nil
}

func GetLatestRelease() (string, error) {
	releases, err := GetReleases()
	if err != nil {
		return "", err
	}
	if len(releases) < 1 {
		return "", fmt.Errorf("no releases found")
	}
	return releases[0], nil
}

func getDefinitionFromCuraGithubRepo(def interface{}, path string, release string) error {
	// generate relative url
	fullUri, err := url.Parse("https://raw.githubusercontent.com/Ultimaker/Cura/" + release + "/" + path)
	if err != nil {
		return err
	}
	err = fetchJson(fullUri.String(), def)
	if err != nil {
		return err
	}
	return nil
}

func getLatestDefinitionFromCuraGithubRepo(def interface{}, path string) error {
	if latestRelease, err := GetLatestRelease(); err != nil {
		return err
	} else if err := getDefinitionFromCuraGithubRepo(def, path, latestRelease); err != nil {
		return err
	} else {
		return nil
	}
}

func fetchJson(uri string, v interface{}) error {
	// fetch URI
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}
