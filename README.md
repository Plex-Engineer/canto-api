# Canto API 
[![Build and test](https://github.com/Plex-Engineer/canto-api/actions/workflows/go.yml/badge.svg)](https://github.com/Plex-Engineer/canto-api/actions/workflows/go.yml)

Open-source backend for efficiently serving Canto data built using [Redis](https://github.com/redis/redis), [Fiber](https://github.com/gofiber/fiber) and [Go](https://github.com/golang/go). Built to minimize load on nodes to allow applications to scale better. 

## Dependencies
- `golang 1.18` or above
- `redis 7.0` ([install here](https://redis.io/docs/getting-started/installation/))
## Quickstart

```bash
CANTO_MAINNET_RPC_URL = 
PORT = :3000
DB_HOST = localhost
DB_PORT = 6379
CANTO_MAINNET_GRPC_URL = 
MULTICALL_ADDRESS = 
QUERY_INTERVAL = 3
```

```bash
# clone repo
git clone git@github.com:Plex-Engineer/canto-api.git

# create .env file and set variables:
mv .env_example .env

# build binary
cd canto-api
go build

# run redis 
redis-server

# run binary
./canto-api
```

## Docker
Use docker compose:

`docker compose up -d`