## —— Make —————————————————————————————————————————————————————————————————————
help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## ___ Dev _____________________________________________________________________
reset_databse: ## Reset the database
	@cd database
	@rm db.db
	@sqlite3 db.sqlite

build_api: ## Build the api
	@go build -o bin/api main.go

start_api: build_api ## Run the api
	@./bin/api