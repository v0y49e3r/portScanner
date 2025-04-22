package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

// y tuong la se quet tung port
// nhung van de la khi goi Dial thi se rat lau cho nen ta can 1 cai moi no ten
//Goroutine ( chay dong bo nhieu thread mot luc)
//chay nhieu concurrent

type Dialer interface {
	// input la network va address
	// method
	Dial(network, address string) (net.Conn, error)
}

type NetDialer struct {
}

func (nd NetDialer) Dial(network, address string) (net.Conn, error) {
	// implement interface in to struct
	return net.Dial(network, address)
}

type Target struct {
	ip     string
	port   int
	wg     *sync.WaitGroup
	dialer Dialer
	// implement for mocktest
	// du dinh se tao 1 file unit_test.go
}

func (target Target) scanPort() {
	// phong truong hop cia scanPort no bi loi
	// cho cai nay de no luon ket thuc
	defer target.wg.Done()

	address := target.ip + ":" + strconv.Itoa(target.port)
	conn, err := target.dialer.Dial("tcp", address)

	if err != nil {
		// fmt.Printf("port:  %d  is closed ❌\n", target.port)
		return
	}
	conn.Close()
	fmt.Printf("Port:  %d  Open ✅ \n", target.port)
}
