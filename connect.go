package twitchkiteventsub

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"sync"
)

// SessionConfig represents the basic information to establish a connection
type SessionConfig struct {
	// Session is a pointer that represents the sessions data
	Session *Session

	// Events represents the channel that events will be sent to
	Events chan<- Event

	// mu represents a RWMutex to pretect session data
	mu sync.RWMutex
}

// Event represents the information of a recieved event
type Event struct {
	// MessageType represents the type of message
	// Welcome, Keepalive, Notification, Reconnect, Revocation
	MessageType string

	// SubscriptionType represents the event type
	SubscriptionType string

	// Data represents all data for the event
	Data []byte
}

// NewEventSubWebsocket is a init function for session config
func NewEventSubWebsocket(eventMessageChan chan<- Event) *SessionConfig {
	// Twitch EventSub WebSocket endpoint
	u := url.URL{Scheme: "wss", Host: "eventsub.wss.twitch.tv", Path: "/ws"}
	urlString := u.String()
	s := &SessionConfig{
		Events: eventMessageChan,
		Session: &Session{
			ReconnectURL: &urlString,
		},
	}
	return s
}

// Connect opens a connection to Twitch Event Sub
func (s *SessionConfig) Connect() {
	// Dial the Twitch WebSocket
	log.Printf("Connecting to %s", *s.Session.ReconnectURL)
	c, _, err := websocket.DefaultDialer.Dial(*s.Session.ReconnectURL, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// Read messages from Twitch
	for {
		websocketMessageType, message, err := c.ReadMessage()
		if err != nil {
			if websocketMessageType == websocket.PingMessage {
				log.Println("Websocket PING IN ERROR BLOCK")
				c.WriteMessage(websocket.PongMessage, nil)
				continue
			}
			log.Println("Read Error:", err)
			return
		}

		if websocketMessageType == websocket.PingMessage {
			log.Println("Websocket PING")
			c.WriteMessage(websocket.PongMessage, nil)
			continue
		}

		// parsing the data as a map
		var raw map[string]any
		err = json.Unmarshal(message, &raw)
		if err != nil {
			log.Println("error:", err)
		}
		metadataMap, ok := raw["metadata"].(map[string]any)
		if !ok {
			log.Println("Error reading message_type")
			return
		}

		// Handle the different message types
		switch metadataMap["message_type"] {
		case "session_welcome":
			var parsedMessage WelcomeMessage
			err := json.Unmarshal(message, &parsedMessage)
			if err != nil {
				log.Println("Welcome Message Decoding Error: ", err)
				return
			}
			s.mu.Lock()
			s.Session = &parsedMessage.Payload.Session
			s.mu.Unlock()
			s.Events <- Event{MessageType: parsedMessage.Metadata.MessageType, SubscriptionType: parsedMessage.Metadata.SubscriptionType, Data: message}

		case "session_keepalive":
			var parsedMessage KeepaliveMessage
			err := json.Unmarshal(message, &parsedMessage)
			if err != nil {
				log.Println("Keep Alive Message Decoding Error: ", err)
				return
			}
			s.Events <- Event{MessageType: parsedMessage.Metadata.MessageType, SubscriptionType: parsedMessage.Metadata.SubscriptionType, Data: message}

		case "notification":
			var parsedMessage NotificationMessage
			err := json.Unmarshal(message, &parsedMessage)
			if err != nil {
				log.Println("Notification Message Decoding Error: ", err)
				return
			}
			s.Events <- Event{MessageType: parsedMessage.Metadata.MessageType, SubscriptionType: parsedMessage.Metadata.SubscriptionType, Data: message}

		case "reconnect_message":
			var parsedMessage ReconnectMessage
			err := json.Unmarshal(message, &parsedMessage)
			if err != nil {
				log.Println("Reconnect Message Decoding Error: ", err)
				return
			}
			s.mu.Lock()
			s.Session.ReconnectURL = parsedMessage.Payload.Session.ReconnectURL
			s.mu.Unlock()
			s.Connect()

		}
	}
}
