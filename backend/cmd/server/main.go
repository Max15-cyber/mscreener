package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
	"mscreener/internal/api"
)

func main() {
    api.StartPriceWorker()

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "ok")
    })

    http.HandleFunc("/price", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "symbol": "BTCUSDT",
            "price":  api.CurrentBTCPrice,
            "time":   time.Now().Format(time.RFC3339),
        })
    })

	fmt.Println("╔══════════════════════════════════════════════════╗")
	fmt.Println("║                 M S C R E E N E R                ║")
	fmt.Println("║             Trading Platform v0.1.0              ║")
	fmt.Println("╠══════════════════════════════════════════════════╣")
	fmt.Println("║ Server running on: http://localhost:8080         ║")
	fmt.Println("║ Press Ctrl+C to stop                             ║")
	fmt.Println("╚══════════════════════════════════════════════════╝")
    http.ListenAndServe(":8080", nil)
}
