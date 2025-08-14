package main

import (
	"dns-resolver/process"
	"fmt"
	"os"
	"time"
)

func main() {
    start := time.Now()

    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <domain>")
        os.Exit(1)
    }

    query := os.Args[1]
    name, data := process.Process(query)

    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Address: %s\n", data)

    fmt.Println("\nexecution time: ", time.Since(start))
}
