ASSETS_DIR = "assets"
build:
	@export GOPATH=$${GOPATH-~/go} && \
	go get github.com/jteeuwen/go-bindata/... github.com/elazarl/go-bindata-assetfs/... && \
	$$GOPATH/bin/go-bindata -o bindata.go ${ASSETS_DIR}/... && \
	go build
