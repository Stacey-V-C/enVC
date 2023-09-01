package services

import (
	"context"
	"nvc/pb"
	"nvc/types"
	"os"
	"time"

	"google.golang.org/grpc"
)

type SQLToChromaBridge struct {
	client        pb.SQLToChromaListenerClient
	SQLChanIn     chan types.NVC_Event
	ChromaChanOut chan types.NVC_Event
	ErrChanOut    chan types.NVC_Event
}

func NewSQLToChromaBridge(
	SQLChanIn, ChromaChanOut, ErrChanOut chan types.NVC_Event,
) SQLToChromaBridge {
	client := pb.NewSQLToChromaListenerClient(nil)

	return SQLToChromaBridge{
		client:        client,
		SQLChanIn:     SQLChanIn,
		ChromaChanOut: ChromaChanOut,
		ErrChanOut:    ErrChanOut,
	}
}

func (s *SQLToChromaBridge) Listen() {
	for {
		select {
		case event := <-s.SQLChanIn:
			if event.Action == types.SQL_RAW_RESULT && event.Validate() {
				if event.Validate() {
					payload := event.Payload.(*types.SQLResultPayload)

					go s.Send(payload, &s.ErrChanOut)
					break
				}
			}

			if event.Action == types.MIGRATE && event.Validate() {
				payload := event.Payload.(*types.Migration)
				if payload.Validate() {
					data := payload.Data.(*[]types.SQLResultPayload)

					for _, result := range *data {
						go s.Send(&result, &s.ErrChanOut)
					}

					break
				}
			}
		}
	}
}

// so wait so the controller should request sql results and then pass a migration event to the bridge??
// it also needs to listen though - do migrations get sent on the sql channel or the chroma channel?
// maybe CMD/GET Don't get broadcast on the same channels as POST results

func (s *SQLToChromaBridge) Send(
	data *types.SQLResultPayload,
	errChan *chan types.NVC_Event,
) {
	chromaAddr := os.Getenv("py_chroma_addr") // make this consistent with python which reads from confi file

	if chromaAddr == "" {
		chromaAddr = "localhost:50051"
	}

	conn, err := grpc.Dial(chromaAddr, grpc.WithInsecure())

	if err != nil {
		*errChan <- types.NVC_Event{
			Action: types.ERROR,
			Id:     "sql_to_chroma_bridge",
			Payload: types.TypedError{
				Type:    types.CHROMA_ERROR,
				Message: err.Error(),
			},
		}
		return
	}

	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*300)
	defer cancel()

	c := pb.NewSQLToChromaListenerClient(conn)

	formattedValues := []*pb.Value{}

	for _, row := range data.Values {
		newValue := &pb.Value{
			Values: row,
		}
		formattedValues = append(formattedValues, newValue)
	}

	r, err := c.LogSQLAction(ctx, &pb.SQLAction{
		DataModel: data.DataModel,
		Table:     data.Table,
		Columns:   data.Columns,
		Values:    formattedValues,
	})

	sendErr := ""

	if err != nil {
		sendErr = err.Error()
	} else if r.GetError() != "" {
		sendErr = r.GetError()
	}

	if sendErr != "" {
		*errChan <- types.NVC_Event{
			Action: types.ERROR,
			Id:     "sql_to_chroma_bridge",
			Payload: types.TypedError{
				Type:    types.CHROMA_ERROR,
				Message: sendErr,
			},
		}
		return
	} else {
		/*
			TODO:
				send to chroma channel, I think?  but what kind of results do we want?
				list of Ids? would just a success message be useful?
				gotta start thinking about how BACKGROUND results are displayed, used, etc
		*/
		return
	}

}

// maybe each listener has a CMD channel and a BG channel???
// how much overhead are channels anyway
// the other options is to have them parse on each event first whether it is a CMD or a BG event hmmm
