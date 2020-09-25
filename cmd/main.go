package main

import (
	"github.com/moneye/protocols"
	"github.com/moneye/responsehandlers"
)

func main() {

	regProtocolInterfaces, regResponseHandlerInterfaces := Initialize()

	// Register custom protocol
	protocols.RegisterProtocolInterface(&regProtocolInterfaces, "ftp", protocols.FTP{})

	// Register custom Response Handler
	responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "output1", responsehandlers.ConsoleHandlerWithMemory{})
	//responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "output2", responsehandlers.ConsoleHandler{})

	// Start monitoring
	Start(regProtocolInterfaces, regResponseHandlerInterfaces)

}
