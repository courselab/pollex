MOCKERY := $(shell command -v mockery 2> /dev/null)
MOCKS := User

dev:
	go run cmd/server/main.go

coverage:
	cd pkg/handlers && go test -coverprofile=c.out && go tool cover -html=c.out

mocks: mockery_dep coverage update tidy
	cd pkg/controllers && for x in $(MOCKS); do \
		mockery --name=$$x --filename=$$(echo $$x | tr '[:upper:]' '[:lower:]').go; \
	done

update: tidy
	go get -u ./...

tidy:
	go mod tidy

deps: mockery_dep

mockery_dep:
ifndef MOCKERY
	@echo $(MOCKERY)
	@echo -n "Mockery is not installed. Would you like to install? [y/N] " && read ans && [ $${ans:-N} = y ]
	go install github.com/vektra/mockery/v2@latest
endif
