# service := $(shell basename $(PWD))

-include .env
export MONGO_URI := $(value MONGO_URI)
export PORT := $(value PORT)

# proto generates code from the most recent proto file(s)
.PHONY: proto
proto:
	cd proto && buf mod update
	buf lint
	# buf breaking --against './.git#branch=main,ref=HEAD~1'
	buf build
	buf generate
	cd proto && buf push
	
.PHONY: run
run:
	dapr run \
		--app-id patient \
		--app-port 8080 \
		--app-protocol http \
	
		go run .

.PHONY: kill
kill:
	-lsof -P -i TCP -s TCP:LISTEN | grep 8080 | awk '{print $$2}' | { read pid; kill -9 $$pid; }
	-lsof -P -i TCP -s TCP:LISTEN | grep 9090 | awk '{print $$2}' | { read pid; kill -9 $$pid; }

.PHONY: test
test:
	go test -v ./handler/...

# ------------------------------------------------------------
# Unit Testing
# - start the dapr server using 'make publisher'
# - publish events using 'make publish'
# ------------------------------------------------------------

.PHONY: publisher
publisher:
	dapr run \
		--app-id patient \
		--app-protocol http \
		

.PHONY: publish
publish:
	dapr publish \
		--publish-app-id patient \
		
	
		--data-file ./.json/events/person.created.json

# clear redis queue
.PHONY: flushq
flushq:
	docker exec -it dapr_redis redis-cli FLUSHALL