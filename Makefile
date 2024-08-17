build:
	go build -o bin/main ./cmd/cnss/main.go 

run:
	./bin/main 

remove: 
	rm -i ./bin/main
