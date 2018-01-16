.PHONY: all clean deps build public-ip internal-ip uptime load-average temperature battery datetime volume install dir copy

all: clean build

clean:
	rm -f public-ip internal-ip uptime load-average temperature battery datetime volume

deps:
	go get -t ./...

build: public-ip internal-ip uptime load-average temperature battery datetime volume

public-ip:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' ./cmd/public-ip/

internal-ip:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' ./cmd/internal-ip/

uptime:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' ./cmd/uptime/

load-average:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' ./cmd/load-average/

temperature:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' ./cmd/temperature/

battery:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' ./cmd/battery/

datetime:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' ./cmd/datetime/

volume:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' ./cmd/volume/

install: build dir copy

dir:
	mkdir -p ~/.config/i3blocks-go

copy:
	cp public-ip ~/.config/i3blocks-go/
	cp internal-ip ~/.config/i3blocks-go/
	cp uptime ~/.config/i3blocks-go/
	cp load-average ~/.config/i3blocks-go/
	cp temperature ~/.config/i3blocks-go/
	cp battery ~/.config/i3blocks-go/
	cp datetime $GOPATH/bin/i3blocks-datetime
	cp volume ~/.config/i3blocks-go/
