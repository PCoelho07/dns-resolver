package process

import (
	"dns-resolver/http"
	"dns-resolver/message"
	"fmt"
	"log"
)

func Process(query string) (name string, data string) {
    for {
        dnsMessage := message.NewMessage(query)

        fmt.Printf("Querying %s for %s\n\n", message.RootDNS, query)

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

        name = answer.Name
        data = answer.RDataParsed
        break
    }

    return name, data
}
