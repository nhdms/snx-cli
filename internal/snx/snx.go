package snx

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Kill() {
	// Kill any existing snx process to ensure clean connection
	exec.Command("snx", "-d").Run()
	exec.Command("killall", "snx").Run()
	time.Sleep(500 * time.Millisecond)
}

func Connect(server, username, password string) error {
	Kill()

	cmd := exec.Command("snx", "-s", server, "-u", username)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdin pipe: %w", err)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start snx: %w", err)
	}

	// Give snx time to prompt for password
	time.Sleep(500 * time.Millisecond)

	_, err = io.WriteString(stdin, password+"\n")
	if err != nil {
		return fmt.Errorf("failed to write password: %w", err)
	}
	stdin.Close()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("snx exited with error: %w", err)
	}

	return nil
}

func Disconnect() error {
	cmd := exec.Command("snx", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to disconnect: %w", err)
	}
	return nil
}

func Status() (bool, error) {
	out, err := exec.Command("ip", "addr", "show", "tunsnx").CombinedOutput()
	if err != nil {
		return false, nil
	}
	return strings.Contains(string(out), "tunsnx"), nil
}
