package services

import (
	"context"
	"fmt"
	"nvc/types"
)

type UIReceiver struct {
	ctx         context.Context
	sqlChannel  chan types.NVC_Event
	heldResults map[string]types.NVC_Event
}

func NewUIReceiver(ctx context.Context, sqlChannel chan types.NVC_Event) *UIReceiver {
	return &UIReceiver{
		ctx:         ctx,
		sqlChannel:  sqlChannel,
		heldResults: make(map[string]types.NVC_Event),
	}
}

func (u *UIReceiver) Listen() {
	go ListenAndStoreResults(u, u.sqlChannel)
}

func ListenAndStoreResults(u *UIReceiver, channel chan types.NVC_Event) {
	for input := range channel {
		fmt.Println("UIReceiver received:", input)

		u.heldResults[input.Id] = input
	}
}

func (u *UIReceiver) PullData(id string) types.NVC_Event {
	return u.heldResults[id]
}

func (u *UIReceiver) GetReceiveChannel() chan types.NVC_Event {
	return u.sqlChannel
}
