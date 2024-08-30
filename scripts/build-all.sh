./scripts/build-frontend.sh

echo "Building the Go RESTful API application (Mac OS Intel processors)"
GOOS=darwin GOARCH=amd64 go build -o ./bin/app-1.0.0-darwin-amd64 -tags prod ./backend/cmd

echo "Building the Go RESTful API application (Mac OS Apple Silicon processors)"
GOOS=darwin GOARCH=arm64 go build -o ./bin/app-1.0.0-darwin-arm64 -tags prod ./backend/cmd

echo "Building the Go RESTful API application (Linux 64 bit)"
GOOS=linux GOARCH=amd64 go build -o ./bin/app-1.0.0-linux-amd64 -tags prod ./backend/cmd

echo "Building the Go RESTful API application (Windows 64 bit)"
GOOS=windows GOARCH=amd64 go build -o ./bin/app-1.0.0-win-amd64.exe -tags prod ./backend/cmd