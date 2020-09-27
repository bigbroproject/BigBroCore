package main

import (
	"fmt"
	"github.com/bigbroproject/bigbrocore/models"
	"github.com/bigbroproject/bigbrocore/models/response"
	"github.com/bigbroproject/bigbrocore/process"
	"github.com/bigbroproject/bigbrocore/protocols"
	"github.com/bigbroproject/bigbrocore/responsehandlers"
	"github.com/bigbroproject/bigbrocore/utilities"
	"github.com/fatih/color"
	"log"
	"time"
)

func Initialize() (map[string]protocols.ProtocolInterface, map[string]responsehandlers.ResponseHandlerInterface) {
	return protocols.DefaultRegisteredProtocolInterfaces(), responsehandlers.DefaultRegisteredResponseHandlers()
}

func Start(registeredProtocolInterfaces map[string]protocols.ProtocolInterface, registeredResponseHandlerInterfaces map[string]responsehandlers.ResponseHandlerInterface) {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf, err := models.ConfigFromFile("config/conf.yml")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[%s] Configuration loaded!", utilities.CreateColorString("Load", color.FgHiBlue))
	processesChannel := make(chan string)
	responseChannel := make(chan response.Response)

	startResponseBroadcaster(&responseChannel, &registeredResponseHandlerInterfaces)
	log.Printf("[%s] Starting Handlers complete!", utilities.CreateColorString("Complete", color.FgHiBlue))

	if err != nil {
		log.Fatal(err.Error())
	}
	for i := range conf.Services {
		service := conf.Services[i]
		for i2 := range service.Protocols {
			protocol := service.Protocols[i2]
			protocolInterface := registeredProtocolInterfaces[protocol.Type]
			if protocols.IsRegistered(&registeredProtocolInterfaces, protocol.Type) {
				proc := process.NewProcess(func() {
					if err = protocolInterface.CheckService(protocol); err == nil { // SUCCESS RESP
						responseChannel <- response.Response{
							ServiceName:  service.Name,
							Protocol:     protocol,
							ResponseType: response.Success,
							Error:        nil}
					} else { // ERROR RESP
						responseChannel <- response.Response{
							ServiceName:  service.Name,
							Protocol:     protocol,
							ResponseType: response.Error,
							Error:        err}
					}

				}, processesChannel)
				process.ScheduleProcess(proc, protocol.Interval)

			} else {
				red := color.New(color.FgRed).SprintFunc()
				log.Println(fmt.Sprintf("[%s] [%s] [%s] [%s - %s - %s] An error as accured: %s", red("ERRO"), time.Now().Format(time.RFC3339), service.Name, protocol.Type, protocol.Server, protocol.Port, "ProtocolInterface "+protocol.Type+" not registered!"))
			}

		}
	}

	for {
		time.Sleep(100 * time.Millisecond)
	}
}

func startResponseBroadcaster(responseChannel *chan response.Response, responseHandlers *map[string]responsehandlers.ResponseHandlerInterface) {
	chanArray := make([]*chan response.Response, 0)
	// Create channels and start handlers
	for _, handler := range *responseHandlers {
		channel := make(chan response.Response)
		chanArray = append(chanArray, &channel)

		// Start handler on created channel
		handler := handler
		go func() {
			handler.Handle(&channel)
		}()
	}

	// Start endless loop that read from responsChannel and publish on handlers channels
	go func() {
		for {
			resp := <-*responseChannel
			for _, channel := range chanArray {
				*channel <- resp
			}
		}
	}()
}
