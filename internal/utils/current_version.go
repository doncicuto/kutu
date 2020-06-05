package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type minikubeVersion struct {
	Commit          string
	MinikubeVersion string
}

type kubectlVersion struct {
	ClientVersion struct {
		Major        string
		Minor        string
		GitVersion   string
		GitCommit    string
		GitTreeState string
		BuildDate    string
		GoVersion    string
		Compiler     string
		Platform     string
	}
}

// CurrentVersion gets the current version for our binary
// executing the command in a shell and parsing the output.
func CurrentVersion(binary string, versionCommand []string) (string, error) {
	switch binary {
	case "kubectl":
		out, err := exec.Command(binary, versionCommand...).Output()
		if err == nil {
			var version kubectlVersion
			err := json.Unmarshal(out, &version)
			if err == nil {
				return string(version.ClientVersion.GitVersion), nil
			}
		}
	case "skaffold":
		out, err := exec.Command(binary, versionCommand...).Output()
		if err == nil {
			return (strings.TrimRight(string(out), "\n")), nil
		}
	case "minikube":
		out, err := exec.Command(binary, versionCommand...).Output()
		if err == nil {
			var version minikubeVersion
			err := json.Unmarshal(out, &version)
			if err == nil {
				return string(version.MinikubeVersion), nil
			}
		}
	case "kustomize":
		out, err := exec.Command(binary, versionCommand...).Output()
		if err == nil {
			re := regexp.MustCompile(`{(.*?) `)
			match := re.FindStringSubmatch(string(out))
			if len(match) >= 2 {
				return match[1], nil
			}
		}
	default:
		return "", fmt.Errorf("Kutu doesn't know how to check or update %s", binary)
	}
	return "", errors.New("Could not parse version")
}
