build:
	@go build

run:
	@go build && ./integration-test

run-api-test:
	@go build && ./integration-test -module='API'

run-gql-test:
	@go build && ./integration-test -module='GQL'