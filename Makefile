ifeq (migrate,$(firstword $(MAKECMDGOALS)))
  DIRECTION := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(DIRECTION):;@:)
endif

.PHONY: migrate 
migrate:
	@go run cmd/migrations/main.go -migrate $(DIRECTION)

startup:
	@go run cmd/migrations/main.go -migrate up
	@go run cmd/rest/main.go

start: 
	@go run cmd/rest/main.go