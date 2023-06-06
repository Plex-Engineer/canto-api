# Canto API 
[![Build and test](https://github.com/Plex-Engineer/canto-api/actions/workflows/go.yml/badge.svg)](https://github.com/Plex-Engineer/canto-api/actions/workflows/go.yml)

Open-source backend for efficiently serving Canto data built using [Redis](https://github.com/redis/redis), [Fiber](https://github.com/gofiber/fiber) and [Go](https://github.com/golang/go). Built to minimize load on nodes to allow applications to scale better. 

# WIP 🚧

## Quickstart
```bash
# clone repo
git clone git@github.com:Plex-Engineer/canto-api.git

# build binary
cd canto-api
go build

# run redis 
redis-server

# run binary
./canto-api
