package twitchkiteventsub

import (
	"encoding/json"
)

// Generic function to decode the event data into a struct
func DecodeNotificationEvent[T any](data []byte) (*T, error) {
	// unmarshal the whole message
	var message NotificationMessage
	err := json.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}

	// marshal the event
	eventBytes, err := json.Marshal(message.Payload.Event)
	if err != nil {
		return nil, err
	}

	// unmarshal and return the event data
	var event T
	err = json.Unmarshal(eventBytes, &event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
