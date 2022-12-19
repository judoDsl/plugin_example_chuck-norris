all: dependecies test compile
	@echo 'Generatied plugin chuck_norris.'
dependecies:
	@echo 'Resove dependies'
	go mod tidy
test: dependecies
	go test -v
compile: test
	go build -buildmode=plugin -o bin/chuck_norris.so main.go
