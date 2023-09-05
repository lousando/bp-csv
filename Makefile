.PHONY: install

./bind/bp-csv: main.go
	go build -o ./bin/bp-csv

install: main.go
	go install
