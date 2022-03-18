W:=${W}
V:=${V}
NAME=micro
IMAGE_NAME=micro/$(NAME)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --abbrev=0 --tags --always --match "v*")
GIT_IMPORT=github.com/matchstalk/mss-boot
CGO_ENABLED=0
BUILD_DATE=$(shell date +%s)
LDFLAGS=-X $(GIT_IMPORT).BuildDate=$(BUILD_DATE) -X $(GIT_IMPORT).GitCommit=$(GIT_COMMIT) -X $(GIT_IMPORT).GitTag=$(GIT_TAG)
IMAGE_TAG=$(GIT_TAG)-$(GIT_COMMIT)
PROTO_FLAGS=--go_opt=paths=source_relative
PROTO_PATH=$(GOPATH)/src:.
SRC_DIR=$(GOPATH)/src

.PHONY: proto
proto:
	#protoc-go-inject-tag -I ./proto -I ${GOPATH}/src  --go_out=plugins=grpc: proto/${W}/${V}/*;
	find proto/ -name '*.proto' -exec protoc --proto_path=$(PROTO_PATH) $(PROTO_FLAGS) --go_out=plugins=grpc:. {} \;


.PHONY: lint
lint:
	golangci-lint run -v ./service/${W}/...;
