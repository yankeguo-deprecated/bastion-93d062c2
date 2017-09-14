build:
	go build
dist:
	mkdir -p dist/public
	go build -o dist/bastion
	cd ui && npm install && NODE_ENV=production npm run build
	cp -r public/* dist/public
	cp -r ui/dist/* dist/public
	cp config.sample.toml dist/
lint:
	gofmt -s -w */*.go *.go
pre-test:
	docker stop sandbox-test || true
	docker rm sandbox-test || true
	rm -rf /tmp/test-bastion
.PHONY: build lint dist pre-test
