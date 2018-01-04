package daemon

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ezhdanovskiy/gocsvparser/ui"
)

func Run() error {
	address := "127.0.0.1:7878"
	log.Printf("Starting, HTTP on: %s\n", address)

	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Printf("Error creating listener: %v\n", err)
		return err
	}

	ui.Start(l)

	waitForSignal()

	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	log.Printf("Got signal: %v, exiting.", s)
}
