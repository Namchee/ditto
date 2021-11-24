## Ditto

[![Go Report Card](https://goreportcard.com/badge/github.com/Namchee/ditto)](https://goreportcard.com/report/github.com/Namchee/ditto)

Ditto is a CLI testing tool that helps you verify if multiple HTTP endpoints have the same outputs. 

One of the use case of Ditto is regression testing between two or more HTTP endpoints.

### Features

1. Works with normal HTTP REST API and GraphQL endpoints with various stringifiable response bodies.
2. Parallel test execution.
3. Configurability.

### Installation

> Make sure that you have [Go](https://golang.org/doc/install) installed in your machine.

Ditto comes with two separate binaries, `ditto` and `ditto-gen`. `ditto` is the core of Ditto that executes your tests while `ditto-gen` is a utility tool that helps you write tests for `ditto`.

You can install Ditto by executing the following command in your terminal.

```bash
go install github.com/Namchee/ditto/cmd/ditto
```

To install `ditto-gen` instead, you can execute the following command.

```bash
go install github.com/Namchee/ditto/cmd/ditto-gen
```

### Test Structure

When Ditto is being executed, Ditto will try to search for test definitions that should be stored in the test directory folder.

A test is defined as a JSON file with the following format.

Key | Type | Description
--- | ---- | -----------
`name` | `string` | Test name. Used when reporting test run results.
`endpoints` | `[]Endpoints` | List of endpoints that should be tested. There must be at least two endpoints for a valid test definition.

An endpoint is defined as an object with the following format.

Key | Type | Description
--- | ---- | -----------
`host` | `string` | Host name. Must be an IP  or an URL.
`method` | `string` | Case-sensitive HTTP method. Must be either `GET`, `POST`, `PUT`, `PATCH`, or `DELETE`
`query` | `object` | Query object to be sent when sending HTTP request.
`body` | `object` | Request body to be sent when sending HTTP request.
`headers` | `object` | Request headers to be sent when sending HTTP request.
`timeout` | `integer` | Endpoint timeout in seconds. A test will automatically fail if an endpoint timeouts.

Please refer to [ditto-test directory](./ditto-test) for test samples.

### Usage

#### `ditto`

The `ditto` command is a command that runs predefined tests in the current working directory. This command does not accept any extra inputs.

#### `ditto-gen`

The `ditto-gen` command is a utility command that helps you create your test definitions. `ditto-gen` can be executed from your terminal with the following command.

```bash
ditto-gen <file_name> <test_name>
```

`ditto-gen` accepts the following arguments:

Name | Type | Description
--- | ---- | -----------
`filename` | `string` | Test file name. Should be suffixed with `.json`.
`testname` | `string` | Test name. Used on reporting and test execution logs.

### Configuration

When Ditto is executed, Ditto will try to look for `ditto.config.json` at the current working directory. If the corresponding file is found, Ditto will use the file as a configuration file that modifies the behavior of Ditto.

Below are the list of possible configuration for Ditto.

Name | Type | Default | Description
---- | ---- | ------- | -----------
`test_directory` | `string` | `ditto-test` | Test directory which stores all test definitions that are going to be executed by Ditto.
`log_directory` | `string` | `ditto-log` | Test execution log directory which stores all test logs when one or more tests are failing.
`strict` | `boolean` | `false` | Determine if Ditto should stop test execution when one or more test definitions are invalid.
`workers` | `integer` | `<all_cores>` | Determine the maximum number of tests that should be executed in parallel.
`status` | `boolean` | `false` | Determine if passing tests should also require the same HTTP response status.

### Test Logs

When a test fails, Ditto will attempt to log the test execution result in a folder with the same name as `log_directory` in JSON format. A test log will be named with the test name corresponding to the test data.

Test logs is an object with the following properties.

Name | Type | Default | Description
---- | ---- | ------- | -----------
`name` | `string` | `ditto-test` | Test name.
`err` | `string` | `ditto-log` | Test errors when calling the endpoint.
`result` | `[]Endpoint` | Fetch result.

### GraphQL Testing

This tool can also be used to test GraphQL queries and mutation. However, you cannot test two different query and mutations and expecting the same result. For example:

```
{
    a {
        foo
    }
}
```

```
{
    b {
        foo
    }    
}
```

Although the above queries returns the exact same result, `ditto` will always mark the test as fail since the query name is different.

To address those issues, you can use [GraphQL alias](https://graphql.org/learn/queries/#aliases). For example, the test above should be executed as:

```
{
    data: a {
        foo
    }
}
```

```
{
    data: b {
        foo
    }
}
```

### License

This project is licensed under the [MIT License](./LICENSE)