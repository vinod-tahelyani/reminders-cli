buildClient:
	@echo "Building client binary"
	go build -o bin/client/client ./cmd/client
buildServer:
	@echo "Building server binary"
	go build -o bin/server/server ./cmd/server