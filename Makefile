

MAIN_ROOT ?= cmd/webserver

up:
	cd $(MAIN_ROOT) && go run .

tidy:
	go mod tidy
