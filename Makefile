build:
	go build -o texture-list-generator

generate_test:
	go generate ./...

test: generate_test
	gotestsum -- ./...