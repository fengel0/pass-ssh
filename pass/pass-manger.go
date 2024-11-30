package pass

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// stores the key in pass
func StoreInPass(path string, content []byte) error {
	cmd := exec.Command("pass", "insert", "--multiline", path)
	cmd.Stdin = strings.NewReader(string(content))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to store key in pass: %v", err)
	}
	return nil
}
