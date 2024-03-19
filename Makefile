VER := $(shell git describe --always --long)-$(shell date -u +%Y-%m-%dT%H-%M-%S%Z)

start:
			air -c .air.toml

test:
			go test ./...

image-build:
			podman build -t webapp:$(VER) -f Containerfile .

ver:
	@echo $(VER)

image:		image-build ver

