package main

import (
	protocols "github.com/bigbroproject/bigbrocore/protocols"
	"github.com/bigbroproject/bigbrocore/responsehandlers"
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
