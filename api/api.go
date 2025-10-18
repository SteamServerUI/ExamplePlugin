package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SteamServerUI/ExamplePlugin/global"
)

func HandleSomethingElse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Something else")
}

func HandleIndexFromAssetsManager(w http.ResponseWriter, r *http.Request) {
	data, err := global.AssetManager.GetAssetString("assets/index.html")
	if err != nil {
		log.Fatalf("Failed to read asset: %v", err)
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "%s", data)
}
