package main

import (
	"dns-resolver/message"
	"fmt"
)

func main() {
    fmt.Println("Hello dns-resolver!")
    question := message.NewQuestion("google.com", message.TypeA, message.ClassIN)
    dnsMessage := message.NewMessage([]message.QuestionType{*question})
    dnsMessageBytes := dnsMessage.ToBytes()
    result, _ := dnsMessage.DecodeFromBytes(dnsMessageBytes)
    fmt.Printf("raw dns message: %v \ndns message in bytes: %v", dnsMessage, dnsMessageBytes)
    fmt.Printf("\ndecoded dns message: %v \n", result)
}
