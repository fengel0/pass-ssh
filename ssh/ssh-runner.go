package ssh

import (
	"pass-ssh/core"
)

type SSHRunner struct {
	KeyPath string
	Args    []string
}

func (r SSHRunner) Run() error {
	args := append([]string{"-i", r.KeyPath}, r.Args...)
	return core.RunCommand("ssh", args, nil)
}

type CopyRunner struct {
	PublicKeyPath string
	Host          string
}

func (r CopyRunner) Run() error {
	return core.RunCommand("ssh-copy-id", []string{"-i", r.PublicKeyPath, r.Host}, nil)
}
