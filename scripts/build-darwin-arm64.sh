./scripts/build-frontend.sh

echo "Building the Go RESTful API application (Mac OS Apple Silicon processors)"
GOOS=darwin GOARCH=arm64 go build -o ./bin/app-1.0.0-darwin-arm64 -tags prod ./backend/cmd