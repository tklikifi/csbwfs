package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print("[" + time.Now().Format("2006-01-02T15:04:05.999") + "] " + string(bytes))
}

func scan(ip_address string, first_port int, last_port int) {
	ips, err := net.LookupIP(ip_address)
	if err != nil {
		log.Println(err)
		return
	}
	ip := ips[0]
	fmt.Println("Scanning IP address:", ip)
	for p := min(first_port, last_port); p <= max(first_port, last_port); p += 1 {
		dst := fmt.Sprintf("%s:%d", ip, p)
		conn, err := net.DialTimeout("tcp", dst, 1000*time.Millisecond)
		if err != nil {
			// log.Println(err)
			continue
		}
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("%s:%d -- %s", ip, p, message)
	}
}

func main() {
	// Set log formatting
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	// Parse command line arguments
	flag.Usage = func() {
		fmt.Printf("Usage: portscan [options] ip [ip ...]\n\n")
		fmt.Printf("Perform port scan against given IP address.\n\n")
		flag.PrintDefaults()
	}
	flag_f := flag.Int("f", 1024, "First port")
	flag_l := flag.Int("l", 49151, "Last port")
	flag.Parse()
	ip_addresses := flag.Args()

	// Scan IP addresses
	for _, ip := range ip_addresses {
		scan(ip, *flag_f, *flag_l)
	}
}
