package services

import (
	"context"
	"fmt"
	"nvc/types"
)

type Broadcaster struct {
	ctx         context.Context
	receivers   types.ReceiverMap
	subscribers types.SubscriberMap
}

func NewBroadcaster() *Broadcaster {
	chroma_receive_channel := make(chan types.NVC_Event)
	sql_receive_channel := make(chan types.NVC_Event)
	error_receive_channel := make(chan types.NVC_Event)
	ui_receive_channel := make(chan types.NVC_Event)

	broadcaster := Broadcaster{
		receivers:   make(types.ReceiverMap),
		subscribers: make(types.SubscriberMap),
	}

	broadcaster.receivers[types.Chroma] = chroma_receive_channel
	broadcaster.receivers[types.SQL] = sql_receive_channel
	broadcaster.receivers[types.AppError] = error_receive_channel
	broadcaster.receivers[types.UI] = ui_receive_channel

	return &broadcaster
}

func (b *Broadcaster) RegisterSubscriber(channel types.EventType, c chan types.NVC_Event) {
	b.subscribers[channel] = append(b.subscribers[channel], c)
}

func (b *Broadcaster) Broadcast(channel types.EventType, event types.NVC_Event) {
	for _, subscriber := range b.subscribers[channel] {
		subscriber <- event
	}
}

func (b *Broadcaster) Listen() {
	go ListenAndBroadCast(b, b.receivers[types.Chroma], types.Chroma)

	go ListenAndBroadCast(b, b.receivers[types.SQL], types.SQL)

	go ListenAndBroadCast(b, b.receivers[types.AppError], types.AppError)

	go ListenAndBroadCast(b, b.receivers[types.UI], types.UI)
}

func (b *Broadcaster) GetChannel(channel types.EventType) chan types.NVC_Event {
	return b.receivers[channel]
}

func ListenAndBroadCast(
	broadcaster *Broadcaster,
	channel chan types.NVC_Event,
	eventType types.EventType,
) {
	for input := range channel {
		fmt.Println("Broadcaster received: ", input, " on channel: ", eventType)
		broadcaster.Broadcast(eventType, input)
	}
}
