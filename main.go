package main

import (
	components "Diffract/services/components"

	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	Detector := components.NewDetectorService()
	Stage := components.NewStageService()
	HVPS := components.NewHVPSService()

	app := NewApp(Stage, Detector, HVPS)

	err := wails.Run(&options.App{
		Title:  "Diffract",
		Width:  1300,
		Height: 950,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        true,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,

		Bind: []interface{}{
			app,
			Stage,
			Detector,
			HVPS,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
