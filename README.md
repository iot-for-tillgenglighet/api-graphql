# api-graphql
A GraphQL API to provide IoT-Hub data to clients

## Deprecation Notice

This repository is deprecated and the code has moved to https://github.com/diwise/api-graphql

# Build and tag with Docker

`docker build -f deployments/Dockerfile -t iot-for-tillgenglighet/api-graphql:latest .`

# Run with Docker

`docker run -it -p 8080:8080 iot-for-tillgenglighet/api-graphql:latest`
