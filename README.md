# Going Live

## What is this?

*Going Live* is a super basic containerized starter for easy development of APIs in Go.  
By leveraging the superpowers of the well-know nodemon and the unmatched speed of the Go compiler, the *Going Live* container is able to rebuild the binary and restart the server every time you save some changes or add new files.

## Requirements
- Docker
- Docker Compose

## Usage

Before First Run Only:
```console
cat .env.local >> .env
```

Run:
```console
docker-compose up
```

Run after the Dockerfile is changed:
```console
docker-compose up --build
```

Try:
```console
curl localhost:3001/titles | jq
curl localhost:3001/titles/1 | jq
```

## Authors
- [advance467](mailto:advance467@hotmail.com)
- [alaindet](https://github.com/alaindet)
