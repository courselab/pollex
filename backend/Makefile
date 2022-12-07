dev:
	go run cmd/server/main.go

coverage:
	cd pkg/handlers && go test -coverprofile=c.out && go tool cover -html=c.out
