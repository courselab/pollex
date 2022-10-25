MOCKS := User

srcs=$(shell find . -type f -name '*.go')

run: backend
	./backend

backend: $(srcs)
	go build -o backend ./cmd/server

coverage:
	cd pkg/handlers && go test -coverprofile=c.out && go tool cover -html=c.out

mocks: coverage update
	cd pkg/controllers && for x in $(MOCKS); do \
		go run github.com/vektra/mockery/v2@latest --name=$$x --filename=$$(echo $$x | tr '[:upper:]' '[:lower:]').go; \
	done

update: tidy
	go get -u ./...

tidy:
	go mod tidy

test:
	AUTH_SERVICE_URL=UNIT_TEST go test -v ./pkg/...

.PHONY: coverage dev mocks update tidy test
