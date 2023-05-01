default: out/example

clean:
	rm -rf out

test: *_test.go
	go test ./...

out/painter: 
	mkdir -p out
	go build -o out/painter ./cmd/painter
