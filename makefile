run:
	@go run cmd/api/main.go $(ARGS)

# Generate custom errors
egen: 
	@go run cmd/egen/egen.go