# Go & PostgreSQL E2E tests example

### Commands
```bash
# run tests
# set DATABASE_URL to your test database
DATABASE_URL=postgres://Michal:@localhost/testDatabase?sslmode=disable go test -v ./...

# to run a binary you need to create .env file and put there your database credentials
# (see .env.example)
go build
./e2e-tests-example
```
