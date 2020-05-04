package connectors

import (
	"fmt"

	"github.com/microlib/simple"
)

// Connections struct - all backend connections in a common object
type Connections struct {
	l *simple.Logger
}

func (r *Connections) Error(msg string, val ...interface{}) {
	r.l.Error(fmt.Sprintf(msg, val...))
}

func (r *Connections) Info(msg string, val ...interface{}) {
	r.l.Info(fmt.Sprintf(msg, val...))
}

func (r *Connections) Debug(msg string, val ...interface{}) {
	r.l.Debug(fmt.Sprintf(msg, val...))
}

func (r *Connections) Trace(msg string, val ...interface{}) {
	r.l.Trace(fmt.Sprintf(msg, val...))
}

// NewClientConnectors returns Connectors struct
func NewClientConnections(logger *simple.Logger) Clients {
	conns := &Connections{l: logger}
	logger.Debug(fmt.Sprintf("Connection details %v\n", conns))
	return conns
}
