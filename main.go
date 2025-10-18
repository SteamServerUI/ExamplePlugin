package main

import (
	"embed"
	"fmt"
	"log"
	"sync"

	"github.com/SteamServerUI/ExamplePlugin/api"
	"github.com/SteamServerUI/ExamplePlugin/global"
	"github.com/SteamServerUI/PluginLib"
)

//go:embed assets/*
var assets embed.FS

var (
	settingsResponse PluginLib.SettingsResponse
	wg               sync.WaitGroup
)

func main() {

	// Register embedded assets
	global.AssetManager = PluginLib.RegisterAssets(&assets)

	PluginLib.InitConfig(global.PluginName, global.DefaultLogLevel)
	GetGameserverRunningStatus()
	LogSomething()
	SaveASetting()
	GetSetting()
	//GetAllSettings()
	ExposeAPI(&wg)
	wg.Wait()
}

func GetGameserverRunningStatus() {
	rsp, err := PluginLib.GetServerStatus()
	if err != nil {
		log.Fatalf("Failed to get server status: %v", err)
	}

	fmt.Println("Gameserver running:", rsp.Status, "UUID:", rsp.UUID)
}

func LogSomething() {
	// allows either a message
	PluginLib.Log("Test")
	// or a message and a log level
	PluginLib.Log("Test", "Info")
	// also allows proper error handling
	err := PluginLib.Log("Test", "Non-Existing-Level")
	if err != nil {
		fmt.Println("Error (expected, since level doesn't exist):", err)
	}
}

func SaveASetting() {

	payload := map[string]string{"GameBranch": "public"}

	if _, err := PluginLib.Post("/api/v2/settings/save", &payload, &settingsResponse); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Setting saved:", settingsResponse.Message)
}

func GetSetting() {
	value, err := PluginLib.GetSetting("GameBranch")
	if err != nil {
		log.Printf("Failed to get setting: %v", err)
		return
	}
	fmt.Printf("Setting 'GameBranch': %v\n", value)
}

func GetAllSettings() {
	settings := PluginLib.GetAllSettings()
	for key, value := range settings {
		fmt.Printf("Setting '%s': %v\n", key, value)
	}
}

func ExposeAPI(wg *sync.WaitGroup) {

	PluginLib.RegisterRoute("/", api.HandleTextFromAssetsManager)
	PluginLib.RegisterRoute("/something", api.HandleSomethingElse)
	PluginLib.RegisterRoute("/image", api.HandleBinaryFromAssetsManager)
	PluginLib.ExposeAPI(wg)
	PluginLib.RegisterPluginAPI()

}
