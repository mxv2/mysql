docker-up:
	@docker-compose up -d

export MYSQL_TEST_USER=gotest
export MYSQL_TEST_PASS=secret
export MYSQL_TEST_ADDR=localhost:3307

test:
	@go test ./...