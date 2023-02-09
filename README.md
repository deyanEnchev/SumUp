# Job Processing Project
## Introduction
This Go project provides a solution to the SumUp job processing interview problem.

## Usage
While in the root of the project, use this command in order to run it:
    `go run src/main.go`

After running the project, you can either execute POST requests through Postman,
or you can use the following command:

    `$ curl -d @mytasks.json http://localhost:4000/... | bash`

## Testing
While in the root of the project, use this command in order to run all tests for it:
    `go test ./...`