package websocket

type MessageEvent struct {
    Id int `json:"id"`
}

type broadcastMessage struct {
    Id int `json:"id"`
}

func (event *MessageEvent) Handle(client *Client) error {
    client.wss.Broadcast(&BroadcastInfo{message: &broadcastMessage{event.Id}})

    return nil
}

func InitMessageEventHandler() IEvent {
    return new(MessageEvent)
}

