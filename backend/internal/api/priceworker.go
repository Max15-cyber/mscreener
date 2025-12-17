package api

import (
	"encoding/json"
	"io"
	"net/http"
	"log"
	"time"
)

var CurrentBTCPrice float64

func StartPriceWorker() {
    log.Println("StartPriceWorker → started")
    go func() {
        for {
            price, err := FetchBTCPrice()
            if err != nil {
                log.Println("Error fetching price:", err)
            } else {
                CurrentBTCPrice = price
                log.Println("Updated BTC price:", price)
            }

            time.Sleep(10 * time.Second) // пауза 10с
        }
    }()
}

func FetchBTCPrice() (float64, error) {
    url := "https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=USDT"

    resp, err := http.Get(url)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return 0, err
    }


    var parsed map[string]float64

    if err := json.Unmarshal(body, &parsed); err != nil {
        return 0, err
    }

    return parsed["USDT"], nil
}

