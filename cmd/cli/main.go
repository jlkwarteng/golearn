package main

import (
	"fmt"
	"golearn/internal/network"
)

func main() {
	fmt.Println("Application Main Entry Point")
	network.Ping("127.0.0.1")
}
