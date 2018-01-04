package daemon

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ezhdanovskiy/gocsvparser/ui"
)

type Config struct {
	ListenSpec string
	UI         ui.Config
}

func Run(cfg *Config) error {
	log.Printf("Starting, HTTP on: %s\n", cfg.ListenSpec)

	l, err := net.Listen("tcp", cfg.ListenSpec)
	if err != nil {
		log.Printf("Error creating listener: %v\n", err)
		return err
	}

	ui.Start(cfg.UI, l)

	waitForSignal()

	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	log.Printf("Got signal: %v, exiting.", s)
}
