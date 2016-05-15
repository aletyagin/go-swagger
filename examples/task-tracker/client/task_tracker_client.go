package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-swagger/go-swagger/examples/task-tracker/client/tasks"
)

// Default task tracker HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new task tracker HTTP client.
func NewHTTPClient(formats strfmt.Registry) *TaskTracker {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("localhost:8322", "/api", []string{"https", "http"})
	return New(transport, formats)
}

// New creates a new task tracker client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *TaskTracker {
	cli := new(TaskTracker)
	cli.Transport = transport

	cli.Tasks = tasks.New(transport, formats)

	return cli
}

// TaskTracker is a client for task tracker
type TaskTracker struct {
	Tasks *tasks.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *TaskTracker) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Tasks.SetTransport(transport)

}
