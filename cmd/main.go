package main

import (
	"github.com/bigbroproject/bigbrocore/core"
	"github.com/bigbroproject/bigbrocore/protocols"
	"github.com/bigbroproject/bigbrocore/responsehandlers"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {

	/*
		for debug purposes
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	*/

	regProtocolInterfaces, regResponseHandlerInterfaces := core.Initialize(os.Args[1])

	// Register custom protocol
	protocols.RegisterProtocolInterface(&regProtocolInterfaces, "ftp", protocols.FTP{})

	// Register custom Response Handler
	responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "consoleWithMemory", responsehandlers.ConsoleHandlerWithMemory{})
	//responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "console", responsehandlers.ConsoleHandler{})

	// Only for debugging
	go func() {
		for  {
			log.Println("ACTIVE GOROUTINES :"+strconv.Itoa(runtime.NumGoroutine()))
			time.Sleep(time.Second * 1)
		}
	}()
	// Start monitoring
	core.Start(regProtocolInterfaces, regResponseHandlerInterfaces)

}
