package main

import (
	"embed"
	"nvc/controllers"
	"nvc/services"
	"nvc/types"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	broadcaster := services.NewBroadcaster()

	controller := controllers.NewController(broadcaster.GetChannel(types.SQL))

	return_adder := services.NewReturnAdder(
		make(chan types.NVC_Event),
		broadcaster.GetChannel(types.SQL),
	)

	broadcaster.RegisterSubscriber(types.SQL, return_adder.GetReceiveChannel())

	ui_receiver := services.NewUIReceiver(
		app.ctx,
		broadcaster.GetChannel(types.SQL),
	)

	broadcaster.RegisterSubscriber(types.SQL, ui_receiver.GetReceiveChannel())

	broadcaster.Listen()

	return_adder.Listen()

	ui_receiver.Listen()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "testing",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			controller,
			ui_receiver,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
