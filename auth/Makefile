srcs=$(shell find . -type f -name '*.go')

run: auth .credentials
	@echo Service key is $$(grep -oP 'TEST_SERVICE_KEY=\K(.+)' .credentials)
	@TOKEN_KEY=$$(grep -oP 'TEST_TOKEN_KEY=\K(.+)' .credentials) SERVICE_KEY=$$(grep -oP 'TEST_SERVICE_KEY=\K(.+)' .credentials) ./auth

auth: $(srcs)
	go build

.credentials:
	@printf TEST_TOKEN_KEY= >> .credentials
	@head -c32 /dev/random | od -A n -v -t x1 | tr -d '[:space:]' >> .credentials
	@echo >> .credentials
	
	@printf TEST_SERVICE_KEY= >> .credentials
	@head -c16 /dev/random | od -A n -v -t x1 | tr -d '[:space:]' >> .credentials
	@echo >> .credentials


.PHONY: run
