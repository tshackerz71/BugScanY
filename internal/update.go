package internal

import (
    "encoding/json"
    "fmt"
    "net/http"
)

const (
    currentVersion = "v1.0.0"
    repoAPI        = "https://api.github.com/repos/tshackerz71/bugscanner/releases/latest"
)

func CheckForUpdate() {
    fmt.Println("🔄 Checking for updates...")

    resp, err := http.Get(repoAPI)
    if err != nil {
        fmt.Println("❌ Unable to check for updates:", err)
        return
    }
    defer resp.Body.Close()

    var result struct {
        TagName string `json:"tag_name"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        fmt.Println("❌ Decode error")
        return
    }

    fmt.Println("🛠️ Current Version:", currentVersion)
    fmt.Println("🌐 Latest Version:", result.TagName)

    if result.TagName != currentVersion {
        fmt.Println("🚀 Update available!")
    } else {
        fmt.Println("✅ Up to date.")
    }
}
