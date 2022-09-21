# load env vars
-include .env
export mongoUri := $(value mongoUri)
export AUTH0_DOMAIN := $(value AUTH0_DOMAIN)
export AUTH0_CLIENT_ID := $(value AUTH0_CLIENT_ID)

# proto generates code from the most recent proto file(s)
.PHONY: proto
proto:
	cd proto && buf mod update
	buf generate
	buf build
	cd proto && buf push

.PHONY: rungo
rungo:
	go run main.go

.PHONY: run
run:
	dapr run \
		--app-id patient \
		--app-port 8080 \
		--app-protocol grpc \
		--config ./.dapr/config.yaml \
		--components-path ./.dapr/components \
		go run .

.PHONY: kill
kill:
	lsof -P -i TCP -s TCP:LISTEN | grep 8080 | awk '{print $2}' | { read pid; kill -9 ${pid}; }
	lsof -P -i TCP -s TCP:LISTEN | grep 9090 | awk '{print $2}' | { read pid; kill -9 ${pid}; }

.PHONY: test
test:
	go test -v ./handler/...
