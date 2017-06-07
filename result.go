package monitaur

import (
	"time"
)

type Result struct {
	Check    Check         `json:"-"`
	Executed time.Time     `json:"executed"`
	Duration time.Duration `json:"duration"`
	Status   Status        `json:"status"`
	Output   string        `json:"output"`
}
