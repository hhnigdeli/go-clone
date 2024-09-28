package getGitCredentials

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Get(repoURL string) (string, string, error) {
	cmd := exec.Command("git", "credential", "fill")

	// Prepare the input for the git credential fill command
	input := fmt.Sprintf("url=%s\n\n", repoURL)
	cmd.Stdin = strings.NewReader(input)

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", "", fmt.Errorf("failed to get git credentials: %w", err)
	}

	// Parse the output to extract the username and password
	output := out.String()
	lines := strings.Split(output, "\n")
	var username, password string
	for _, line := range lines {
		if strings.HasPrefix(line, "username=") {
			username = strings.TrimPrefix(line, "username=")
		} else if strings.HasPrefix(line, "password=") {
			password = strings.TrimPrefix(line, "password=")
		}
	}

	if username == "" || password == "" {
		return "", "", fmt.Errorf("failed to parse git credentials")
	}

	return username, password, nil
}
