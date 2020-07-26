package testapp

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func Build(filepath string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "hacktherank-*")
	if err != nil {
		return "", fmt.Errorf("failed to create temp directory: %w", err)
	}
	executableFilepath := fmt.Sprintf("%s/executable", tmpDir)

	cmd := exec.Command("go", "build", "-o", executableFilepath, filepath)
	if _, err := cmd.Output(); err != nil {
		switch actualErr := err.(type) {
		case *exec.ExitError:
			return "", fmt.Errorf("failed to build go executable: %s", string(actualErr.Stderr))
		default:
			return "", fmt.Errorf("failed to build go executable: %w", err)
		}
	}

	return executableFilepath, nil
}
