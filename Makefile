# All tests are stored in the test/ directory and run using the builtin go tools
tests:
	cd test && go test -v -race

# Simple for now, all configuration happens through config-{dev/prod}.json
run:
	go run main.go

# golangci looks for logical or syntactical errors in all files in sub directories
# we have to specify subdirs for the go fmt formatter
format:
	golangci-lint run
	go fmt .
	go fmt ./model/
	go fmt ./api/
	go fmt ./config/
	go fmt ./test/
