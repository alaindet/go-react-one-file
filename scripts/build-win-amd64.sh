./scripts/build-frontend.sh

echo "Building the Go RESTful API application (Windows 64 bit)"
GOOS=windows GOARCH=amd64 go build -o ./bin/app-1.0.0-win-amd64.exe -tags prod ./backend/cmd