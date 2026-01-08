package twitchkiteventsub

// Reward represents a channel points reward
type Reward struct {
	// ID is the unique id of the reward
	ID string `json:"id"`

	// Title is the name of the reward
	Title string `json:"title"`

	// Cost is the number of points required to redeem
	Cost int `json:"cost"`

	// Prompt is the reward description
	Prompt string `json:"prompt"`
}

// ChannelPointsRedemptionEvent is triggered when a viewer redeems a reward
type ChannelPointsRedemptionEvent struct {
	// ID is the unique id for this redemption
	ID string `json:"id"`

	// BroadcasterUserID is the broadcaster receiving the redemption
	BroadcasterUserID string `json:"broadcaster_user_id"`

	// BroadcasterUserName is the display name of the broadcaster
	BroadcasterUserName string `json:"broadcaster_user_name"`

	// UserID is the ID of the user redeeming the reward
	UserID string `json:"user_id"`

	// UserName is the display name of the user redeeming the reward
	UserName string `json:"user_name"`

	// UserInput is an optional message provided by the user
	UserInput string `json:"user_input"`

	// Status indicates redemption state: "FULFILLED", "CANCELED", "UNFULFILLED"
	Status string `json:"status"`

	// RedeemedAt is the timestamp when the reward was redeemed
	RedeemedAt string `json:"redeemed_at"`

	// Reward contains details of the redeemed reward
	Reward Reward `json:"reward"`
}

// ChannelRaidEvent is triggered when a broadcaster raids another channel
type ChannelRaidEvent struct {
	// FromBroadcasterUserID is the ID of the raiding broadcaster
	FromBroadcasterUserID string `json:"from_broadcaster_user_id"`

	// FromBroadcasterUserLogin is the login of the raiding broadcaster
	FromBroadcasterUserLogin string `json:"from_broadcaster_user_login"`

	// FromBroadcasterUserName is the display name of the raiding broadcaster
	FromBroadcasterUserName string `json:"from_broadcaster_user_name"`

	// ToBroadcasterUserID is the ID of the raided broadcaster
	ToBroadcasterUserID string `json:"to_broadcaster_user_id"`

	// ToBroadcasterUserLogin is the login of the raided broadcaster
	ToBroadcasterUserLogin string `json:"to_broadcaster_user_login"`

	// ToBroadcasterUserName is the display name of the raided broadcaster
	ToBroadcasterUserName string `json:"to_broadcaster_user_name"`

	// Viewers is the number of viewers in the raid
	Viewers int `json:"viewers"`
}

// AdBreakEvent is triggered when an ad break starts
type AdBreakEvent struct {
	// DurationSeconds is the length of the ad in seconds
	DurationSeconds int `json:"duration_seconds"`

	// StartedAt is the timestamp when the ad started
	StartedAt string `json:"started_at"`

	// IsAutomatic indicates if the ad was automatically triggered
	IsAutomatic bool `json:"is_automatic"`

	// BroadcasterUserID is the ID of the broadcaster running the ad
	BroadcasterUserID string `json:"broadcaster_user_id"`

	// BroadcasterUserLogin is the login of the broadcaster running the ad
	BroadcasterUserLogin string `json:"broadcaster_user_login"`

	// BroadcasterUserName is the display name of the broadcaster
	BroadcasterUserName string `json:"broadcaster_user_name"`

	// RequesterUserID is the ID of the user who requested the ad (manual)
	RequesterUserID string `json:"requester_user_id"`

	// RequesterUserLogin is the login of the requester
	RequesterUserLogin string `json:"requester_user_login"`

	// RequesterUserName is the display name of the requester
	RequesterUserName string `json:"requester_user_name"`
}

// StreamOnlineEvent is triggered when a broadcaster goes live
type StreamOnlineEvent struct {
	// ID is the unique id of the stream
	ID string `json:"id"`

	// BroadcasterUserID is the broadcaster going live
	BroadcasterUserID string `json:"broadcaster_user_id"`

	// BroadcasterUserLogin is the login of the broadcaster
	BroadcasterUserLogin string `json:"broadcaster_user_login"`

	// BroadcasterUserName is the display name of the broadcaster
	BroadcasterUserName string `json:"broadcaster_user_name"`

	// Type is the stream type (usually "live")
	Type string `json:"type"`

	// StartedAt is the timestamp when the stream started
	StartedAt string `json:"started_at"`
}

// StreamOfflineEvent is triggered when a broadcaster goes offline
type StreamOfflineEvent struct {
	// BroadcasterUserID is the broadcaster going offline
	BroadcasterUserID string `json:"broadcaster_user_id"`

	// BroadcasterUserLogin is the login of the broadcaster
	BroadcasterUserLogin string `json:"broadcaster_user_login"`

	// BroadcasterUserName is the display name of the broadcaster
	BroadcasterUserName string `json:"broadcaster_user_name"`
}

// ChannelUpdateEvent is triggered when a broadcaster updates their channel information
type ChannelUpdateEvent struct {
	// BroadcasterUserID is the ID of the broadcaster
	BroadcasterUserID string `json:"broadcaster_user_id"`

	// BroadcasterUserLogin is the login of the broadcaster
	BroadcasterUserLogin string `json:"broadcaster_user_login"`

	// BroadcasterUserName is the display name of the broadcaster
	BroadcasterUserName string `json:"broadcaster_user_name"`

	// Title is the updated stream title
	Title string `json:"title"`

	// Language is the updated stream language (ISO 639-1)
	Language string `json:"language"`

	// CategoryID is the updated game/category ID
	CategoryID string `json:"category_id"`

	// ContentClassificationLabels are the updated content classification labels
	ContentClassificationLables []string `json:"content_classification_labels"`
}
