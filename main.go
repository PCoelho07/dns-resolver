package main

import (
	"dns-resolver/http"
	"dns-resolver/message"
	"fmt"
	"log"
	"time"
)

func main() {
    start := time.Now()

    dnsMessage := message.NewMessage("google.com")
    fmt.Printf("raw dns message: %+v", dnsMessage)

    dnsMessageParsed := dnsMessage.ToBytes()

    plainResult, err := http.DoRequest(dnsMessageParsed)
    if err != nil {
        log.Fatalf("error: %s", err)
        return
    }

    fmt.Printf("\nresponse from server: %+v \n", plainResult)
    result, _ := message.DnsMessageFromBytes(plainResult)

    if result.HasError() {
        errorMap := []string{"no error", "Format Error", "Server Failure", "Name Error", "Not Implemented", "Refused"}
        fmt.Printf("\n error code: %s\n", errorMap[result.Header.Flags.RCode])
    }

    fmt.Printf("\ndecoded dns message: %+v \n", result)
    fmt.Println("\nexecution time: ", time.Since(start))
}
