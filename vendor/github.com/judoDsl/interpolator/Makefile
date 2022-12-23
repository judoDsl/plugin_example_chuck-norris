fmt:
	@echo "==> Formating code"
	go fmt

test: fmt
	@echo "==> Running tests"
	GO111MODULE=on go test -v

test-cover: fmt
	@echo "==> Running Tests with coverage"
	GO111MODULE=on go test -cover .


