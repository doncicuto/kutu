package utils

import (
	"fmt"
	"strings"

	"github.com/muultipla/kutu/internal/config"
	"github.com/tcnksm/go-latest"
)

// CheckVersion checks if our current version is outdated
// using tcnksm's awesome go-latest package.
// CheckVersion gets our current version and formats it
func CheckVersion(binary string, info config.BinaryConfig, current string) (*latest.CheckResponse, error) {

	switch binary {
	case "kubectl":
		ourCurrentVersion := strings.Replace(current, "v", "", 1)
		res, err := latest.Check(&info.GithubTag, ourCurrentVersion)
		if err == nil {
			if !res.Outdated {
				res.Current = ourCurrentVersion
			}
			return res, err
		}
		return nil, err

	case "skaffold":
		var current, err = CurrentVersion(binary, info.VersionCommand)
		if err == nil {
			ourCurrentVersion := strings.TrimRight(strings.Replace(current, "v", "", 1), "\n")
			res, err := latest.Check(&info.GithubTag, ourCurrentVersion)
			if err == nil {
				if !res.Outdated {
					res.Current = ourCurrentVersion
				}
				return res, err
			}
		}
		return nil, err

	case "minikube":
		var current, err = CurrentVersion(binary, info.VersionCommand)
		if err == nil {
			ourCurrentVersion := strings.Replace(current, "v", "", 1)
			res, err := latest.Check(&info.GithubTag, ourCurrentVersion)
			if err == nil {
				if !res.Outdated {
					res.Current = ourCurrentVersion
				}
				return res, err
			}
		}
		return nil, err

	case "kustomize":
		var current, err = CurrentVersion(binary, info.VersionCommand)
		if err == nil {
			ourCurrentVersion := strings.Replace(current, `kustomize/v`, "", 1)
			res, err := latest.Check(&info.GithubTag, ourCurrentVersion)
			if err == nil {
				if !res.Outdated {
					res.Current = ourCurrentVersion
				}
				return res, err
			}
		}
		return nil, err

	default:
		return nil, fmt.Errorf("Kutu doesn't know how to check or update %s", binary)
	}
}
