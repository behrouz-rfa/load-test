# ==============================================================================
# Main

run:
	go run ./cmd/app/main.go -to 4000 -d 5 -c 5 -u https://jsonplaceholder.typicode.com/todos/1

build:
	go build ./cmd/app/main.go




# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod download

tidy:
	go mod tidy
	go mod download

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod download

deps-cleancache:
	go clean -modcache



# ==============================================================================
# Go migrate postgresql

# ==============================================================================
# Docker compose commands

develop:
	echo "Starting docker environment"
	docker compose -f docker-compose.yml up --build

local:
	echo "Starting local environment"
	docker compose -f docker-compose.yml up --build


# ==============================================================================
# Docker support

FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)