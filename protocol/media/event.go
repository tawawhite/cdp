// Code generated by cdpgen. DO NOT EDIT.

package media

import (
	"github.com/mafredri/cdp/rpcc"
)

// PlayerPropertiesChangedClient is a client for PlayerPropertiesChanged events.
// This can be called multiple times, and can be used to set / override /
// remove player properties. A null propValue indicates removal.
type PlayerPropertiesChangedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*PlayerPropertiesChangedReply, error)
	rpcc.Stream
}

// PlayerPropertiesChangedReply is the reply for PlayerPropertiesChanged events.
type PlayerPropertiesChangedReply struct {
	PlayerID   PlayerID         `json:"playerId"`   // No description.
	Properties []PlayerProperty `json:"properties"` // No description.
}

// PlayerEventsAddedClient is a client for PlayerEventsAdded events. Send
// events as a list, allowing them to be batched on the browser for less
// congestion. If batched, events must ALWAYS be in chronological order.
type PlayerEventsAddedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*PlayerEventsAddedReply, error)
	rpcc.Stream
}

// PlayerEventsAddedReply is the reply for PlayerEventsAdded events.
type PlayerEventsAddedReply struct {
	PlayerID PlayerID      `json:"playerId"` // No description.
	Events   []PlayerEvent `json:"events"`   // No description.
}

// PlayerMessagesLoggedClient is a client for PlayerMessagesLogged events.
// Send a list of any messages that need to be delivered.
type PlayerMessagesLoggedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*PlayerMessagesLoggedReply, error)
	rpcc.Stream
}

// PlayerMessagesLoggedReply is the reply for PlayerMessagesLogged events.
type PlayerMessagesLoggedReply struct {
	PlayerID PlayerID        `json:"playerId"` // No description.
	Messages []PlayerMessage `json:"messages"` // No description.
}

// PlayerErrorsRaisedClient is a client for PlayerErrorsRaised events. Send a
// list of any errors that need to be delivered.
type PlayerErrorsRaisedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*PlayerErrorsRaisedReply, error)
	rpcc.Stream
}

// PlayerErrorsRaisedReply is the reply for PlayerErrorsRaised events.
type PlayerErrorsRaisedReply struct {
	PlayerID PlayerID      `json:"playerId"` // No description.
	Errors   []PlayerError `json:"errors"`   // No description.
}

// PlayersCreatedClient is a client for PlayersCreated events. Called whenever
// a player is created, or when a new agent joins and receives a list of active
// players. If an agent is restored, it will receive the full list of player
// ids and all events again.
type PlayersCreatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*PlayersCreatedReply, error)
	rpcc.Stream
}

// PlayersCreatedReply is the reply for PlayersCreated events.
type PlayersCreatedReply struct {
	Players []PlayerID `json:"players"` // No description.
}