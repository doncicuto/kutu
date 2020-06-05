package config

import (
	"strings"

	"github.com/tcnksm/go-latest"
)

//BinaryConfig stores the binary description,
//the command arguments to get its version and
//the Github repository information to check if
//version is up to date.
type BinaryConfig struct {
	Description    string
	VersionCommand []string
	GithubTag      latest.GithubTag
}

// Binaries List of binaries managed gby Kutu
var Binaries = make(map[string]BinaryConfig)

func init() {

	Binaries["kubectl"] = BinaryConfig{
		Description:    "is a command line tool for controlling Kubernetes clusters.",
		VersionCommand: []string{"version", "--client", "-o", "json"},
		GithubTag: latest.GithubTag{
			Owner:             "kubernetes",
			Repository:        "kubectl",
			FixVersionStrFunc: latest.DeleteFrontV(),
		},
	}

	Binaries["skaffold"] = BinaryConfig{
		Description:    "handles the workflow for building, pushing and deploying your application, allowing you to focus on what matters most: writing code. ",
		VersionCommand: []string{"version"},
		GithubTag: latest.GithubTag{
			Owner:             "GoogleContainerTools",
			Repository:        "skaffold",
			FixVersionStrFunc: latest.DeleteFrontV(),
		},
	}

	Binaries["minikube"] = BinaryConfig{
		Description:    "is a tool that makes it easy to run Kubernetes locally.",
		VersionCommand: []string{"version", "-o", "json"},
		GithubTag: latest.GithubTag{
			Owner:             "kubernetes",
			Repository:        "minikube",
			FixVersionStrFunc: latest.DeleteFrontV(),
		},
	}

	Binaries["kustomize"] = BinaryConfig{
		Description:    "traverses a Kubernetes manifest to add, remove or update configuration options without forking.",
		VersionCommand: []string{"version", "--short"},
		GithubTag: latest.GithubTag{
			Owner:      "kubernetes-sigs",
			Repository: "kustomize",
			FixVersionStrFunc: func(version string) string {
				return strings.Replace(version, `kustomize/v`, "", 1)
			},
		},
	}
}
