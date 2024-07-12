# Internal GitHub Actions

This repository contains a collection of internal GitHub Actions that are used by Thoughtgears. All of them are released 
using the [Thoughtgears GitHub Release](./actions/github-release/README.md) action.  
All the actions are located under the `actions` directory. 

## Go actions

Actions written in go are all built and bundled during the CI process. All go actions will have a [JS shim](./template/run.js)
that will call the go binary. This will allow the actions to be used in the same way as a JS action and will have a much
faster execution time compared to `composite` or `container` actions.

## Development

### Prerequisites

- [Go](https://golang.org/dl/) for development of go actions.
- [Task](https://taskfile.dev/#/installation) for running the tasks.
- [GolangCI-Lint](https://golangci-lint.run/usage/install/) for linting and static analysis.


### Adding a new action

To add a new action, run `task scaffold -- github-release`, it will create a new directory under `actions` with the
required files and directories. 

### Testing

To test the actions, run `task your-action:test`, it will run the tests for all the actions.  
To lint the actions, run `task your-action:lint`, it will run the linter for all the actions.

### Building

To build the actions, run `task your-action:bundle`, it will build the go actions and create a `bin/action-name-sha` binary.

