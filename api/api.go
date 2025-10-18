package api

import (
	"fmt"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>SSUI PluginLib</title>
    <style>
        body {
            font-family: 'Segoe UI', Arial, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
            margin: 0;
            background: linear-gradient(135deg, #1a1a2e, #16213e);
            color: #e0e0e0;
        }
        h1 {
            color: #00d4ff;
            font-size: 2em;
            margin: 0 0 10px;
        }
        p {
            color: #b0b0cc;
            font-size: 1em;
            margin: 0 0 15px;
        }
        img {
            max-width: 100%;
            border-radius: 8px;
            border: 2px solid #00d4ff;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Example Page</h1>
        <p>SSUI PluginLib</p>
        <img src="https://placecats.com/millie_neo/300/200" alt="Cute cat">
    </div>
</body>
</html>
    `
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "%s", html)
}

func HandleSomethingElse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Something else")
}
