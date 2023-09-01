package types

import (
	"reflect"
)

type EventType int8

const (
	AppError EventType = iota
	Chroma   EventType = iota
	SQL      EventType = iota
	UI       EventType = iota
)

type ErrorType int8

const (
	SQL_ERROR    ErrorType = iota
	CHROMA_ERROR ErrorType = iota
	UI_ERROR     ErrorType = iota
	JSON_ERROR   ErrorType = iota
)

type ActionType string

const (
	ERROR                 ActionType = "error"
	SQL_RAW_RESULT        ActionType = "sql_raw_result"
	SQL_FILTER_RESULT     ActionType = "sql_filter_result"
	MIGRATE_SQL_TO_CHROMA ActionType = "migrate_sql_to_chroma"
	SQL_QUERY             ActionType = "sql_query"
)

type ReceiverMap map[EventType]chan NVC_Event // one event channel per event type

type SubscriberMap map[EventType][]chan NVC_Event // many channels to broadcast to per event type

type NVC_Event struct {
	Action  ActionType `json:"action"`
	Id      string     `json:"id"`
	Payload any        `json:"payload"`
}

type Typed_NVC_Event[T any] struct {
	Action  ActionType `json:"action"`
	Id      string     `json:"id"`
	Payload T          `json:"payload"`
}

type TypedError struct {
	Type    ErrorType
	Message string
}

type SQLResultPayload struct {
	DataModel string
	Table     string
	Columns   []string
	Values    [][]string
}

var emptyString = ""

var requiredTypes = map[ActionType]reflect.Type{
	ERROR:                 reflect.TypeOf(&TypedError{}),
	SQL_RAW_RESULT:        reflect.TypeOf(&SQLResultPayload{}),
	SQL_FILTER_RESULT:     reflect.TypeOf(&SQLResultPayload{}),
	MIGRATE_SQL_TO_CHROMA: reflect.TypeOf(&[]NVC_Event{}),
	SQL_QUERY:             reflect.TypeOf(&emptyString),
}

func (n *NVC_Event) validate() bool {
	return reflect.TypeOf(n.Payload) == requiredTypes[n.Action]
}

type Subscriber interface {
	Listen()
}
