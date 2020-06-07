package utils

import (
	"fmt"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/cavaliercoder/grab"
)

// Update downloads new binaries from GitHub repositories
func Update(binary string, filePath string, version string) error {
	var url string
	os := runtime.GOOS

	switch binary {
	case "kubectl":
		url = fmt.Sprintf("https://storage.googleapis.com/kubernetes-release/release/%s/bin/%s/amd64/kubectl", version, os)

	case "skaffold":
		url = fmt.Sprintf("https://storage.googleapis.com/skaffold/releases/%s/skaffold-%s-amd64", version, os)

	case "minikube":
		url = fmt.Sprintf("https://github.com/kubernetes/minikube/releases/download/%s/minikube-%s-amd64", version, os)

	case "kustomize":
		url = fmt.Sprintf("https://github.com/kubernetes-sigs/kustomize/releases/download/%s/%s_%s_amd64.tar.gz", version, strings.Replace(version, "/", "_", 1), os)
	}

	if binary == "kustomize" {
		tmpFilePath := fmt.Sprintf("/tmp/%s.tar.gz", strings.Replace(version, "/", "_", 1))
		_, err := grab.Get(tmpFilePath, url)

		if err == nil {
			_, cmdErr := exec.Command("tar", "xzf", tmpFilePath, "-C", path.Dir(filePath)).Output()
			if cmdErr != nil {
				return fmt.Errorf("Could not extract kustomize, you may not have enough privileges. Please re-run it using sudo")
			}
			return cmdErr
		}
		return err
	}
	_, err := grab.Get(filePath, url)
	if err != nil {
		return fmt.Errorf("Could not update %s, you may not have enough privileges. Please re-run it using sudo", binary)
	}
	return err
}
