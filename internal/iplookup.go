package internal

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func LookupIP(ip string) {
    resp, err := http.Get("https://ipinfo.io/" + ip + "/json")
    if err != nil {
        fmt.Println("❌ Error:", err)
        return
    }
    defer resp.Body.Close()

    var data map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        fmt.Println("❌ Decode error:", err)
        return
    }

    for k, v := range data {
        fmt.Printf("  %s: %v\n", k, v)
    }
}
