
outfile = ./bin/define

files = ./cmd/define.go

run:
	go run $(files)

build:
	go build -o $(outfile) $(files)

test:
	go test ./lib -v -coverprofile=coverage.out -bench=.
