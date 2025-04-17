# Go/React in one file

This repository demonstrates how to embed a React application and a Go web server in a single binary file. The application is named **YATA** for "Yet Another Todo Application" and it's a simple React SPA coupled to a Go RESTful API with an in-memory database and pre-seeded mock data. You can optionally load and dump a simple JSON file as your database.

This is not production-ready and is somewhat impractical for most real-world uses, but a modern frontend application embedded in a single Go binary like this is just one file to deploy and can effectively serve as a simpler alternative for small applications.

The final bundle is an executable binary weighting **~8 Mb** that can be easily built for Windows, Mac or Linux from any platform just by changing the build flags in the `./build.sh` script (thank you Go team). A comparable lightweight Docker image or Electron application could easily weight 10-20 times more.

## Requirements

- Go 1.22+
- Node 20+

## Demo

To see this in action, run these commands

```shell
./scripts/build.sh
./bin/app
```

## Flags

- `--port` accepts a number to declare a different port for the web server
- `--jsondb` accepts a path to optionally read and write a JSON file with that path, instead of using an in-memory database

## Development

- Build the front end first to allow the Go app to compile and embed it

  ```shell
  ./scripts/build-frontend.sh
  ```

- Start the backend in live reload mode on port 8080 with WGO (https://github.com/bokwoon95/wgo)

  ```shell
  wgo run -verbose \
    -xdir=frontend \
    -xdir=docs \
    -xdir=scripts \
    -xdir=. \
    -dir=backend \
    ./backend/cmd --port=8080
  ```

- Start the backend without live reloading on port 8080

  ```shell
  go run ./backend/cmd --port=8080
  ```

- On another terminal, start the front end Vite development server on port 5173
  ```shell
  npm run --prefix=frontend dev
  ```

## Build

Run one of these scripts

- `./scripts/build.sh`: Builds for your current platform
- `./scripts/build-darwin-amd64.sh`: Builds for Mac OS with Intel processors
- `./scripts/build-darwin-arm64.sh`: Builds for Mac OS with Apple Silicon processors
- `./scripts/build-linux-amd64.sh`: Builds for Linux 64 bit
- `./scripts/build-win-amd64.sh`: Builds for Windows 64bit
- `./scripts/build-all.sh`: Builds for all the above platforms

### Resources

- https://github.com/bokwoon95/wgo
- https://marketplace.visualstudio.com/items?itemName=humao.rest-client
