# twitchkit-eventsub
"twitchkit-eventsub" is a Go library for interacting with Twitch’s EventSub API, allowing you to listen to events such as follows, subscriptions, donations, and more.

To register EventSub subscriptions, use [`twitchkit-helix`](https://github.com/brandonwhited-dev/twitchkit-helix).
EventSub subscriptions cannot be created until Twitch assigns a WebSocket session.
Twitch provides this session ID in the session_welcome message.

For detailed field info, see the [pkg.go.dev documentation](https://pkg.go.dev/github.com/brandonwhited-dev/twitchkit-eventsub).
# Features
- Creates a connection to Twitch Event Sub over Web Socket
- Manages state of the sessions
- Decodes all events received
# Installation 
```bash
go get github.com/brandonwhited-dev/twitchkit-eventsub
```
# Example
See [`twitchkit-examples`](https://github.com/brandonwhited-dev/twitchkit-examples) for examples of twitckhit implementation.
