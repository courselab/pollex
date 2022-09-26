USER=courselab
PROJECT_NAME=pollex
GOPATH=$(shell go env | grep GOPATH | cut -d= -f2)

.PHONY: install_proj

run:
	cd backend && go run main.go

install_proj: go_check path_check
		@echo "Installing project in GOPATH"
		ln -s "$(PWD)" "$(GOPATH)/src/$(USER)/$(PROJECT_NAME)"

go_check:
		@echo GOPATH found in $(GOPATH)
		@if [ -z "$(GOPATH)" ]; then \
			echo "GOPATH is not set"; \
			exit 1; \
		fi

path_check:
		@if [ ! -d "$(GOPATH)/src" ]; then \
			echo "Creating directory $(GOPATH)/src"; \
			mkdir -p "$(GOPATH)/src"; \
		fi
		@if [ ! -d "$(GOPATH)/src/$(USER)" ]; then \
			echo "Creating directory $(GOPATH)/src/$(USER)"; \
			mkdir -p "$(GOPATH)/src/$(USER)"; \
		fi
