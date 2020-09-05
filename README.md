# Basic Go from Docker

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
```
