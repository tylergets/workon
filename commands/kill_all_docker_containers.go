package commands

import (
	"context"
	"github.com/tylergets/workon/logger"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type KillAllDockerContainers struct{}

func (c *KillAllDockerContainers) Execute(log *logger.Logger, m map[string]string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Error("Error creating Docker client: %s", err)
		return err
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Error("Error listing containers: %s", err)
		return err
	}

	for _, container := range containers {
		err := cli.ContainerKill(context.Background(), container.ID, "SIGKILL")
		if err != nil {
			log.Error("Failed to kill container %s: %s", container.ID[:10], err)
		} else {
			log.Info("Killed container %s", container.ID[:10])
		}
	}

	log.Success("All Docker containers killed.")
	return nil
}

func init() {
	Register("kill-all-docker-containers", &KillAllDockerContainers{})
}
