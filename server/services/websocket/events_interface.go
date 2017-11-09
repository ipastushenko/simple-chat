package websocket

type InitHandler func() IEvent

type IEvent interface {
    Handle(client *Client) error
}

type IEventsService interface {
    RegisterEvent(name string, handler InitHandler) error
    ResolveEvent(name string) (IEvent, error)
}
