package main

import (
	"fmt"
	sys "synf/handlers/data"
	"synf/server"
)

func main() {

	// clear terminal screen ASCII sequence
	fmt.Print("\033[H\033[2J")
	fmt.Print(sys.CpuName())

	server.RawConnect("127.0.0.1", "5000")
	server.RegistrationEndpoint()
	fmt.Println("API STARTED...")
	fmt.Println(server.GetOutboundIp())
	
}
