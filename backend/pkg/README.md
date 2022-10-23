
## Code coverage
Before running the command go to backend root (backend/)<br>

`make coverage`

## Update dependencies
Before running the command go to backend root (backend/)<br>

`go get -u ./... `

`go mod tidy`

## Mock create/update example
Before running the command go to controller package (controllers/)<br>

`mockery --name=User --filename=user.go`
