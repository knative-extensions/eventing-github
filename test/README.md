# Eventing Github TESTS

This directory contains tests and testing docs for `Knative Eventing Github`.

## Running tests with scripts
### Presubmit tests

`presubmit-tests.sh` is the entrypoint for the tests before code submission

You can run it simply with:

```shell
./test/presubmit-tests.sh
```

_By default, this script will run `build tests`, `unit tests` and
`integration tests`._ If you only want to run one type of tests, you can run
this script with corresponding flags like below:

```shell
./test/presubmit-tests.sh --build-tests
./test/presubmit-tests.sh --unit-tests
./test/presubmit-tests.sh --integration-tests
```

## Running tests with `go test` command

### Running unit tests

You can also use `go test` command to run unit tests:

```shell
go test -v ./pkg/...
```
