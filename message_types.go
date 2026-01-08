package twitchkiteventsub

import "time"

// Metadata is common information shared by all EventSub messages
type Metadata struct {
	// MessageID is the unique id of the message
	MessageID string `json:"message_id"`

	// MessageType is the type of message (e.g., "session_welcome", "notification")
	MessageType string `json:"message_type"`

	// MessageTimestamp is when the message was sent
	MessageTimestamp time.Time `json:"message_timestamp"`

	// SubscriptionType is the type of subscription (optional)
	SubscriptionType string `json:"subscription_type,omitempty"`

	// SubscriptionVersion is the version of the subscription (optional)
	SubscriptionVersion string `json:"subscription_version,omitempty"`
}

// Session contains information about an EventSub session
type Session struct {
	// ID is the session's unique id
	ID string `json:"id"`

	// Status is the current session status
	Status string `json:"status"`

	// KeepaliveTimeoutSeconds is how often a keepalive message is expected
	KeepaliveTimeoutSeconds *int `json:"keepalive_timeout_seconds"`

	// ReconnectURL is the URL to reconnect the session if needed
	ReconnectURL *string `json:"reconnect_url"`

	// ConnectedAt is when the session was established
	ConnectedAt time.Time `json:"connected_at"`
}

// Subscription contains details about an EventSub subscription
type Subscription struct {
	// ID is the subscription's unique id
	ID string `json:"id"`

	// Status is the current status of the subscription
	Status string `json:"status"`

	// Type is the subscription type (e.g., "stream.online")
	Type string `json:"type"`

	// Version is the subscription version
	Version string `json:"version"`

	// Cost is the cost in bits to create this subscription
	Cost int `json:"cost"`

	// Condition defines the specific rules for the subscription
	Condition any `json:"condition"`

	// Transport describes how the subscription delivers events
	Transport Transport `json:"transport"`

	// CreatedAt is when the subscription was created
	CreatedAt time.Time `json:"created_at"`
}

// Transport contains the delivery method of an EventSub subscription
type Transport struct {
	// Method is the delivery method (websocket or webhook)
	Method string `json:"method"`

	// SessionID is the session id associated with this transport
	SessionID string `json:"session_id"`
}

// ====================== MESSAGE TYPES ======================

// WelcomeMessage is sent when a session is first created
type WelcomeMessage struct {
	// Metadata contains common message information
	Metadata Metadata `json:"metadata"`

	// Payload contains the session information
	Payload struct {
		// Session is the session info for this welcome message
		Session Session `json:"session"`
	} `json:"payload"`
}

// KeepaliveMessage is sent periodically to keep the session alive
type KeepaliveMessage struct {
	// Metadata contains common message information
	Metadata Metadata `json:"metadata"`

	// Payload is empty for keepalive messages
	Payload struct{} `json:"payload"`
}

// NotificationMessage is sent when a subscription event occurs
type NotificationMessage struct {
	// Metadata contains common message information
	Metadata Metadata `json:"metadata"`

	// Payload contains the subscription and event details
	Payload struct {
		// Subscription contains info about the subscription
		Subscription Subscription `json:"subscription"`

		// Event contains the event data; type depends on subscription
		Event interface{} `json:"event"`
	} `json:"payload"`
}

// ReconnectMessage is sent when the session should reconnect
type ReconnectMessage struct {
	// Metadata contains common message information
	Metadata Metadata `json:"metadata"`

	// Payload contains the session info for reconnect
	Payload struct {
		// Session is the session info to reconnect to
		Session Session `json:"session"`
	} `json:"payload"`
}

// RevocationMessage is sent when a subscription is revoked
type RevocationMessage struct {
	// Metadata contains common message information
	Metadata Metadata `json:"metadata"`

	// Payload contains the revoked subscription info
	Payload struct {
		// Subscription contains info about the revoked subscription
		Subscription Subscription `json:"subscription"`
	} `json:"payload"`
}
