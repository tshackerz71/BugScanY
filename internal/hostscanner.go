package internal

import (
    "fmt"
    "net/http"
    "time"
)

func ScanHost(url string) {
    client := http.Client{Timeout: 5 * time.Second}
    resp, err := client.Get(url)
    if err != nil {
        fmt.Println("❌ Host not reachable:", err)
        return
    }
    defer resp.Body.Close()

    fmt.Println("✅ Status:", resp.Status)
    for key, val := range resp.Header {
        fmt.Printf("  %s: %s\n", key, val)
    }
}
