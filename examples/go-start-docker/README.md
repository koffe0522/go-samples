# go-start-docker

Basic Go from Docker.
Minimal docker go execution environment.

## set up

```bash
# Duild
$ docker-compose build

# Display image information
$ docker-compose images

# start `-b` is Start in background option
$ docker-compose up -d

# status
$ docker-compose ps

# shell
$ docker-compose exec app /bin/bash

# run
$ docker-compose exec app go run main.go

# down
$ docker-compose down
```
