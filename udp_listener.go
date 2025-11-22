package main

import (
	"log"
	"net"
	"sync"
)

type udpListenerCb func(*udpListener, []byte)

type udpListener struct {
	nconn     *net.UDPConn
	logPrefix string
	cb        udpListenerCb
}

func newUdpListener(port int, logPrefix string, cb udpListenerCb) (*udpListener, error) {
	nconn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: port,
	})
	if err != nil {
		return nil, err
	}

	l := &udpListener{
		nconn:     nconn,
		logPrefix: logPrefix,
		cb:        cb,
	}
	l.log("opened on: %d", port)
	return l, nil
}

func (l *udpListener) log(format string, args ...interface{}) {
	log.Printf("["+l.logPrefix+" listener] "+format, args...)
}

func (l *udpListener) run(wg sync.WaitGroup) {
	defer wg.Done()
	buf := make([]byte, 2048)
	for {
		n, _, err := l.nconn.ReadFromUDP(buf)
		if err != nil {
			l.log("read error: %v", err)
			break
		}

		l.cb(l, buf[:n])
	}
}
