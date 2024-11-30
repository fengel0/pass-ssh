package core

import (
	"fmt"
)

type CommandType int

const (
	SSH CommandType = iota
	Git
	Create
	Copy
)

type Command struct {
	Type    CommandType
	KeyName string
	KeyPath string
	Args    []string
}

func get_git_key_path(key_name string) string {
	return fmt.Sprintf("ssh-keys/git/%s", key_name)
}

func get_ssh_key_path(key_name string) string {
	return fmt.Sprintf("ssh-keys/server/%s", key_name)
}

func ParseCommand(args []string) (Command, error) {
	if len(args) < 2 {
		return Command{}, fmt.Errorf("insufficient arguments")
	}

	switch args[1] {
	case "git":
		if len(args) < 3 {
			return Command{}, fmt.Errorf("insufficient arguments for git command")
		}
		return Command{
			Type:    Git,
			KeyName: get_git_key_path(args[0]),
			Args:    args[2:],
		}, nil
	case "ssh":
		if len(args) < 3 {
			return Command{}, fmt.Errorf("insufficient arguments for ssh command")
		}
		return Command{
			Type:    SSH,
			KeyName: get_ssh_key_path(args[0]),
			Args:    args[2:],
		}, nil
	case "create":
		key_name := ""
		switch args[0] {
		case "git":
			key_name = get_git_key_path(args[2])
		case "ssh":
			key_name = get_ssh_key_path(args[2])
		default:
			return Command{}, fmt.Errorf("unknown command type '%s'", args[0])
		}
		return Command{
			Type:    Create,
			KeyName: key_name,
			Args:    args[2:],
		}, nil
	case "copy":
		if len(args) < 3 {
			return Command{}, fmt.Errorf("insufficient arguments for copy command")
		}
		return Command{
			Type:    Copy,
			KeyName: get_ssh_key_path(args[0]),
			Args:    args[2:],
		}, nil
	default:
		return Command{}, fmt.Errorf("unknown command type '%s'", args[1])
	}
}
