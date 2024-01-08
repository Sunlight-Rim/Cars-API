# Run app
run:
	@go run cmd/api/main.go $(ARGS)

# Generate custom errors
errors: 
	@go run cmd/egen/egen.go

# Generate swagger docs
documentation:
	@swagger generate spec -m -o docs/swagger.json