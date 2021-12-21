export GOBIN = $(shell pwd)/toolbin

GOIMPORTS = $(GOBIN)/goimports
GOLINT = $(GOBIN)/golint
STATICCHECK = $(GOBIN)/staticcheck
SWAGGER = $(GOBIN)/swagger

.PHONY: run
run:
	@mkdir -p build
	@go build -o build/app
	@./build/app

.PHONY: genswag
genswag:
	@rm -rf models/apimodels restapi/operations
	@$(SWAGGER) generate server \
		-f swagger/swagger.yaml \
		-C swagger/config.yaml \
		-A microservice \
		-P models.Principal \
		-m models/apimodels \
		--keep-spec-order

.PHONY: tools
tools:
	@rm -rf toolbin
	@go install golang.org/x/lint/golint
	@go install honnef.co/go/tools/cmd/staticcheck
	@go install github.com/go-swagger/go-swagger/cmd/swagger
	@go install golang.org/x/tools/cmd/goimports

.PHONY: fmt
fmt:
	@$(GOIMPORTS) -l -w .

.PHONY: lint
lint:
	@./ci/lint.sh
