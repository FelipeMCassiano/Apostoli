package pkg

import (
	"os/exec"
)

func CloneRepo(url, path string) (string, error) {
	out, err := exec.Command("git", "clone", url, path).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
