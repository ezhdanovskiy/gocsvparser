ASSETS_DIR = "assets"
build:
	@export GOPATH=$${GOPATH-~/go} && \
	go get github.com/jteeuwen/go-bindata/... github.com/elazarl/go-bindata-assetfs/... && \
	$$GOPATH/bin/go-bindata -o bindata.go -tags builtinassets ${ASSETS_DIR}/... && \
	go build -tags builtinassets -ldflags "-X main.builtinAssets=${ASSETS_DIR}"
