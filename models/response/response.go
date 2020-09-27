package response

import "github.com/bigbroproject/bigbrocore/models"

type ResponseType string

const (
	Success ResponseType = "Success"
	Error                = "Error"
)

type Response struct {
	ServiceName  string
	Protocol     models.Protocol
	ResponseType ResponseType
	Error        error
}
