package commands

import "github.com/tylergets/workon/logger"

type Command interface {
	Execute(*logger.Logger, map[string]string) error
}
