package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/go-github/github"

	"github.com/muultipla/kutu/internal/config"
)

// CheckVersionResponse tells us if current version is outdated
// and which version is the latest
type CheckVersionResponse struct {
	Outdated bool
	Latest   string
	Current  string
}

func getKubeCtlStableRelease(url string) (*string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	version := strings.TrimRight(string(body), "\n")
	if strings.HasPrefix(version, "v") {
		return &version, err
	}
	return nil, fmt.Errorf("Error: could not get latest releases from GitHub (API rate exceeded?), try again later")
}

func getGitHubReleaseInfo(url string) ([]github.RepositoryRelease, error) {
	var releases []github.RepositoryRelease
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &releases)
	if err == nil {
		return releases, err
	}
	return nil, fmt.Errorf("Error: could not get latest releases from GitHub (API rate exceeded?), try again later")
}

// CheckVersion checks if our current version is outdated
func CheckVersion(binary string, info config.BinaryConfig, current string) (*CheckVersionResponse, error) {
	var latest *string
	var err error
	var releases []github.RepositoryRelease

	if binary != "kubectl" {
		releases, err = getGitHubReleaseInfo(info.GithubReleaseURL)
		if err == nil {
			for _, release := range releases {
				latest = release.TagName
				if strings.HasPrefix(*latest, info.VersionPrefix) && !*release.Prerelease {
					break
				}
			}
		}
	} else {
		latest, err = getKubeCtlStableRelease(info.GithubReleaseURL)
	}

	if err == nil {
		return &CheckVersionResponse{
			Latest:   *latest,
			Current:  current,
			Outdated: current != *latest,
		}, nil
	}

	return nil, err
}
