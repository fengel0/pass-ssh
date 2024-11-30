package main

import (
	"log"
	"os"
	"pass-ssh/core"
	"pass-ssh/git"
	"pass-ssh/ssh"
	"pass-ssh/ssh_key_manager"
)

func main() {
	if len(os.Args) < 2 {
		core.PrintUsage()
		os.Exit(1)
	}

	cmd, err := core.ParseCommand(os.Args[1:])
	if err != nil {
		log.Fatalf("Error parsing command: %v", err)
	}

	km, err := ssh_key_manager.NewKeyManager(cmd.KeyName)
	if err != nil {
		log.Fatalf("Error initializing key manager: %v", err)
	}
	defer km.Cleanup()

	switch cmd.Type {
	case core.SSH, core.Git:
		if err := km.Setup(); err != nil {
			log.Fatalf("Error setting up keys: %v", err)
		}

		var runner core.Runner
		switch cmd.Type {
		case core.SSH:
			runner = ssh.SSHRunner{KeyPath: km.KeyPath, Args: cmd.Args}
		case core.Git:
			runner = git.GitRunner{KeyPath: km.KeyPath, Args: cmd.Args}
		}

		if err := runner.Run(); err != nil {
			log.Fatalf("Command failed: %v", err)
		}
	case core.Create:
		if err := km.GenerateKeys(); err != nil {
			log.Fatalf("Error generating keys: %v", err)
		}
	case core.Copy:
		if err := km.Setup(); err != nil {
			log.Fatalf("Error setting up keys: %v", err)
		}
		runner := ssh.CopyRunner{PublicKeyPath: km.KeyPath + ".pub", Host: cmd.Args[0]}
		if err := runner.Run(); err != nil {
			log.Fatalf("Error copying key: %v", err)
		}
	}
}
