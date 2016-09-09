package main

import "flag"

var ip = flag.String("ip", "0.0.0.0", "servser IP")
var port = flag.String("port", "9099", "servser port")

func main() {
	flag.Parse()
	server := NewServer()
	server.Run(*ip, *port)
}
