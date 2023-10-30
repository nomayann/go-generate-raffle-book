APP=generate-raffle-book

.PHONY: build
build: ## Standard build
	go build -o bin/${APP} main.go

.PHONY: run
run: ## Run 
	go run main.go

.PHONY: compile
compile: ## Compiling for every OS and Platform
	@echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o bin/${APP}-amd64-darwin main.go
	GOOS=linux GOARCH=amd64 go build -o bin/${APP}-amd64-linux main.go
	GOOS=windows GOARCH=amd64 go build -o bin/${APP}-amd64-windows main.go

.PHONY: customs
customs: ## Generate customs.json file
	@echo "Copy customs file to customs.json file"
	cp assets/customs.json.dist assets/customs.json

.PHONY: help
help: ## Display this help.
	@printf "$$(cat $(MAKEFILE_LIST) | egrep -h '^[^:]+:[^#]+## .+$$' | sed -e 's/:[^#]*##/:/' -e 's/\(.*\):/\\033[92m\1\\033[0m:/' | sort -d | column -c2 -t -s :)\n"
