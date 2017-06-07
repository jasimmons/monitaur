package monitaur

import (
	"time"
)

type Event struct {
	Client   *Client       `json:"client"`
	Check    Check         `json:"check"`
	Executed time.Time     `json:"executed"`
	Duration time.Duration `json:"duration"`
	Status   Status        `json:"status"`
	History  []Status      `json:"history"`
	Output   string        `json:"output"`
}

type History []Event

func GenerateEvent(res Result, cli *Client) Event {
	return Event{
		Client:   cli,
		Check:    res.Check,
		Executed: res.Executed,
		Duration: res.Duration,
		Status:   res.Status,
		Output:   res.Output,
		History:  make([]Status, 0, 1),
	}
}
