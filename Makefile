BINARY = aports-api

.PHONY: db

build:
	go build -o $(BINARY)

clean:
	rm $(BINARY)

db:
	docker-compose up -d db

db-shell:
	docker-compose run --rm db psql -U postgres -h db aports

down:
	docker-compose down

lint:
	golint -set_exit_status ./...

run:
	go run main.go
