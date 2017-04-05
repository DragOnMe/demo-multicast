package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"syscall"
	"time"
)

var names = make(map[string]time.Time)

func main() {
	name, ok := os.LookupEnv("HOSTNAME")
	if !ok {
		name = "sample-hostname"
	}

	addr, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
	if err != nil {
		panic(err)
	}

	go server(name, addr)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		if sig == syscall.SIGINT || sig == syscall.SIGTERM {
			done <- true
		}
	}()

	pinger := time.NewTicker(time.Second * 2)
	go func() {
		for _ = range pinger.C {
			ping(name, addr)
		}
	}()

	reaper := time.NewTicker(time.Second * 3)
	go func() {
		for _ = range reaper.C {
			fmt.Println("Reaping ...")
			for n, t := range names {
				if time.Now().After(t.Add(5 * time.Second)) {
					fmt.Println("Peer", name, "timed out.")
					delete(names, n)
				}
			}
			fmt.Println("... done.")
			keys := []string{}
			for k := range names {
				keys = append(keys, k)
			}
			fmt.Println("Known peers alive are:", strings.Join(keys, ", "))
		}
	}()

	<-done
}

func server(name string, address *net.UDPAddr) {
	l, err := net.ListenMulticastUDP("udp", nil, address)
	if err != nil {
		panic(err)
	}

	l.SetReadBuffer(8192)

	for {
		b := make([]byte, 8192)
		len, _, err := l.ReadFromUDP(b)

		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		} else {
			n := string(b[:len])

			if n != name {
				if _, ok := names[n]; !ok {
					fmt.Println("Peer", n, "was discovered.")
					ping(name, address)
				}
				names[n] = time.Now()
			}
		}
	}
}

func ping(name string, address *net.UDPAddr) {
	c, err := net.DialUDP("udp", nil, address)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sending ping")

	c.Write([]byte(name))
}
