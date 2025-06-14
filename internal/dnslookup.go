package internal

import (
    "fmt"
    "net"
)

func LookupDNS(domain string) {
    fmt.Println("🔍 DNS Records for:", domain)

    ips, _ := net.LookupHost(domain)
    fmt.Println("🟢 A Records:")
    for _, ip := range ips {
        fmt.Println("➤", ip)
    }

    cname, err := net.LookupCNAME(domain)
    if err == nil {
        fmt.Println("🔁 CNAME:", cname)
    }

    mx, _ := net.LookupMX(domain)
    for _, record := range mx {
        fmt.Println("📩 MX:", record.Host, "Pref:", record.Pref)
    }

    ns, _ := net.LookupNS(domain)
    for _, record := range ns {
        fmt.Println("🌐 NS:", record.Host)
    }

    txt, _ := net.LookupTXT(domain)
    for _, t := range txt {
        fmt.Println("🧾 TXT:", t)
    }
}
