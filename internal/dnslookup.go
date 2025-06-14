package internal

import (
    "fmt"
    "net"
    "sort"
    "strings"
    "sync"
    "time"
)

var commonPorts = []int{21, 22, 23, 25, 53, 80, 110, 143, 443, 3306, 8080}

func scanPort(host string, port int, wg *sync.WaitGroup, results chan int) {
    defer wg.Done()
    address := fmt.Sprintf("%s:%d", host, port)
    conn, err := net.DialTimeout("tcp", address, 2*time.Second)
    if err == nil {
        conn.Close()
        results <- port
    }
}

func PortScan(host string, ports []int) {
    var wg sync.WaitGroup
    results := make(chan int, len(ports))

    fmt.Println("🔍 Scanning ports...")

    for _, port := range ports {
        wg.Add(1)
        go scanPort(host, port, &wg, results)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    var openPorts []int
    for port := range results {
        openPorts = append(openPorts, port)
    }

    sort.Ints(openPorts)
    if len(openPorts) == 0 {
        fmt.Println("❌ No open ports.")
    } else {
        fmt.Println("🚪 Open Ports:")
        for _, port := range openPorts {
            fmt.Println("➤", port)
        }
    }
}

func StartPortScanner() {
    var host, choice, portStr string
    fmt.Print("Enter host: ")
    fmt.Scanln(&host)

    fmt.Print("Use default ports? (y/n): ")
    fmt.Scanln(&choice)

    if strings.ToLower(choice) == "y" {
        PortScan(host, commonPorts)
    } else {
        fmt.Print("Enter ports (comma-separated): ")
        fmt.Scanln(&portStr)
        var custom []int
        for _, p := range strings.Split(portStr, ",") {
            var port int
            fmt.Sscanf(strings.TrimSpace(p), "%d", &port)
            if port > 0 {
                custom = append(custom, port)
            }
        }
        PortScan(host, custom)
    }
}
