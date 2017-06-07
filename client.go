package monitaur

import (
	"github.com/jasimmons/monitaur/log"
	"os"
	"os/signal"
	"time"
)

type Client struct {
	Name     string      `json:"name"`
	Addr     string      `json:"address"`
	cfg      *Config     `json:"-"`
	Checks   []Check     `json:"-"`
	Handlers []Handler   `json:"-"`
	Results  chan Result `json:"-"`
	Errors   chan error  `json:"-"`
	Cache    *EventCache `json:"-"`
}

func NewClient(cfg *Config, checks []Check) *Client {
	return &Client{
		cfg:     cfg,
		Checks:  checks,
		Results: make(chan Result),
		Errors:  make(chan error),
		Cache:   NewEventCache(),
	}
}

func (client *Client) RunChecks() {
	for _, check := range client.Checks {
		//log.Debugf("starting check %s on interval %d\n", check.Name, check.Interval)
		go func(cli *Client, ch Check) {
			for range time.Tick(ch.GetInterval() * time.Second) {
				cli.Results <- ch.Execute()
			}
		}(client, check)
	}
}

func (client *Client) HandleResults() {
	go func() {
		for result := range client.Results {
			log.Infof("%+v\n", result)
			if hist, ok := client.Cache.Get(client, result.Check); ok {
				// There is a history of events with this ID already

				lastIndex := len(hist) - 1
				if hist[lastIndex].Status != result.Status {
					//log.Debugf("check %s on client %s changed status from %d to %d!\n", result.Check.Name, client.Name, hist[lastIndex].Status, result.Status)
					for _ = range result.Check.GetHandlers() {

					}
				}
				newEvent := GenerateEvent(result, client)
				client.Cache.Save(newEvent)
			} else {
				// No event exists with this ID yet
				//log.Debugf("new event created for check %s on client %s\n", result.Check.Name, client.Name)
				newEvent := GenerateEvent(result, client)
				client.Cache.Save(newEvent)
			}
		}
	}()
}

func (client *Client) Run() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill, os.Interrupt)

	for {
		client.RunChecks()
		client.HandleResults()

		select {
		case s := <-sigChan:
			log.Infof("received signal %s\n", s.String())
			for _, check := range client.Checks {
				check.Stop()
			}
			return
		}
	}
}
