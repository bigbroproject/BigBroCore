package responsehandlers

import "github.com/bigbroproject/bigbrocore/models/response"

type ResponseHandlerInterface interface {
	Handle(*chan response.Response)
}

func RegisterResponseHandlerInterface(registeredResponseHandlerInterfaces *map[string]ResponseHandlerInterface, responseInterfaceName string, responseHandlerInterface ResponseHandlerInterface) {
	(*registeredResponseHandlerInterfaces)[responseInterfaceName] = responseHandlerInterface
}

func DefaultRegisteredResponseHandlers() map[string]ResponseHandlerInterface {
	registeredHandlers := make(map[string]ResponseHandlerInterface, 0)
	//registeredHandlers["console"] = ConsoleHandler{}
	return registeredHandlers
}
