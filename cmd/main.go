package main

import (
	protocols "github.com/bigbrocore/protocols"
	"github.com/bigbrocore/responsehandlers"
)

func main() {

	regProtocolInterfaces, regResponseHandlerInterfaces := Initialize()

	// Register custom protocol
	protocols.RegisterProtocolInterface(&regProtocolInterfaces, "ftp", protocols.FTP{})

	// Register custom Response Handler
	responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "consoleWithMemory", responsehandlers.ConsoleHandlerWithMemory{})
	responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "console", responsehandlers.ConsoleHandler{})

	// Start monitoring
	Start(regProtocolInterfaces, regResponseHandlerInterfaces)

}
