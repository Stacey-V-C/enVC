package services

import (
	"context"
	"encoding/json"
	"nvc/types"
)

type SQLListener struct {
	conn       *SQLConnection
	proc       *SQLProcesser
	sqlChanIn  chan types.NVC_Event
	sqlChanOut chan types.NVC_Event
	errChanOut chan types.NVC_Event
}

func NewSQLListener(
	ctx context.Context,
	sqlIn chan types.NVC_Event,
	sqlOut chan types.NVC_Event,
	errOut chan types.NVC_Event,
) SQLListener {
	conn := NewSQLConnection(ctx)
	proc := NewSQLProcesser()

	return SQLListener{
		conn:       &conn,
		proc:       &proc,
		sqlChanIn:  sqlIn,
		sqlChanOut: sqlOut,
		errChanOut: errOut,
	}
}

func (s *SQLListener) Listen() {
	for {
		select {
		case event := <-s.sqlChanIn:
			switch event.Action {
			case types.SQL_QUERY:
				payload := event.Payload.(*string)
				s.RunQuery(event.Id, *payload) // TODO: type assertion
			}
		}
	}
}

func (s *SQLListener) RunQuery(id, q string) {
	// how do we get the datamodel hmmmm have to use some sort of context

	requestedData, modifiedQuery, err := s.proc.parseQueryReturning(NVCDataModel, q)

	if err != nil {
		eStr := err.Error()

		s.errChanOut <- types.NVC_Event{
			Action:  "sql_error",
			Id:      id,
			Payload: &eStr,
		}
		return
	}

	var numCols int
	var query string

	if modifiedQuery == nil {
		numCols = len(requestedData.GetColumns())
		query = q
	} else {
		numCols = len(NVCDataModel.GetTable(requestedData.GetName()).GetColumns())
		query = *modifiedQuery
	}

	resultSet, err := s.conn.GetResultSetAsStrings(query, numCols)

	if err != nil {
		eStr := err.Error()

		s.errChanOut <- types.NVC_Event{
			Action:  "sql_error",
			Id:      id,
			Payload: &eStr,
		}
		return
	}

	if modifiedQuery != nil {
		resultBytes, err := json.Marshal(resultSet)

		if err != nil {
			eStr := err.Error()

			s.errChanOut <- types.NVC_Event{
				Action:  "json_error",
				Id:      id,
				Payload: &eStr,
			}
			return
		}

		resultStr := string(resultBytes)

		s.sqlChanOut <- types.NVC_Event{
			Action:  "sql_raw_result",
			Id:      id,
			Payload: &resultStr,
		}

		filteredResults, err := s.proc.filterResults(
			requestedData.GetColumns(),
			resultSet,
			NVCDataModel.GetTable(requestedData.GetName()).GetColumns(),
		)

		if err != nil {
			eStr := err.Error()

			s.errChanOut <- types.NVC_Event{
				Action:  "sql_error",
				Id:      id,
				Payload: &eStr,
			}
			return
		}

		filteredResultBytes, err := json.Marshal(filteredResults)

		if err != nil {
			eStr := err.Error()

			s.errChanOut <- types.NVC_Event{
				Action: types.ERROR,
				Id:     id,
				Payload: &types.TypedError{
					Type:    types.JSON_ERROR,
					Message: eStr,
				},
			}
			return
		}

		filteredResultStr := string(filteredResultBytes)

		s.sqlChanOut <- types.NVC_Event{
			Action:  "sql_result_filtered",
			Id:      id,
			Payload: &filteredResultStr,
		}
	}
}
