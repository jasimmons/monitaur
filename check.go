package monitaur

import (
	"encoding/json"
	"github.com/jasimmons/monitaur/log"
	"time"
)

var (
	DefaultCheckInterval = 60 * time.Second
	DefaultCheckRunner   = DefaultShellRunner
)

type Check interface {
	Execute() Result
	Stop()
	GetName() string
	GetInterval() time.Duration
	GetHandlers() []string
}

type DefaultCheck struct {
	Type       string        `json:"type"`
	Name       string        `json:"name"`
	Command    string        `json:"command"`
	Interval   time.Duration `json:"interval"`
	Handlers   []string      `json:"handlers"`
	RunnerHint string        `json:"runner,omitempty"`
	runner     Runner
}

func ParseCheckWithDefaults(checkJson []byte) (Check, error) {
	var check DefaultCheck
	err := json.Unmarshal(checkJson, &check)
	if err != nil {
		return nil, err
	}

	if check.Interval == 0 {
		check.Interval = DefaultCheckInterval
	}
	check.runner = ParseRunner(check.RunnerHint)

	return &check, nil
}

func (c *DefaultCheck) Execute() Result {
	startTime := time.Now()
	code, stdout, _ := c.runner.Run(c.Command)
	endTime := time.Now()

	return Result{
		Check:    c,
		Executed: startTime,
		Duration: endTime.Sub(startTime),
		Status:   Status(code),
		Output:   stdout,
	}
}

func (c *DefaultCheck) Stop() {
	log.Infof("stopping check %s\n", c.Name)
}

func (c *DefaultCheck) GetName() string {
	return c.Name
}

func (c *DefaultCheck) GetInterval() time.Duration {
	return c.Interval
}

func (c *DefaultCheck) GetHandlers() []string {
	return c.Handlers
}

func (c *DefaultCheck) SetRunner(runner Runner) {
	c.runner = runner
}
