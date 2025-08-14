package main

import (
	"dns-resolver/http"
	"dns-resolver/message"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
    start := time.Now()

    if len(os.Args[1]) == 0 {
        log.Fatal("you must provide a url")
    }

    query := os.Args[1]

    for true {
        dnsMessage := message.NewMessage(query)
        plainResult, err := http.DoRequest(dnsMessage.ToBytes())
        if err != nil {
            log.Fatalf("error: %s", err)
            return
        }

        result, _ := dnsMessage.DnsMessageFromBytes(plainResult)

        if result.HasError() {
            errorMap := []string{"no error", "Format Error", "Server Failure", "Name Error", "Not Implemented", "Refused"}
            fmt.Printf("\n error code: %s\n", errorMap[result.Header.Flags.RCode])
            break
        }

        answer := result.Answers[0]
        if answer.Type == message.TypeCNAME { 
            query = answer.RDataParsed
            continue
        }

        fmt.Println("\n**********************************")
        fmt.Printf("Name: %s\n", answer.Name)
        fmt.Printf("Address: %s\n", answer.RDataParsed)
        fmt.Println("**********************************")
        break
    }

    fmt.Println("\nexecution time: ", time.Since(start))
}
