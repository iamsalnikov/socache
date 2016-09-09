package main

import "flag"

func main() {
	var ip = flag.String("ip", "0.0.0.0", "servser IP")
	var port = flag.String("port", "9099", "servser port")

	flag.Parse()
	server := NewServer()
	server.Run(*ip, *port)
}
