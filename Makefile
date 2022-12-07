dev_server:
	make -C backend dev
auth_server:
	make -C auth run

.PHONY: dev_server auth_server
