package internal

import (
    "bufio"
    "fmt"
    "net"
    "net/http"
    "os/exec"
    "strings"
    "time"
)

func RunWhois(domain string) {
    fmt.Println("📄 WHOIS Info:")
    out, err := exec.Command("whois", domain).Output()
    if err != nil {
        fmt.Println("❌ 'whois' not found. Please install it or try using online lookup.")
        return
    }

    scanner := bufio.NewScanner(strings.NewReader(string(out)))
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "Registrar") || strings.Contains(line, "OrgName") {
            fmt.Println("➤", line)
        }
    }
}

func GrabHTTPHeaders(host string) {
    urls := []string{"http://" + host, "https://" + host}
    for _, url := range urls {
        client := http.Client{Timeout: 3 * time.Second}
        resp, err := client.Get(url)
        if err != nil {
            fmt.Println("❌", url, "unreachable")
            continue
        }
        defer resp.Body.Close()

        fmt.Println("🌐 Headers from:", url)
        for key, val := range resp.Header {
            fmt.Println("➤", key+":", strings.Join(val, ", "))
        }
        break
    }
}

func RunHostInfo(target string) {
    ip := target
    if net.ParseIP(target) == nil {
        ips, err := net.LookupIP(target)
        if err != nil {
            fmt.Println("❌ Can't resolve domain.")
            return
        }
        ip = ips[0].String()
    }

    RunWhois(target)
    GrabHTTPHeaders(target)
    fmt.Println("📡 IP Info:")
    LookupIP(ip)
}
