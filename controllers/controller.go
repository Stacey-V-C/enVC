package controllers

import (
	"context"
	"nvc/types"
	"strings"

	"github.com/google/uuid"
)

type Controller struct {
	ctx         context.Context
	sendChannel chan types.NVC_Event
	errChannel  chan types.NVC_Event
}

func NewController(c, e chan types.NVC_Event) *Controller {
	return &Controller{
		ctx:         context.Background(),
		sendChannel: c,
		errChannel:  e,
	}
}

func (c *Controller) SendSQL(query string) (id string, valid bool) {
	id = uuid.New().String()

	eventType := strings.Split(query, " ")[0]

	eventType = strings.TrimSpace(eventType)

	valid = checkEventType(eventType)

	if !valid {
		c.errChannel <- types.NVC_Event{
			Action: types.ERROR,
			Id:     id,
			Payload: &types.TypedError{
				Type:    types.SQL_ERROR,
				Message: "Invalid SQL action: " + query,
			},
		}
	} else {
		c.sendChannel <- types.NVC_Event{
			Action:  types.SQL_QUERY,
			Id:      id,
			Payload: &query,
		}
	}

	// should we be returning the error here too?
	// was thinking that following the async model we should just
	// give id and let the front end object get the data it needs
	// from there but we should at least tell them it errored
	// so that they know to grab from the right channel - would it be possible
	// for there to be a handler further down the line - or earlier? (middleware?)
	// that could catch and correct an error before the front end gets
	// response - i.e. automatic spelling correction? hmmm
	return
}

func checkEventType(e string) bool {
	SQL_ACTIONS := [5]string{ // check if we can make this const
		"SELECT",
		"INSERT",
		"UPDATE",
		"SHOW",
		"",
	}

	for _, v := range SQL_ACTIONS {
		if e == v {
			return true
		}
	}

	return false
}
