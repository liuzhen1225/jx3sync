package main

import (
	"changeme/service"
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	role := service.NewRole()
	// Create application with options
	err := wails.Run(&options.App{
		Title:         "jx3sync",
		Width:         600,
		Height:        400,
		Frameless:     true,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			role.Startup(ctx)
		},
		Bind: []interface{}{
			app,
			role,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
