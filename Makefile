# Create a make rule that compiles main.go into main and then runs it.

build: main.go
	go build -o main main.go
	./main