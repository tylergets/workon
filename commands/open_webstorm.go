package commands

import (
	"errors"
	"github.com/tylergets/workon/logger"
	"os/exec"
)

type OpenWebstorm struct{}

func (c *OpenWebstorm) Execute(log *logger.Logger, data map[string]string) error {
	project := data["project"]
	if project == "" {
		log.Error("No project specified.")
		return errors.New("no project specified")
	}
	log.Info("Opening webstorm to: %s", project)

	// Execute the WebStorm command with the project path.
	cmd := exec.Command("webstorm", project)
	err := cmd.Start()
	if err != nil {
		log.Error("Failed to open WebStorm: %s", err)
		return err
	}

	log.Info("WebStorm opened successfully")
	return nil
}

func init() {
	Register("open-webstorm", &OpenWebstorm{})
}
