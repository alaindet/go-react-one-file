# Go/React in one file

This repository demonstrates how to embed a compiled React application in a Go single binary. The application is named "YATA" for "Yet Another Todo Application" and it's a simple React SPA coupled to a Go RESTful API with an in-memory database and pre-seeded mock data.

This is not production-ready and is somewhat impractical for most real-world uses, but a frontend application embedded in a single Go binary like this is far easier to deploy and can effectively serve as a simpler alternative to Electron or Docker for desktop-like web applications and small websites that are hence very easy to deploy.

The final bundle is an executable binary weighting **~8 Mb** that can be easily built for Windows, Mac or Linux from any platform just by changing the build flags in the `./build.sh` script. A comparable lightweight Docker image could easily weight 10-20 times more.

## Requirements

- Go 1.22+
- Node 20+

## Demo

To see this in action, run these commands

```shell
./build.sh
./bin/app

# To run it in another port
./bin/app --port=3333
```

## Development

```shell
# Backend exposed on port 8080 (with wgo installed, see Resources section)
wgo run -verbose -xdir=frontend -xdir=docs -xdir=. -dir=backend ./backend/cmd

# Backend exposed on port 8080 (without wgo installed)
go run ./backend/cmd

# Frontend exposed on port 5173
cd ./frontend && npm install && npm run dev
```

### Resources

- https://github.com/bokwoon95/wgo
- https://marketplace.visualstudio.com/items?itemName=humao.rest-client
