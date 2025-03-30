package main

import (
	"dns-resolver/message"
	"fmt"
)

func main() {
    fmt.Println("Hello dns-resolver!")
    question := message.NewQuestion("google.com", message.TypeA, message.ClassIN)
    dnsMessage := message.NewMessage([]message.QuestionType{*question})
    fmt.Printf("raw dns message: %v \ndns message in bytes: %v", dnsMessage, dnsMessage.ToBytes())
}
