package main

import (
	"github.com/moneye/protocols"
	"github.com/moneye/responsehandlers"
)

func main() {

	regProtocolInterfaces, regResponseHandlerInterfaces := Initialize()

	// Register custom protocol
	protocols.RegisterProtocolInterface(&regProtocolInterfaces,"ftp", protocols.FTP{})

	// Register custom Response Handler
	responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces,"console",responsehandlers.ConsoleHandlerWithMemory{})
	Start(regProtocolInterfaces,regResponseHandlerInterfaces)

}


