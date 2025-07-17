package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

const (
	// defaultNtpPort is the default port for NTP servers.
	// this server will be used if no other server is specified.

	ntpPacketSize = 48

	// ntpDelta is the number of seconds between the NTP epoch (1900) and the Unix epoch (1970).
	ntpDelta = 2208988800
)

func main() {
	// parse the command line flags
	serverPtr := flag.String("server", "pool.ntp.org", "NTP server to query (e.g., time.google.com or pool.ntp.org)")

	// parse CLI flags
	flag.Parse()

	// get the NTP Server Address from parsed flag value
	ntpServer := *serverPtr

	// ensure server address includes a port
	if _, _, err := net.SplitHostPort(ntpServer); err != nil {
		ntpServer = net.JoinHostPort(ntpServer, "123")
	}

	// network setup
	addr, err := net.ResolveUDPAddr("udp", ntpServer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving address '%s': %v\n", ntpServer, err)
		os.Exit(1) // exit with non-zero error code
	}

	// establish UDP connection
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error dialing UDP address '%s': %v\n", ntpServer, err)
		os.Exit(1) // exit with non-zero error code
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(5 * time.Second)) // max 5 seconds to wait for a response

	requestPacket := make([]byte, ntpPacketSize) // craft NTP request packet

	requestPacket[0] = 0x1B // NTP version 3, client mode

	t1 := time.Now()

	_, err = conn.Write(requestPacket)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending request to '%s': %v\n", ntpServer, err)
		os.Exit(1) // exit with non-zero error code
	}

	responsePacket := make([]byte, ntpPacketSize)

	_, err = conn.Read(responsePacket)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Fprintf(os.Stderr, "Timeout waiting for response from '%s'\n", ntpServer)
		} else {
			fmt.Fprintf(os.Stderr, "Error reading response from '%s': %v\n", ntpServer, err)
		}

		os.Exit(1) // exit with non-zero error code
	}

	t4 := time.Now() // get current time to calculate the round trip delay

	txTimeInt := binary.BigEndian.Uint32(responsePacket[40:44])  // extract transmit timestamp from response
	txTimeFrac := binary.BigEndian.Uint32(responsePacket[44:48]) // extract transmit timestamp fraction

	serverUnixTime := float64(txTimeInt) + (float64(txTimeFrac) / (1 << 32)) - ntpDelta                        // convert NTP time to Unix time
	serverTime := time.Unix(int64(serverUnixTime), int64((serverUnixTime-float64(int64(serverUnixTime)))*1e9)) // convert to time.Time

	offset := serverTime.Sub(t4).Seconds()

	delay := t4.Sub(t1).Seconds()

	fmt.Printf("\n--- NTP Test Results ---\n")
	fmt.Printf("NTP Server: %s\n", ntpServer)
	fmt.Printf("Server Time: %s\n", serverTime.Format(time.RFC3339Nano))
	fmt.Printf("Offset: %.6f seconds\n", offset)
	fmt.Printf("Round Trip Delay: %.6f seconds\n", delay)
}
