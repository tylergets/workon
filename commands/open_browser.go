package commands

import (
	"errors"
	"github.com/tylergets/workon/logger"
	"os"
	"os/exec"
)

type OpenBrowser struct{}

func (c *OpenBrowser) Execute(log *logger.Logger, data map[string]string) error {
	url := data["url"]
	if url == "" {
		log.Error("No url specified.")
		return errors.New("no url specified")
	}
	log.Info("Opening browser to: %s", url)

	browser := os.Getenv("BROWSER")
	if browser == "" {
		log.Error("Environment variable BROWSER is not set.")
		return errors.New("environment variable BROWSER is not set")
	}

	cmd := exec.Command(browser, url)
	if err := cmd.Start(); err != nil {
		log.Error("Failed to open browser: %s", err)
		return err
	}

	return nil
}

func init() {
	Register("open-browser", &OpenBrowser{})
}
