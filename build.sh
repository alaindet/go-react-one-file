echo "[01/05] Building a demo Go/React app in one file"

echo "[02/05] Building the React front end"
cd ./frontend
npm install
npm run build

echo "[03/05] Moving the compiled assets to /backend/cmd/dist"
rm -rf ../backend/cmd/dist
mv ./dist ../backend/cmd/dist

echo "[04/05] Building the Go application"
cd ..
go build -o ./bin/app -tags prod ./backend/cmd

echo "[05/05] Done"
