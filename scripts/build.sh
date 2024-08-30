./scripts/build-frontend.sh

echo "Building the Go RESTful API application (current platform)"
go build -o ./bin/app-1.0.0 -tags prod ./backend/cmd