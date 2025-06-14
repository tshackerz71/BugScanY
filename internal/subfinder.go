
package internal

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func RunSubfinder() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter domain for subfinder: ")
    domain, _ := reader.ReadString('\n')
    domain = strings.TrimSpace(domain)

    if domain == "" {
        fmt.Println("❌ Domain is required.")
        return
    }

    fmt.Println("🔍 Running Subfinder for:", domain)
    cmd := exec.Command("subfinder", "-d", domain)
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("❌ Error running subfinder:", err)
        return
    }
    fmt.Println(string(output))
}
