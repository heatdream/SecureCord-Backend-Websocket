package server

import (
	"encoding/json"
)

type EventHandler func(*Event, *Client)

type Event struct {
	Name        string `json:"event"`
	To          string `json:"to"`
	From        string `json:"from"`
	Transaction struct {
		Username string      `json:"username"`
		Data     interface{} `json:"data"`
	} `json:"transaction"`
}

func NewEvenFromRaw(rawData []byte) (*Event, error) {
	event := new(Event)
	err := json.Unmarshal(rawData, event)
	return event, err
}

func (e *Event) Raw() []byte {
	raw, _ := json.Marshal(e)
	return raw
}
