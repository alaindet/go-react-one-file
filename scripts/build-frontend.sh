echo "Building the React front end"
cd ./frontend
npm install
npm run build

echo "Moving the compiled assets to /backend/cmd/dist"
rm -rf ../backend/cmd/dist
mv ./dist ../backend/cmd/dist
cd ..