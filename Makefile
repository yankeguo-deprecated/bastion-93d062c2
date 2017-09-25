build:
	go build
dist:
	rm -rf dist
	mkdir -p dist/public
	go build -o dist/bastion
	cd ui && make dist
	cd blackbox && go build -o ../dist/blackbox
	cp -r ui/dist/* dist/public
	cp config.sample.toml dist/
lint:
	gofmt -s -w */*.go *.go
pre-test:
	docker stop sandbox-test || true
	docker rm sandbox-test || true
	rm -rf /tmp/test-bastion
.PHONY: build lint dist pre-test
