package main

import (
	"dns-resolver/message"
	"fmt"
)

func main() {
    fmt.Println("Hello dns-resolver!")
    dnsMessage := message.NewMessage("google.com", message.TypeA)
    fmt.Printf("raw dns message: %v \ndns message in bytes: %v", dnsMessage, dnsMessage.ToBytes())
}
