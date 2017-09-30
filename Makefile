DEFAULT_GOAL: build

.PHONY: build
build:
	$(info --- Building Watchtower)
	go build .

.PHONY: watch
watch:
	$(info --- Starting Watchtower in watch mode)
	gin run main.go
