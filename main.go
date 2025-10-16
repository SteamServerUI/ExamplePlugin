package main

import (
	"fmt"

	"github.com/SteamServerUI/PluginLib"
)

type ServerStatus struct {
	Status bool   `json:"isRunning"`
	UUID   string `json:"uuid"`
}

func main() {
	var response ServerStatus

	if err := PluginLib.Get("/api/v2/server/status", &response); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Server running:", response.Status, "UUID:", response.UUID)
}
