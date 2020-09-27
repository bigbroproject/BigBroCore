# BigBro Core

An extensible monitoring tool for user defined services and protocols, all wrote in GoLang.

## Introduction

BigBro Core is the core library of the main project BigBro, it can be used in standalone way (without any web service or user interface, recommend for low spec devices (as OrangePi or others).

## Defaults Protocols supported
- http
- https (with ssl)
- icmp
- icmp6

## Configuration file

## Struct of the Core and implementation of custom protocols and handlers

parlare della registrazione di protocolli e handlers

BigBro core structured in two main parts: `protocolInterface` and `responseHandlerInterface`. 
### protocolInterface
BigBro Core allows the implementation of new type of network protocol or services to be monitor. To develop an implementation the developer must implements the `protocolInterface`, in particular the developer must implement the `CheckService` function, and the struct that defines the protocol or service.

In the `CheckService` function must be defined the check login of the specified protocol / service and then return the `error` (if occurred) or `nil`. 

An example of implementation of a http protocol:

```go
package protocols

import (
	"github.com/bigbroproject/bigbrocore/models"
	httplib "net/http"
	"strconv"
	"time"
)

type Http struct {
}

// If error is nil, then service is up
func (http Http) CheckService(Protocol models.Protocol) error {
	// CHECK
	tr := &httplib.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	client := &httplib.Client{Transport: tr}
	url := "http://" + Protocol.Server
	if Protocol.Port == 0 {
		url += ":80"
	} else {
		url += ":" + strconv.Itoa(Protocol.Port)
	}
	_, err := client.Get(url)
	return err
}
```

In this example the `CheckService` try to perform a simply GET via http to the specific server or service and then return the http error if occurred, `nil` otherwise.


### responseHandlerInterface
A responseHandlers receives the responses (the responses are define by the `models/response/Response` struct) from the monitoring process, defined by the config files and registered protocols. 

A struct that defines the handler and a `Handle` function must be implemented to mange the received `Response` from the monitoring processes. It must had an infinite loop to manage the receiving Responses.

A response is structured as follow:
```go
type ResponseType string

const (
	Success ResponseType = "Success"
	Error                = "Error"
)

type Response struct {
	ServiceName  string          // The service name defined in configuration file
	Protocol     models.Protocol // The protocol used to monitoring defined in configuration file
	ResponseType ResponseType    // The ReponseType (Success or Error)
	Error error                  // The error if there is any
}
```

The channel which is passed as input in `Handle` function is used where the responses arrive.

Following, an example of `ConsoleHandler`, that simply prints the `Responses` in console:

```go
package responsehandlers

import (
	"fmt"
	"github.com/bigbroproject/bigbrocore/models/response"
	"github.com/fatih/color"
	"log"
	"strconv"
	"time"
)

type ConsoleHandler struct {}


func ( handler ConsoleHandler) Handle(channel *chan response.Response){
    for  { //the infinite loop
	resp := <- *channel // the receiving response
	if resp.ResponseType == response.Error {
            message = fmt.Sprintf("[%s] [%s] [%s] [%s - %s - %s] An error as accured: %s", "ERRO", now.Format(time.RFC3339), resp.ServiceName, resp.Protocol.Type, resp.Protocol.Server, port, resp.Error.Error())
        } else {
            message = fmt.Sprintf("[ %s ] [%s] [%s] [%s - %s - %s] Service seems OK.", "OK", now.Format(time.RFC3339), resp.ServiceName, resp.Protocol.Type, resp.Protocol.Server, port)
        }
        log.Println(message)
    }
}
```




