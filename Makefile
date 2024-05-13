SHELL=/bin/bash
COVERAGE_THRESHOLD=80
COVERAGE_FILE=coverage.out

start:
	go run .

start.watch:
	air -d

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" \
    -o main ./cmd/app

start.prod:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" \
    -o main ./cmd/app && ./main

lint:
	@echo "FORMATTING"
	go fmt ./...
	@echo "LINTING: golangci-lint"
	golangci-lint run --fix ./...

start.test:
	go test -v ./...

test.cov:
		go test -coverprofile="$(COVERAGE_FILE)" $$( go list ./internal/... | grep services ) >/dev/null && \
			coverage=$$(go tool cover -func="$(COVERAGE_FILE)" | grep total | awk '{print $$3}' | sed 's/%//') && \
			if [ $$(echo "$$coverage < $(COVERAGE_THRESHOLD)" | bc -l) -eq 1 ]; then \
				rm "$(COVERAGE_FILE)"; \
				echo "Cannot push due to insufficient code coverage."; \
				echo "Current coverage: $(COVERAGE_THRESHOLD)% ($$coverage%)"; \
				exit 1; \
			else \
				echo "The code coverage is sufficient: $(COVERAGE_THRESHOLD)% ($$coverage%)"; \
				rm "$(COVERAGE_FILE)"; \
			fi
