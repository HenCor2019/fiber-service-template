SHELL=/bin/bash
COVERAGE_THRESHOLD=80
COVERAGE_FILE=coverage.out
COVER_MODULES = ./api/users/services/ ./api/tasks/services/ ./api/pokemons/services/

start:
	go run .

start.watch:
	air -d

start.prod:
	go build -o main && ./main

lint:
	@echo "FORMATTING"
	go fmt ./...
	@echo "LINTING: golangci-lint"
	golangci-lint run --fix ./...

start.test:
	go test -v ./...

test.cov:
		go test -coverprofile="$(COVERAGE_FILE)" ./internal/... >/dev/null && \
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
