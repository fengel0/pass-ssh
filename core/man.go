package core

import (
	"fmt"
	"os"
)

func PrintUsage() {
	fmt.Printf(`Usage:
  SSH:    %s <key_name> ssh     <user@host>   [ssh_options]
  Git:    %s <key_name> git     <git_command> [git_args]
  Create: %s <git/ssh>  create  <key_name>
  Copy:   %s <key_name> copy    <user@host>

Examples:
  SSH:    %s work_laptop ssh user@example.com
  Git:    %s github_key git clone git@github.com:user/repo.git
          %s github_key git push origin main
  Create: %s new_key_name create
  Copy:   %s work_laptop copy user@example.com
Environment:
  SSH_KEY_ENCRYPTION_LEVEL - Key encryption level (default: 4096)
`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}
