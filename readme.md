# Going Live

## Introduction:

*Going Live* is a super basic containerized starter for easy development of APIs in Go.  
By leveraging the superpowers of the well-know nodemon and the unmatched speed of the Go compiler, the *Going Live* container is able to rebuild the binary and restart the server everytime you save some changes or add new files.  

___

## Requirements:
- Docker
- Docker Compose

## Usage:

- start the project.

```console
docker-compose up
```

- then run the contents of the initilization_script.sql file against the database using your favorite client (it's exposed to the host on the port 5433).
- now you can try the included demo endpoints like this:

```console
curl localhost:3001/titles | jq
```

```console
curl localhost:3001/titles/1 | jq
```

___
## Now *Go* ahead and start developing your awesome APIs !
___

## Author:
- [advance467](mailto:advance467@hotmail.com)