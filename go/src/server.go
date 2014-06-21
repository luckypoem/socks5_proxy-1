package main

import (
	"secureconn"
	"socks5proxy"
)

var (
	Key = []byte{
		102, 57, 31, 13, 11, 131, 64, 191,
		211, 221, 171, 121, 176, 173, 205, 1,
		61, 5, 3, 7, 19, 23, 41, 37,
		53, 61, 71, 91, 83, 99, 100}
)

func main() {
	sps := &socks5proxy.Socks5ProxyServer{}
	sps.ListenAndServe("tcp", ":1080", secureconn.RC4, Key)
}