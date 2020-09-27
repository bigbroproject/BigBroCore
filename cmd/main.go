package main

import (
	"github.com/bigbroproject/bigbrocore/protocols"
	"github.com/bigbroproject/bigbrocore/responsehandlers"
	"os"
)

func main() {

	/*
		for debug purposes
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	*/

	regProtocolInterfaces, regResponseHandlerInterfaces := Initialize(os.Args[1])

	// Register custom protocol
	protocols.RegisterProtocolInterface(&regProtocolInterfaces, "ftp", protocols.FTP{})

	// Register custom Response Handler
	responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "consoleWithMemory", responsehandlers.ConsoleHandlerWithMemory{})
	responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "console", responsehandlers.ConsoleHandler{})

	// Start monitoring
	Start(regProtocolInterfaces, regResponseHandlerInterfaces)

}
