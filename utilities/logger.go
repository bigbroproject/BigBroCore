package utilities

import (
	"fmt"
	"github.com/bigbrocore/models"
	"github.com/fatih/color"
	"log"
	"strconv"
	"time"
)

/*
PrintStatus prints the result of a CheckService result and returns the same error, if there was, otherwise nil will be returned.
*/
func PrintStatus(service *models.Service, protocol *models.Protocol, err error) (string, error) {

	now := time.Now()
	port := strconv.Itoa(protocol.Port)
	message := ""
	if port == "0" {
		port = "No port"
	}
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		message = fmt.Sprintf("[%s] [%s] [%s] [%s - %s - %s] An error as accured: %s", red("ERRO"), now.Format(time.RFC3339), service.Name, protocol.Type, protocol.Server, port, err.Error())
	} else {
		green := color.New(color.FgHiGreen).SprintFunc()
		message = fmt.Sprintf("[ %s ] [%s] [%s] [%s - %s - %s] Service seems OK.", green("OK"), now.Format(time.RFC3339), service.Name, protocol.Type, protocol.Server, port)
	}
	log.Println(message)
	return message, err
}

func printColorString(str string, clr color.Attribute) {
	color.New(color.FgHiBlue).SprintFunc()("Load")
	log.Printf("%s", color.New(clr).SprintFunc()(str))
}
