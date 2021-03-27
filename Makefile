BINARY := open-song-remote-display

build:
	go build -o $(BINARY) ./cmd/main.go

clean:
	rm -f $(BINARY)
