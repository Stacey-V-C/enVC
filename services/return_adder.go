package services

import (
	"fmt"
	"nvc/types"
	"strings"
)

type ReturnAdder struct {
	receive_channel chan types.NVC_Event
	send_channel    chan types.NVC_Event
}

func NewReturnAdder(receive_channel chan types.NVC_Event, send_channel chan types.NVC_Event) *ReturnAdder {
	return &ReturnAdder{
		receive_channel: receive_channel,
		send_channel:    send_channel,
	}
}

func (r *ReturnAdder) Listen() {
	go ListenAndAddReturnStatement(r)
}

func ListenAndAddReturnStatement(r *ReturnAdder) {
	for input := range r.receive_channel {
		fmt.Println("ReturnAdder received:", input)
		statements := strings.Split(input.Payload, ";")

		statements = statements[:len(statements)-1]

		for i, statement := range statements {
			if !strings.Contains(statement, "returning") {
				statements[i] = statement + " returning *;"
			}
		}
		r.send_channel <- types.NVC_Event{
			Action:  "return",
			Id:      input.Id,
			Payload: strings.Join(statements, ""),
		}
	}
}

func (r *ReturnAdder) GetReceiveChannel() chan types.NVC_Event {
	return r.send_channel
}
