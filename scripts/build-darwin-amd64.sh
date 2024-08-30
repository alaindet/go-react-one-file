./scripts/build-frontend.sh

echo "Building the Go RESTful API application (Mac OS Intel processors)"
GOOS=darwin GOARCH=amd64 go build -o ./bin/app-1.0.0-darwin-amd64 -tags prod ./backend/cmd