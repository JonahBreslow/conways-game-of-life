# Create a make rule that compiles main.go into main and then runs it.

build: 
	go build -o src/main src/main.go
	./src/main