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
```go
package main

import (
    "log"
    "context"
    "time"

    "github.com/brandonwhited-dev/twitchkit-eventsub"
    "github.com/brandonwhited-dev/twitchkit-helix"
)

func main() {
    oauth := "YOUR_TWITCH_OAUTH"       // no OAuth Prefix
    clientID := "YOUR_CLIENT_ID"       // Client ID linked with the oauth

    helixClient := twitchkithelix.NewClient(clientID, oauth, nil)

    eventSubChan := make(chan twitchkiteventsub.Event)
    eventSubWebsocket := twitchkiteventsub.NewEventSubWebsocket(eventSubChan)
    
    go func() {
        eventSubWebsocket.Connect()
    }

    firstWelcome := true // Guarding against multiple session welcomes
    for event := range eventSubChan {
        switch event.MessageType {
        case "session_welcome":
            if firstWelcome {
                firstWelcome = false

                //initialze connections
                ctx := context.Background()
                ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
                defer cancel()

                _, err := helixClient.EventStreamOnline(
                    ctx,
                    eventSubWebsocket.Session.ID,
                    twitchkithelix.ConditionStreamOnline{
                        BroadcasterUserID: "BROADCASTERS_ID",
                    },
                )
                if err != nil {
                    log.Fatal("SUBSCRIPTION INIT ERROR: ", err)
                }
            }
        case "notification":
            // Handle the event
            log.Println(event)
        }
}
```
