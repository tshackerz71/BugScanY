package internal

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func FileToolkitMenu() {
    fmt.Print("Enter file path: ")
    var path string
    fmt.Scanln(&path)

    file, err := os.Open(path)
    if err != nil {
        fmt.Println("❌ Error:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    seen := make(map[string]bool)
    var lines []string

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if !seen[line] && line != "" {
            seen[line] = true
            lines = append(lines, line)
        }
    }

    fmt.Printf("✅ Unique Lines: %d\n", len(lines))
    for _, l := range lines {
        fmt.Println("➤", l)
    }
}
