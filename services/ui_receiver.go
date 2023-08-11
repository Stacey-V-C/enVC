package services

import (
	"context"
	"fmt"
	"nvc/types"
)

type UIReceiver struct {
	ctx          context.Context
	sql_channel  chan types.NVC_Event
	held_results map[string][]string
}

func NewUIReceiver(ctx context.Context, sql_channel chan types.NVC_Event) *UIReceiver {
	return &UIReceiver{
		ctx:          ctx,
		sql_channel:  sql_channel,
		held_results: make(map[string][]string),
	}
}

func (u *UIReceiver) Listen() {
	go ListenAndStoreResults(u, u.sql_channel)
}

func ListenAndStoreResults(u *UIReceiver, channel chan types.NVC_Event) {
	for input := range channel {
		fmt.Println("UIReceiver received:", input)
		results := [2]string{input.Payload, input.Action}
		u.held_results[input.Id] = results[:]
	}
}

func (u *UIReceiver) PullData(id string) []string {
	return u.held_results[id]
}

func (u *UIReceiver) GetReceiveChannel() chan types.NVC_Event {
	return u.sql_channel
}
