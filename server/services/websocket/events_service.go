package websocket

import (
    "sync"
    "errors"
    "fmt"
)

type EventsService struct {
    events map[string]InitHandler
}

var (
    eventService *EventsService
    eventOnce sync.Once
)

func NewEventsService() IEventsService {
    eventOnce.Do(func () {
        eventService = &EventsService{events: make(map[string]InitHandler)}
    })

    return eventService
}

func (service *EventsService) RegisterEvent(
    name string,
    handler InitHandler,
) error {
    service.events[name] = handler
    return nil
}

func (service *EventsService) ResolveEvent(name string) (IEvent, error) {
    if service.events[name] == nil {
        return nil, errors.New(fmt.Sprintf("Event '%v' has not found", name))
    }

    return service.events[name](), nil
}
