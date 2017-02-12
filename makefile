design_path="github.com/viejodecaldas/ms-object-minesweeper/src/web/design"
BUILD := $(shell date -u +%m%d%H%M)

.PHONY: regen
regen: backup clean gen restore

.PHONY: backup
backup:
	@echo "backing up"
	@find src/web/. -maxdepth 1 -name "*.go" -exec cp -f {} {}.backup \;

.PHONY: clean
clean:
	@echo "cleaning"
	@rm -rf src/web/app src/web/client src/web/swagger src/web/tool
	@find src/web/. -maxdepth 1 -name "*.go" -exec rm -f {} \;

.PHONY: gen
gen:
	@echo "generating $(design_path)"
	@goagen bootstrap -d $(design_path) -o src/web

.PHONY: restore
restore:
	@echo "restoring"
	@./src/web/restore.py

.PHONY: clean_backups
clean_backups:
	@find src/web/. -maxdepth 1 -name "*.backup" -exec rm -f {} \;

.PHONY: clean_all
clean_all: clean clean_backups

clean-ebs:
	@echo "cleaning"
	@rm -rf src/web/app src/web/client src/web/swagger src/web/tool

build:
	@cd src/web && go build -o ../../bin/server -ldflags "-X main.BUILD=$(BUILD)"

run: build
	cd bin/ && ./server

ebs-local:
	go build -o bin/ebs/application -tags="production" -ldflags "-s" -ldflags "-X main.BUILD=$(BUILD)" .

ebs-build: clean-ebs gen
	rm -rf bin/ebs
	@cd src/web && GOOS=linux GOARCH=amd64 go build -o ../../bin/ebs/application -tags="production" -ldflags "-s" -ldflags "-X main.BUILD=$(BUILD)" .
	cd bin/ebs && zip -r app-$(shell date -u +%y%m%d-%H%M).zip .