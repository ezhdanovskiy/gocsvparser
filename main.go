package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ezhdanovskiy/gocsvparser/daemon"
)

func main() {
	fmt.Printf("Start csv parser. 4\n")

	cfg := &daemon.Config{}
	cfg.ListenSpec = "localhost:3000"
	cfg.UI.Assets = http.Dir("/Users/evgenii.zhdanovskiy/go/src/github.com/ezhdanovskiy/gocsvparser/assets/")

	if err := daemon.Run(cfg); err != nil {
		log.Printf("Error in main(): %v", err)
	}
	fmt.Printf("Stop csv parser.\n")
}
