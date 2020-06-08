package utils_test

import (
	"testing"

	"github.com/muultipla/kutu/internal/config"
	"github.com/muultipla/kutu/internal/utils"
)

//TestBinaryConfig stores the binary description,
//the command arguments to get its version and
//the Github repository information to check if
//version is up to date.
type TestBinaryConfig struct {
	Description      string
	VersionCommand   []string
	GithubReleaseURL string
	VersionPrefix    string
	OldVersion       string
	LatestVersion    string
}

// TestBinaries List of binaries managed by Kutu
var TestBinaries = make(map[string]TestBinaryConfig)

func init() {

	TestBinaries["kubectl"] = TestBinaryConfig{
		LatestVersion: "v1.18.3",
		OldVersion:    "v1.17.6",
	}

	TestBinaries["skaffold"] = TestBinaryConfig{
		LatestVersion: "v1.10.1",
		OldVersion:    "v1.9.1",
	}

	TestBinaries["minikube"] = TestBinaryConfig{
		LatestVersion: "v1.11.0",
		OldVersion:    "v1.10.1",
	}

	TestBinaries["kustomize"] = TestBinaryConfig{
		LatestVersion: "kustomize/v3.6.1",
		OldVersion:    "kustomize/v3.6.0",
	}
}

func TestCheckVersion(t *testing.T) {
	for binary, info := range config.Binaries {
		version := TestBinaries[binary].OldVersion
		got, _ := utils.CheckVersion(binary, info, version)
		if !got.Outdated {
			t.Errorf("CheckVersion %s for version %s was %t ; want true", binary, version, got.Outdated)
		}

		version = TestBinaries[binary].LatestVersion
		got, _ = utils.CheckVersion(binary, info, version)
		if got.Outdated {
			t.Errorf("CheckVersion %s for version %s was %t ; want false", binary, version, got.Outdated)
		}
	}
}
