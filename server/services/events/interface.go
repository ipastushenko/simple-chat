package events

type IEventData interface {}

type EventHandler func()

type IEvent interface {
    GetName() string
    InitData() IEventData
}

type IEvents interface {
    RegisterEvent(IEvent) error
}
