package services

import (
	"context"
	"fmt"
	"nvc/types"
)

type UIReceiver struct {
	ctx         context.Context
	sqlChannel  chan types.NVC_Event
	heldResults map[string][]string
}

func NewUIReceiver(ctx context.Context, sqlChannel chan types.NVC_Event) *UIReceiver {
	return &UIReceiver{
		ctx:         ctx,
		sqlChannel:  sqlChannel,
		heldResults: make(map[string][]string),
	}
}

func (u *UIReceiver) Listen() {
	go ListenAndStoreResults(u, u.sqlChannel)
}

func ListenAndStoreResults(u *UIReceiver, channel chan types.NVC_Event) {
	for input := range channel {
		fmt.Println("UIReceiver received:", input)
		results := [2]string{*input.Payload, input.Action}
		u.heldResults[input.Id] = results[:]
	}
}

func (u *UIReceiver) PullData(id string) []string {
	return u.heldResults[id]
}

func (u *UIReceiver) GetReceiveChannel() chan types.NVC_Event {
	return u.sqlChannel
}
