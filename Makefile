BINARY = aports-api
COVER_FILE = coverage.out
COVER_REPORT = coverage.html

.PHONY: db

build:
	go build -v -o $(BINARY)

clean:
	rm -f $(BINARY) $(COVER_FILE) $(COVER_REPORT)

coverage:
	go tool cover -html $(COVER_FILE) -o $(COVER_REPORT)

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

test:
	go test -coverprofile $(COVER_FILE) ./...
