./scripts/build-frontend.sh

echo "Building the Go RESTful API application (Linux 64 bit)"
GOOS=linux GOARCH=amd64 go build -o ./bin/app-1.0.0-linux-amd64 -tags prod ./backend/cmd