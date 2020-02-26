default: out/example

clean:
	rm -rf out

test:
	go vet
	go test
	
out/example: implementation.go cmd/example/main.go
	mkdir -p out
	echo package main > cmd/example/buildVersion.go & FOR /F %a in ('git describe') do echo var buildVersion = "%a" >> cmd/example/buildVersion.go
	go build -o out/example ./cmd/example