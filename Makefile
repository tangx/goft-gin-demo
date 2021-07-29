

MAIN_ROOT ?= cmd/demo

up:
	cd $(MAIN_ROOT) && go run .

tidy:
	go mod tidy
