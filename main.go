package main

// https://gobyexample.com/goroutines
import (
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	targetIP := "192.168.0.103"
	startPort := 0
	lastPort := 65535
	for port := startPort; port <= lastPort; port++ {
		target := Target{
			ip:     targetIP,
			port:   port,
			wg:     &waitGroup,
			dialer: NetDialer{},
		}
		waitGroup.Add(1)
		go target.scanPort()
	}
	waitGroup.Wait()
}
