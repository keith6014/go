PROJECT:=dbExample-goa
export GOPATH=/home/user/go
export GOROOT=/opt/go
GO=/opt/go
GOAGEN:=${GOPATH}/bin/goagen

clean:
	rm -rf build deps
deps: design/design.go
	${GOAGEN} bootstrap -d ${PROJECT}/design
	>$@
build: deps
	${GO}/bin/go build -i
	>$@
