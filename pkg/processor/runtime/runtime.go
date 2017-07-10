package runtime

import (
	"github.com/nuclio/nuclio/pkg/logger"
	"github.com/nuclio/nuclio/pkg/processor/event"
)

type Runtime interface {
	ProcessEvent(event event.Event) (interface{}, error)
}

type AbstractRuntime struct {
	Logger  logger.Logger
	Context *Context
}

func NewAbstractRuntime(logger logger.Logger, configuration *Configuration) *AbstractRuntime {
	return &AbstractRuntime{
		Logger:  logger,
		Context: newContext(logger, configuration),
	}
}