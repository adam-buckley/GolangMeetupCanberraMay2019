### How to set this up

You'll need 3 terminal tabs:
- go into server, run `npm i` then `node index.js`
- on the next tab go into client, run `npm i` then `npm run dev`
- the final tab is used for building the WASM binary, to do that go into server/go-wasm and run `GOOS=js GOARCH=wasm go build -o main.wasm`

Navigate to the URL from step 2, enjoy some dots.