package main

import (
	"fmt"
	"log"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/ezhdanovskiy/gocsvparser/daemon"
)

func main() {
	fmt.Printf("Start csv parser.\n")

	cfg := &daemon.Config{}
	cfg.ListenSpec = "localhost:3000"
	cfg.Db.FilePath = "/Users/evgenii.zhdanovskiy/go/src/bitbucket.org/ezhdanovskiy/csvparser.go/data/dictionary.csv"

	log.Printf("Running with builtin assets.\n")
	cfg.UI.Assets = &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "assets"}

	if err := daemon.Run(cfg); err != nil {
		log.Printf("Error in main(): %v", err)
	}
	fmt.Printf("Stop csv parser.\n")
}
