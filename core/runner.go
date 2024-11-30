package core

import (
	"os"
	"os/exec"
)

// Interface for running commands
type Runner interface {
	Run() error
}

// RunCommand runs a command with the given arguments and environment variables
func RunCommand(command string, args []string, env []string) error {
	cmd := exec.Command(command, args...)
	if env != nil {
		cmd.Env = env
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
