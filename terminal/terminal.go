package terminal

import (
	"fmt"
	"github.com/tylergets/workon/logger"
	"os/exec"
	"syscall"
)

// Launch opens a Kitty terminal in the specified directory and runs the specified command.
func Launch(command string, cwd *string, log *logger.Logger) error {
	// Construct the command to launch Kitty with the specified working directory and command.

	//cmd := exec.Command("kitty", "-d", cwd, "sh", "-c", command)
	// cwd may be null
	cmd := exec.Command("kitty", "sh", "-c", command)
	if cwd != nil {
		log.Info("Setting working directory to: %s", *cwd)
		cmd = exec.Command("kitty", "-d", *cwd, "sh", "-c", command)
	} else {
		log.Info("No working directory specified, using: %s", "~")
		cmd = exec.Command("kitty", "sh", "-c", command)
	}

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	// Run the command and return any errors.
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to launch Kitty: %w", err)
	}

	return nil
}
