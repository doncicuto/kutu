package config

//BinaryConfig stores the binary description,
//the command arguments to get its version and
//the Github repository information to check if
//version is up to date.
type BinaryConfig struct {
	Description      string
	VersionCommand   []string
	GithubReleaseURL string
	VersionPrefix    string
	DownloadURL      string
}

// Binaries List of binaries managed by Kutu
var Binaries = make(map[string]BinaryConfig)

func init() {

	Binaries["kubectl"] = BinaryConfig{
		Description:      "is a command line tool for controlling Kubernetes clusters.",
		VersionCommand:   []string{"version", "--client", "-o", "json"},
		GithubReleaseURL: "https://storage.googleapis.com/kubernetes-release/release/stable.txt",
		DownloadURL:      "https://storage.googleapis.com/kubernetes-release/release/%s/bin/%s/amd64/kubectl",
	}

	Binaries["skaffold"] = BinaryConfig{
		Description:      "handles the workflow for building, pushing and deploying your application, allowing you to focus on what matters most: writing code. ",
		VersionCommand:   []string{"version"},
		GithubReleaseURL: "https://api.github.com/repos/GoogleContainerTools/skaffold/releases",
		VersionPrefix:    "v",
		DownloadURL:      "https://storage.googleapis.com/skaffold/releases/%s/skaffold-%s-amd64",
	}

	Binaries["minikube"] = BinaryConfig{
		Description:      "is a tool that makes it easy to run Kubernetes locally.",
		VersionCommand:   []string{"version", "-o", "json"},
		GithubReleaseURL: "https://api.github.com/repos/kubernetes/minikube/releases",
		VersionPrefix:    "v",
		DownloadURL:      "https://github.com/kubernetes/minikube/releases/download/%s/minikube-%s-amd64",
	}

	Binaries["kustomize"] = BinaryConfig{
		Description:      "traverses a Kubernetes manifest to add, remove or update configuration options without forking.",
		VersionCommand:   []string{"version", "--short"},
		GithubReleaseURL: "https://api.github.com/repos/kubernetes-sigs/kustomize/releases",
		VersionPrefix:    "kustomize/v",
		DownloadURL:      "https://github.com/kubernetes-sigs/kustomize/releases/download/%s/%s_%s_amd64.tar.gz",
	}
}
