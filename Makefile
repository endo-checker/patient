

# proto generates code from the most recent proto file(s)
.PHONY: proto
proto:
	cd proto && buf mod update
	buf lint
	buf build
	# buf breaking --against './.git#branch=main,ref=HEAD~1'
	buf generate
	cd proto && buf push

.PHONY: run
run:
	go run -tags jwx_es256k .

.PHONY: lint
lint:
	golangci-lint run ./...
	
.PHONY: test
test:
	go test -v -cover -tags jwx_es256k ./...
