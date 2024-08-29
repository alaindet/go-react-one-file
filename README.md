# Go/React in one file

This repository demonstrates how to embed a compiled React application in a Go single binary. The application is named "YATL" for "Yet Another Todo List" and is a simple SPA + RESTful API with an in-memory database, pre-seeded on startup.

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
# Backend exposed on port 8080
cd ./backend/cmd && wgo run .

# Frontend exposed on port 5173
cd ./frontend && npm run dev
```

### Commands

### Resources

- https://github.com/bokwoon95/wgo
- https://marketplace.visualstudio.com/items?itemName=humao.rest-client
