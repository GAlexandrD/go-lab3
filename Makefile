default: out/example

clean:
	rm -rf out

test:
	go test ./...

out/painter: cmd/painter/main.go
	mkdir -p out
	go build -o out/painter ./cmd/painter
