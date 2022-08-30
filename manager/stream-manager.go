package gomultisse

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StreamManager is the top-level component of this package. It keeps track of living
// event streams and kills them when they become unused.
type StreamManager struct {
	Streams  map[primitive.ObjectID]EventStream
	DoneChan chan primitive.ObjectID
}

// Initializes and returns a new stream manager for incoming connections.
func New() *StreamManager {
	mgr := StreamManager{
		Streams:  make(map[primitive.ObjectID]EventStream),
		DoneChan: make(chan primitive.ObjectID),
	}

	go mgr.ClearUnusedStreams()

	return &mgr
}

// Returns an existing stream or creates a new stream with the provided streamId as 
// the associated ID.
func (mgr *StreamManager) GetStream(streamId primitive.ObjectID) (*EventStream, error) {
	// Ignore "nil" ObjectID if asked for
	if streamId == primitive.NilObjectID {
		return nil, errors.New("requested streamId was null value")
	}

	if es, exists := mgr.Streams[streamId]; exists {
		return &es, nil
	}

	return NewClient(streamId, mgr.DoneChan), nil
}

// Waits for a stream to send a done message then deletes it.
func (mgr *StreamManager) ClearUnusedStreams() {
	for {
		doneId := <-mgr.DoneChan
		delete(mgr.Streams, doneId)
	}
}
