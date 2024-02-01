# Genezio go server with ts client and go client

## Step 0: Install genezio CLI (branch with go features)

```bash
git clone https://github.com/Genez-io/genezio.git
cd genezio
git checkout go-ast-sdk-demo
npm i && npm run install-locally-dev
```

## Step 1: Deploy the server

```bash
cd go-server && genezio deploy
```

## Step 2: Get the SDK for the typescript client

```bash
cd ts-client && npm add @genezio-sdk/go-server_eu-central-1@1.0.0
```

## Step 3: Run the typescript client

```bash
npm i
tsc index.ts && node index.js
```

## Step 4: Get the SDK for the go client

```bash
cd go-client @@ genezio sdk --language go
```

## Step 5: Run the go client

```bash
go run main.go
```
