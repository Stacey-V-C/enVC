package types

type EventType int8

const (
	AppError EventType = iota
	Chroma   EventType = iota
	SQL      EventType = iota
	UI       EventType = iota
)

type ReceiverMap map[EventType]chan NVC_Event // one event channel per event type

type SubscriberMap map[EventType][]chan NVC_Event // many channels to broadcast to per event type

type NVC_Event struct {
	Action  string
	Id      string
	Payload *string
}

type Subscriber interface {
	Listen()
}
