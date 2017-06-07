package monitaur

import (
	"encoding/json"
	"errors"
	"time"
)

type Handler interface {
	Handle(event Event) (output string, err error)
}

type DefaultHandler struct {
	Type       string    `json:"type"`
	Name       string    `json:"name"`
	Command    string    `json:"command"`
	RunnerHint string    `json:"runner,omitempty"`
	Timestamp  time.Time `json:"timestamp"`
	Output     string    `json:"output"`
	runner     Runner
	Event      Event
}

func (h *DefaultHandler) Handle(event Event) (string, error) {
	h.Event = event
	h.Timestamp = time.Now()
	_, stdout, stderr := h.runner.Run(h.Command)
	if stderr != "" {
		return stdout, errors.New(stderr)
	}
	return stdout, nil
}

func ParseHandlerWithDefaults(handlerJson []byte) (Handler, error) {
	var handler *DefaultHandler
	err := json.Unmarshal(handlerJson, &handler)
	return handler, err
}
