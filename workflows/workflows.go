package workflows

import (
	"github.com/tylergets/workon/commands"
	"github.com/tylergets/workon/config"
	"github.com/tylergets/workon/logger"
	"github.com/tylergets/workon/terminal"
	"os/exec"
)

func ExecuteWorkflow(name string, cfg config.Config) {
	workflow, ok := cfg.Workflows[name]
	if !ok {
		logger.NewLogger(0).Error("Workflow '%s' not found.", name)
		return
	}

	for i, step := range workflow {
		log := logger.NewLogger(i)
		stepName := step.Name

		if step.Run != nil {
			for _, cmdStr := range step.Run {
				if step.Windowed {
					// Launch Command In A New Kitty Terminal
					terminal.Launch(cmdStr, step.Cwd, log)
				} else {
					if err := ExecuteCommand(cmdStr, step.Cwd, log); err != nil {
						return
					}
				}
			}
		} else {
			log.Info("Running predefined command: %s", stepName)
			predefinedCommand := commands.Get(stepName)
			if predefinedCommand != nil {
				err := predefinedCommand.Execute(log, step.Data)
				if err != nil {
					return
				}
			} else {
				log.Error("No predefined command found for: %s", stepName)
			}
		}
	}
}

func ExecuteCommand(cmdStr string, cwd *string, log *logger.Logger) error {
	log.Info("Running command: %s", cmdStr)

	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	if cwd != nil {
		log.Info("Setting working directory to: %s", *cwd)
		cmd.Dir = *cwd
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("Command '%s' failed: %s", cmdStr, err.Error())
		log.Error("Command output: %s", string(out))
		return err
	}
	log.Info("Command output: %s", string(out))
	return nil
}
