# Compile & run app
run:
	@go run cmd/api/main.go $(ARGS)

# Generate easyjson marshallers
json:
	@easyjson -all internal/adapters/primary/rest/response
	@easyjson -all internal/adapters/primary/rest/request

# Generate swagger docs
documentation:
	@docker run --rm -v .:/source -w /source quay.io/goswagger/swagger:0.30.3 generate spec -m -o ./docs/swagger.json

# Generate custom errors
errors:
	@go run cmd/egen/egen.go