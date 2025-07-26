package collector

import (
	"os"
	"os/user"
)

func (c *SystemCollector) getUser() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return currentUser.Username, nil
}

func (c *SystemCollector) getHostname() (string, error) {
	return os.Hostname()
}

func (c *SystemCollector) getShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "Unknown"
	}

	for i := len(shell) - 1; i >= 0; i-- {
		if shell[i] == '/' {
			return shell[i+1:]
		}
	}

	return shell
}
