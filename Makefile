IMAGE = mikulas/color-demo
VERSION = 1.0.0

.PHONY: _
_: build publish

.PHONY: build
build:
	docker build -t $(IMAGE):$(VERSION) --file Dockerfile .

.PHONY: publish
publish:
	docker push $(IMAGE):$(VERSION)
