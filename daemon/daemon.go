package daemon

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/ezhdanovskiy/gocsvparser/db"
	"github.com/ezhdanovskiy/gocsvparser/model"
	"github.com/ezhdanovskiy/gocsvparser/ui"
)

type Config struct {
	ListenSpec string
	Db         db.Config
	UI         ui.Config
}

func Run(cfg *Config) error {
	log.Printf("Starting, HTTP on: %s\n", cfg.ListenSpec)

	l, err := net.Listen("tcp", cfg.ListenSpec)
	if err != nil {
		log.Printf("Error creating listener: %v\n", err)
		return err
	}

	db, err := db.InitDb(cfg.Db)
	if err != nil {
		log.Printf("Error initializing database: %v\n", err)
		return err
	}

	m := model.New(db)

	ui.Start(cfg.UI, m, l)

	url := "http://" + cfg.ListenSpec
	log.Printf("Open %s\n", url)
	err = openBrowser(url)
	if err != nil {
		log.Printf("Error open %s: %v\n", url, err)
		return err
	}

	waitForSignal()

	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	log.Printf("Got signal: %v, exiting.", s)
}

func openBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}
