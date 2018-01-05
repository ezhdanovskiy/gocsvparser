package main

import (
	"fmt"
	"log"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/ezhdanovskiy/gocsvparser/daemon"
)

var builtinAssets string

func main() {
	fmt.Printf("Start csv parser.\n")

	cfg := &daemon.Config{}
	cfg.ListenSpec = "localhost:3000"
	cfg.Db.FilePath = "dictionary.csv"

	assetsPath := "/Users/evgenii.zhdanovskiy/go/src/github.com/ezhdanovskiy/gocsvparser/assets/"
	if builtinAssets != "" {
		log.Printf("Running with builtin assets.\n")
		cfg.UI.Assets = &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: builtinAssets}
	} else {
		log.Printf("Assets served from %q.\n", assetsPath)
		cfg.UI.Assets = http.Dir(assetsPath)
	}

	if err := daemon.Run(cfg); err != nil {
		log.Printf("Error in main(): %v", err)
	}
	fmt.Printf("Stop csv parser.\n")
}
