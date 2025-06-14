package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "bugscanner/internal"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("═════════════════════════════════════════")
        fmt.Println("      🐞 TS HACKE - CLI TOOL (v1.0.0)")
        fmt.Println("═════════════════════════════════════════")
        fmt.Println("[1] Host Scanner")
        fmt.Println("[2] Subfinder")
        fmt.Println("[3] IP Lookup")
        fmt.Println("[4] File Toolkit")
        fmt.Println("[5] Port Scanner")
        fmt.Println("[6] DNS Record Lookup")
        fmt.Println("[7] Host Info")
        fmt.Println("[8] Help")
        fmt.Println("[9] Update")
        fmt.Println("[0] Exit")
        fmt.Print("Enter your choice: ")

        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)

        switch choice {
        case "1":
            fmt.Println("🔍 HOST SCANNER")
            fmt.Print("Enter URL (with http/https): ")
            url, _ := reader.ReadString('\n')
            internal.ScanHost(strings.TrimSpace(url))
        case "2":
            internal.RunSubfinder()
        case "3":
            fmt.Print("Enter IP or domain: ")
            input, _ := reader.ReadString('\n')
            internal.LookupIP(strings.TrimSpace(input))
        case "4":
            internal.FileToolkitMenu()
        case "5":
            fmt.Println("🚪 PORT SCANNER")
            internal.StartPortScanner()
        case "6":
            fmt.Println("🌐 DNS RECORD LOOKUP")
            fmt.Print("Enter domain: ")
            domain, _ := reader.ReadString('\n')
            internal.LookupDNS(strings.TrimSpace(domain))
        case "7":
            fmt.Println("🧠 HOST INFO")
            fmt.Print("Enter domain or IP: ")
            input, _ := reader.ReadString('\n')
            internal.RunHostInfo(strings.TrimSpace(input))
        case "8":
            internal.ShowHelp()
        case "9":
            internal.CheckForUpdate()
        case "0":
            fmt.Println("👋 Are you sure you want to exit? (y/n): ")
            confirm, _ := reader.ReadString('\n')
            if strings.TrimSpace(confirm) == "y" {
                fmt.Println("🛑 Exiting BugScanner. Goodbye!")
                os.Exit(0)
            }
        default:
            fmt.Println("❌ Invalid choice.")
        }

        fmt.Println()
    }
}
