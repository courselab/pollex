MOCKS := User

dev: backend
	./backend

backend:
	go build -o backend cmd/server/main.go

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