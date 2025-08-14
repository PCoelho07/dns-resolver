package http

import (
	"dns-resolver/message"
	"fmt"
	"io"
	"net"
	"time"
)

func DoRequest(data []byte) ([]byte, error) {
    addr := fmt.Sprintf("%s:%d", message.RootDNS, message.RootDNSPort)

    c, err := net.Dial("udp", addr)

    if err != nil { 
        return nil, fmt.Errorf("failed to connect to the DNS server: %s", err)
    }

    defer c.Close()

    c.SetDeadline(time.Now().Add(5 * time.Second))

    if _, err := c.Write(data); err != nil {
        return nil, fmt.Errorf("failed send data to server: %s", err)
    }

    buf := make([]byte, 1024)
    n, err := c.Read(buf)

    if err == io.EOF {
        return nil, fmt.Errorf("end of input reached: %s", err)
    }

    if err != nil {
        return nil, fmt.Errorf("failed to read data from server: %s", err)
    }

    return buf[:n], nil
}

