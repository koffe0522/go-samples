.PHONY: start
## start: build & run
start:
	docker-compose up -d --build
	@echo "Building..."
	@echo "http://localhost:8000/"

.PHONY: stop
## build: stop the application
stop:
	docker-compose down

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
