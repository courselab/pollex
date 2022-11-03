docker:
	docker-compose up -d --scale backend=3

dev_server:
	make -C backend dev
auth_server:
	make -C auth run

.PHONY: dev_server auth_server
