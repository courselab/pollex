docker: .test-env
	docker-compose --env-file .test-env up -d --scale backend=3

.test-env:
	make -C auth .credentials
	grep -oP 'TEST_\K(.+)' auth/.credentials > .test-env

dev_server:
	make -C backend dev
auth_server:
	make -C auth run

.PHONY: dev_server auth_server
