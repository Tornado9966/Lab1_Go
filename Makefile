default: out/example

clean:
	rm -rf out

test:
	go vet
	go test
	
out/example: implementation.go cmd/example/main.go
	mkdir -p out
	echo package main > cmd/example/buildVersion.go
	echo -n "var buildVersion = \"" >> cmd/example/buildVersion.go
	echo -n "`git describe`" >> cmd/example/buildVersion.go
	echo "\"" >> cmd/example/buildVersion.go
	go build -o out/example ./cmd/example
