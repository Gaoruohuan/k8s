export tag=v1.0

build:
	echo "building httpserver binary"
	mkdir -p bin/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/httpserver .

release: build
	echo "building httpserver container"
	docker build -t grh9611/httpserver:${tag} .

push: release
	echo "pushing grh9611/httpserver"
	docker push grh9611/httpserver:${tag}