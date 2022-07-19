export SHELL:=/bin/bash

.PHONY: mysql-fs-image

mysql-fs-image: GOOS=linux CGO_ENABLED=0 go build -v -ldflags '-extldflags -static' -o lightflow cmd/mysql_fs/main.go

.PHONY: fs-mysql-image

fs-mysql-image: GOOS=linux CGO_ENABLED=0 go build -v -ldflags '-extldflags -static' -o lightflow cmd/fs_mysql/main.go
