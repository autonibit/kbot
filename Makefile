APP := $(shell basename $(shell git remote get-url origin))
BUILD_DIR = build
REGISTRY := europe-west3-docker.pkg.dev/devops-kuber-2023/default-repo
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
CURRENT_DIST = $(subst /, , $(word 4, $(shell go version))) # the current dist in format: linux arm64
TARGETOS = $(word 1, $(CURRENT_DIST))
TARGETARCH = $(word 2, $(CURRENT_DIST))

lint: 
	golint

test:
	go test -v

format:
	gofmt -s -w ./

get:
	go get

build: format get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o ./${BUILD_DIR}/kbot${EXT} -ldflags "-X="github.com/autonibit/kbot/cmd.appVersion=${VERSION}

define build_target
	$(MAKE) build TARGETOS=$(1) TARGETARCH=$(if $(2),$(2),amd64) EXT=$(3)
endef

linux: 
	$(call build_target,linux,$(word 2,$(MAKECMDGOALS))) 

macos:
	$(call build_target,darwin,$(word 2, $(MAKECMDGOALS)))

windows:
	$(call build_target,windows,$(word 2, $(MAKECMDGOALS)),.exe)

image: ## Build container image for defaul OS/Arch [linux/amd64]
	docker build . -t ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH} --build-arg TARGETOS=${TARGETOS} --build-arg TARGETARCH=${TARGETARCH}

define image_target
	$(MAKE) image TARGETOS=$(1) TARGETARCH=$(if $(2),$(2),amd64)
endef

image-linux:
	$(call image_target,linux,$(word 2, $(MAKECMDGOALS)))

image-macos:
	$(call image_target,darwin,$(word 2, $(MAKECMDGOALS)))

image-windows:
	$(call image_target,windows,$(word 2, $(MAKECMDGOALS)))

push: 
	docker push ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

clean: 
	rm -rf ./${BUILD_DIR}

clean-image: 
	docker rmi ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH} -f

clean-all: clean clean-image 

%::
	@true