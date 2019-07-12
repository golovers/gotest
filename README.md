# Testing in Go

## Structure:
    xxx.go
    xxx_test.go
        TestXXX(t *testing.T)
        BenchmarkXXX(t *testing.B)
    xxx_integration_test.go
        TestXXX(t *testing.T)
    testdata/
        xxx.txt
        xxy.csv

## Basic commands:
    go test
    go test -v
    go test ./...
    go test -run=TestRotate

## Benchmark
    go test -benchmem -bench=.

## Coverage
    go test -cover
    go test ./... -coverprofile=coverage.out
    go tool cover -func=coverage.out
    go tool cover -html=coverage.out

## Test suite
    xxx_integration_test.go
        // +build integration
    go test -tags=integration

## HTTP:
    httptest.ResponseRecorder
    httptest.Server
## Mocks:
    https://github.com/golang/mock

## Blogs:
    https://blog.golang.org/cover