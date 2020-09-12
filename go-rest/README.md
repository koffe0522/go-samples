# Go REST

## set up

```bash
# start & build
$ make start

# stop
$ make stop

# prints this help message
$ make list
```

## routes

| Methods | URL                              |      Params       |
| :-----: | :------------------------------- | :---------------: |
|   GET   | http://localhost:8000/books      |                   |
|   GET   | http://localhost:8000/books/{id} |                   |
|  POST   | http://localhost:8000/books      | { title: string } |
