package main

import (
	"github.com/moneye/protocols"
	"github.com/moneye/responsehandlers"
)

func main() {

	regProtocolInterfaces, regResponseHandlerInterfaces := Initialize()
	protocols.RegisterProtocolInterface(&regProtocolInterfaces,"ftp", protocols.FTP{})
	responsehandlers.RegisterHandlerInterface(&regResponseHandlerInterfaces,"console",responsehandlers.ConsoleHandler{})
	Start(regProtocolInterfaces,regResponseHandlerInterfaces)

}


