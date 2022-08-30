package gomultisse

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ClientChan is an alias for a string channel. It is used by being sent over
// other channels to communicate about clients.
type ClientChan chan string

// EventStream represents a living event stream. It tracks its own ID, channels
// for communicating about new or closed client connections, living channels,
// and a channel for sending messages.
type EventStream struct {
	Id               primitive.ObjectID
	Message          chan string
	NewClientChan    chan ClientChan
	ClosedClientChan chan ClientChan
	Clients          map[chan string]bool
}

// Creates a new event stream server and opens channels
func NewClient(streamId primitive.ObjectID, done chan primitive.ObjectID) *EventStream {
	es := &EventStream{
		Id:               streamId,
		Message:          make(chan string),
		NewClientChan:    make(chan ClientChan),
		ClosedClientChan: make(chan ClientChan),
		Clients:          make(map[chan string]bool),
	}

	go es.listen(done)

	return es
}

// Handles adding clients, removing clients, and streaming messages
func (es *EventStream) listen(done chan primitive.ObjectID) {
	for {
		select {
		case client := <-es.NewClientChan:
			es.Clients[client] = true

		case client := <-es.ClosedClientChan:
			delete(es.Clients, client)

			// Alert manager if this event stream is unused
			if len(es.Clients) == 0 {
				done <- es.Id
			}

		case eventMsg := <-es.Message:
			for clientMsgChan := range es.Clients {
				clientMsgChan <- eventMsg
			}
		}
	}
}
