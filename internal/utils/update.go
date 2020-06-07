package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/cavaliercoder/grab"
)

// Update downloads new binaries from GitHub repositories
func Update(binary string, filePath string, downloadURL string, version string) error {
	var url, tmpFilePath string
	var fi os.FileInfo
	goos := runtime.GOOS

	// Set download url and temp file
	if binary == "kustomize" {
		url = fmt.Sprintf(downloadURL, version, strings.Replace(version, "/", "_", 1), goos)
		tmpFilePath = fmt.Sprintf("/tmp/%s.tar.gz", strings.Replace(version, "/", "_", 1))
	} else {
		url = fmt.Sprintf(downloadURL, version, goos)
		tmpFilePath = fmt.Sprintf("/tmp/%s", binary)
	}

	// Download release file
	_, err := grab.Get(tmpFilePath, url)
	if err != nil {
		return fmt.Errorf("Could not download %s. Maybe url %s is no longer valid", binary, url)
	}

	// Get current permissions
	fi, err = os.Lstat(filePath)
	if err != nil {
		return fmt.Errorf("Could not check permissions for %s. Please re-run kutu using sudo", filePath)
	}
	perms := fi.Mode().Perm()

	// Extract or move downloaded file
	if binary == "kustomize" {
		// Extract file
		_, err = exec.Command("tar", "xzf", tmpFilePath, "-C", path.Dir(filePath)).Output()
		if err != nil {
			return fmt.Errorf("Could not replace %s, you may not have enough privileges. Please re-run kutu using sudo", binary)
		}
	} else {
		// Move file
		err = os.Rename(tmpFilePath, filePath)
		if err != nil {
			return fmt.Errorf("Could not replace %s. Please re-run kutu using sudo", filePath)
		}
	}

	// Set new permissions
	err = os.Chmod(filePath, perms)
	if err != nil {
		return fmt.Errorf("Could not change %s permissions. Please re-run kutu using sudo", filePath)
	}

	return err
}
