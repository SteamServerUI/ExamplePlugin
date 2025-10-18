package main

import (
	"fmt"
	"log"
	"sync"
	"test/api"

	"github.com/SteamServerUI/PluginLib"
)

var settingsResponse PluginLib.SettingsResponse
var wg sync.WaitGroup

func main() {

	PluginLib.InitConfig("ExamplePlugin", "Info")
	GetStatus()
	LogSomething()
	SaveASetting()
	GetSetting()
	//GetAllSettings()
	ExposeAPI(&wg)
	wg.Wait()
}

func GetStatus() {
	rsp, err := PluginLib.GetServerStatus()
	if err != nil {
		log.Fatalf("Failed to get server status: %v", err)
	}

	fmt.Println("Server running:", rsp.Status, "UUID:", rsp.UUID)
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

	if err := PluginLib.Post("/api/v2/settings/save", &payload, &settingsResponse); err != nil {
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

	PluginLib.RegisterRoute("/", api.HandleIndex)
	PluginLib.RegisterRoute("/something", api.HandleSomethingElse)
	PluginLib.ExposeAPI(wg)
	PluginLib.RegisterPluginAPI()
	PluginLib.Log("Registered in SSUI API")

}
