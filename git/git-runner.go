package git

import (
	"fmt"
	"os"
	"pass-ssh/core"
)

type GitRunner struct {
	KeyPath string
	Args    []string
}

func (r GitRunner) Run() error {
	env := append(os.Environ(),
		fmt.Sprintf("GIT_SSH_COMMAND=ssh -i %s -o IdentitiesOnly=yes", r.KeyPath),
	)
	return core.RunCommand("git", r.Args, env)
}
