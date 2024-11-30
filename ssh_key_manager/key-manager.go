package ssh_key_manager

import (
	"fmt"
	"os"
	"os/exec"
	"pass-ssh/pass"
	"path/filepath"
	"strconv"
	"strings"
)

type KeyManager struct {
	TempDir  string
	KeyPath  string
	PassPath string
}

func NewKeyManager(path string) (*KeyManager, error) {
	tempDir, err := os.MkdirTemp("", "ssh-pass-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %v", err)
	}

	return &KeyManager{
		TempDir:  tempDir,
		KeyPath:  filepath.Join(tempDir, "id_rsa"),
		PassPath: path,
	}, nil
}

func (km *KeyManager) Cleanup() {
	os.RemoveAll(km.TempDir)
}

func (km *KeyManager) Setup() error {
	// Check if key exists
	if err := km.checkKey(); err != nil {
		return fmt.Errorf("key '%s' not found in pass", km.PassPath)
	}

	// Extract keys
	if err := km.extractKeys(); err != nil {
		return fmt.Errorf("failed to extract keys: %v", err)
	}

	return nil
}

func (km *KeyManager) checkKey() error {
	cmd := exec.Command("pass", "show", km.PassPath)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (km *KeyManager) extractKeys() error {
	// Extract private key
	if err := km.extractKey(km.PassPath, km.KeyPath, 0600); err != nil {
		return err
	}

	// Extract public key
	return km.extractKey(km.PassPath+".pub", km.KeyPath+".pub", 0644)
}

func (km *KeyManager) extractKey(passPath, outputPath string, perm os.FileMode) error {
	output, err := exec.Command("pass", "show", passPath).Output()
	if err != nil {
		return fmt.Errorf("failed to get key from pass: %v", err)
	}

	cleanOutput := strings.TrimSpace(string(output)) + "\n"
	if err := os.WriteFile(outputPath, []byte(cleanOutput), perm); err != nil {
		return fmt.Errorf("failed to write key: %v", err)
	}

	return nil
}

func (km *KeyManager) GenerateKeys() error {
	// Get encryption level from environment variable, default to 4096
	encryptionLevel := 4096
	if envLevel := os.Getenv("SSH_KEY_ENCRYPTION_LEVEL"); envLevel != "" {
		if parsedLevel, err := strconv.Atoi(envLevel); err == nil {
			encryptionLevel = parsedLevel
		} else {
			fmt.Printf("Invalid SSH_KEY_ENCRYPTION_LEVEL value '%s', falling back to default: 4096\n", envLevel)
		}
	}

	// Generate private key
	keyGenCmd := exec.Command("ssh-keygen", "-t", "rsa", "-b", strconv.Itoa(encryptionLevel), "-N", "", "-f", km.KeyPath)
	keyGenCmd.Stdout = os.Stdout
	keyGenCmd.Stderr = os.Stderr
	if err := keyGenCmd.Run(); err != nil {
		return fmt.Errorf("failed to generate SSH key: %v", err)
	}

	// Read the generated private key
	privateKey, err := os.ReadFile(km.KeyPath)
	if err != nil {
		return fmt.Errorf("failed to read private key: %v", err)
	}

	// Store private key in pass
	if err := pass.StoreInPass(km.PassPath, privateKey); err != nil {
		return err
	}

	// Read the generated public key
	publicKey, err := os.ReadFile(km.KeyPath + ".pub")
	if err != nil {
		return fmt.Errorf("failed to read public key: %v", err)
	}

	// Store public key in pass
	if err := pass.StoreInPass(km.PassPath+".pub", publicKey); err != nil {
		return err
	}

	fmt.Printf("SSH key pair generated with %d-bit encryption and stored in pass as '%s' and '%s.pub'\n", encryptionLevel, km.PassPath, km.PassPath)
	return nil
}
