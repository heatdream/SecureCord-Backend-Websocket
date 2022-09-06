package server

import "log"

type Hub struct {
	// Registered clients.
	Clients map[string]*Client

	// Inbound messages from the clients.
	Send chan Event

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	Events map[string]EventHandler
}

func NewHub() *Hub {
	return &Hub{
		Send:       make(chan Event),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]*Client),
		Events:     make(map[string]EventHandler),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.ID] = client
			log.Println(h.Clients)
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.ID]; ok {
				delete(h.Clients, client.ID)
				close(client.Send)
				log.Println(h.Clients)
			}
		case message := <-h.Send:
			if client, ok := h.Clients[message.To]; ok {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client.ID)
				}
			}
		}
	}
}

func (h *Hub) On(name string, handle EventHandler) {
	h.Events[name] = handle
}


func (h *Hub) ParseEvent(e *Event, c *Client) {
	if handle, ok := h.Events[e.Name]; ok {
		go handle(e, c)
	}
}
