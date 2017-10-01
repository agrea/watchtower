DEFAULT_GOAL: build
ENVIRONMENT:=tmp/environment

$(ENVIRONMENT):
	$(info --- Launching environment)
	docker-compose up --build -d
	mkdir -p $(@D)
	touch $@

.PHONY: build
build:
	$(info --- Building Watchtower)
	go build .

.PHONY: clean
clean:
	$(info --- Cleaning environment)
	@docker-compose down
	rm -rf $(ENVIRONMENT)

.PHONY: watch
watch: $(ENVIRONMENT)
	$(info --- Starting Watchtower in watch mode)
	docker-compose run --rm -p 3000:3000 watchtower realize run .
