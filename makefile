# Run app
run:
	@go run cmd/api/main.go $(ARGS)

# Generate custom errors
errors: 
	@go run cmd/egen/egen.go

# Generate swagger docs
documentation:
	@docker run --rm -v $(pwd):/source -w /source quay.io/goswagger/swagger:0.30.3 generate spec -m -o ./docs/swagger.json