package monitaur

import (
	"encoding/json"
)

func MarshalEvent(event Event) []byte {
	evt, _ := json.Marshal(event)
	return evt
}

func FormatCheckRequest(client *Client, check Check) ([]byte, error) {
	obj := struct {
		Client *Client `json:"client"`
		Check  Check   `json:"check"`
	}{
		Client: client,
		Check:  check,
	}
	return json.Marshal(obj)
}
