package responsehandlers

import (
	"github.com/bigbroproject/bigbrocore/models"
	"github.com/bigbroproject/bigbrocore/models/response"
)

type ResponseHandlerInterface interface {

	// Handle - manage the responses.
	// configuration - is the same configuration model loaded from configFile.
	// channel - is the channel where responses arrive from check processes.
	Handle(*models.Config, *chan response.Response)
}

func RegisterResponseHandlerInterface(registeredResponseHandlerInterfaces *map[string]ResponseHandlerInterface, responseInterfaceName string, responseHandlerInterface ResponseHandlerInterface) {
	(*registeredResponseHandlerInterfaces)[responseInterfaceName] = responseHandlerInterface
}

func DefaultRegisteredResponseHandlers() map[string]ResponseHandlerInterface {
	registeredHandlers := make(map[string]ResponseHandlerInterface, 0)
	//registeredHandlers["console"] = ConsoleHandler{}
	return registeredHandlers
}
