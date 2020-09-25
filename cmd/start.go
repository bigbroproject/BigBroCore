package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/moneye/models"
	"github.com/moneye/models/response"
	"github.com/moneye/process"
	"github.com/moneye/protocols"
	"github.com/moneye/responsehandlers"
	"log"
	"time"
)

func Initialize() (map[string]protocols.ProtocolInterface, map[string]responsehandlers.ResponseHandlerInterface) {
	return protocols.DefaultRegisteredProtocolInterfaces(), responsehandlers.DefaultRegisteredResponseHandlers()
}

func Start(registeredProtocolInterfaces map[string]protocols.ProtocolInterface, registeredResponseHandlerInterfaces map[string]responsehandlers.ResponseHandlerInterface) {

	conf, err := models.ConfigFromFile("config/conf.yml")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[%s] Configuration loaded!", color.New(color.FgHiBlue).SprintFunc()("Load"))
	processesChannel := make(chan string)
	responseChannel := make(chan response.Response)

	//startResponseHandlers(&responseChannel,registeredResponseHandlerInterfaces)
	startResponseBroadcaster(&responseChannel, registeredResponseHandlerInterfaces)

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

func startResponseBroadcaster(c *chan response.Response, interfaces map[string]responsehandlers.ResponseHandlerInterface) {
	go func() {
		for {
			resp := <-*c
			for _, itInterface := range interfaces {
				*itInterface.GetReceiveChannel() <- resp
			}
		}
	}()

}

/*
func startResponseHandlers(responseChannel *chan response.Response, registeredResponseHandlerInterfaces map[string]responsehandlers.ResponseHandlerInterface) {
	//registeredResponseHandlers := registeredResponseHandlerInterfaces
	for s := range registeredResponseHandlerInterfaces {
		go func() {
			registeredResponseHandlerInterfaces[s].Handle(responseChannel)
		}()
	}
	log.Printf("Response Handlers registered!")
}
*/
