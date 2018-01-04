package main

import (
	"fmt"
	"log"

	"github.com/ezhdanovskiy/gocsvparser/daemon"
)

func main() {
	fmt.Printf("Start csv parser.\n")

	if err := daemon.Run(); err != nil {
		log.Printf("Error in main(): %v", err)
	}
	fmt.Printf("Stop csv parser.\n")
}
