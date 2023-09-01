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

type DataSource int8

const (
	SOURCE_SQL    DataSource = iota
	SOURCE_CHROMA DataSource = iota
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
	ERROR             ActionType = "error"
	SQL_RAW_RESULT    ActionType = "sql_raw_result"
	SQL_FILTER_RESULT ActionType = "sql_filter_result"
	MIGRATE           ActionType = "migrate_sql_to_chroma"
	SQL_QUERY         ActionType = "sql_query"
)

type ReceiverMap map[EventType]chan NVC_Event // one event channel per event type

type SubscriberMap map[EventType][]chan NVC_Event // many channels to broadcast to per event type

type NVC_Event struct {
	Action  ActionType `json:"action"`
	Id      string     `json:"id"`
	Payload any        `json:"payload"`
}

type Type_NVC_Event[T NVC_Payload] struct {
	Action  ActionType `json:"action"`
	Id      string     `json:"id"`
	Payload *T         `json:"payload"`
}

type NVC_Payload interface {
	string | TypedError | SQLResultPayload
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

type MigrationPayload interface {
	SQLResultPayload
}

type Migration struct {
	from DataSource
	to   DataSource
	Data any
}

func (m *Migration) Validate() bool {
	switch m.from {
	case SOURCE_SQL:
		valid := reflect.TypeOf(m.Data) == reflect.TypeOf(&[]SQLResultPayload{})

		return valid
	}
	return false
}

var emptyString = ""

var requiredTypes = map[ActionType]reflect.Type{
	ERROR:             reflect.TypeOf(&TypedError{}),
	SQL_RAW_RESULT:    reflect.TypeOf(&SQLResultPayload{}),
	SQL_FILTER_RESULT: reflect.TypeOf(&SQLResultPayload{}),
	MIGRATE:           reflect.TypeOf(&Migration{}),
	SQL_QUERY:         reflect.TypeOf(&emptyString),
}

func (n *NVC_Event) Validate() bool {
	return reflect.TypeOf(n.Payload) == requiredTypes[n.Action]
}

type Subscriber interface {
	Listen()
}
