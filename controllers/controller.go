package controllers

import (
	"context"
	"nvc/types"

	"github.com/google/uuid"
)

type Controller struct {
	ctx          context.Context
	send_channel chan types.NVC_Event
}

func NewController(c chan types.NVC_Event) *Controller {
	return &Controller{
		ctx:          context.Background(),
		send_channel: c,
	}
}

func (c *Controller) WaveInSQL() string {
	id := uuid.New().String()

	c.send_channel <- types.NVC_Event{
		Action:  "select",
		Id:      id,
		Payload: "SELECT * FROM users",
	}

	return id
}
