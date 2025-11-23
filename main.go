package main

import "sync"

type program struct {
	rtspPort     int
	rtpPort      int
	rtcpPort     int
	mutex        sync.Mutex
	rtspl        *rtspListener
	rtpl         *udpListener
	rtcpl        *udpListener
	clients      map[*rtspClient]struct{}
	streamAuthor *rtspClient
	streamSdp    []byte
}
