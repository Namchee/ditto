## Ditto

Ditto is a CLI testing tool that helps you verify if multiple HTTP endpoints have the same outputs. Ditto is designed to be performant out of the box by being able to execute tests in parallel without any extra steps.

Ditto is also designed to be configurable out of the box by using a configuration file.

One of the use case of Ditto is regression testing between two or more HTTP endpoints.

### Installation

TBD

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
`name` | `string` | Test name. Used when reporting test run results.
`endpoints` | `Endpoints` | List of endpoints that should be tested.

### Usage

#### `ditto`

TBD

#### `ditto-gen`

`ditto-gen` is a utility that helps you create your test definitions. `ditto-gen` can be executed from your terminal with the following command.

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

### License

This project is licensed under the [MIT License](./LICENSE)