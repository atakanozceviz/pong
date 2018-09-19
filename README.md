# Pong

Pong is a cli application that can run user defined commands in many user defined paths. 
Pong can run commands in parallel if you wish.

## Install

```
go get github.com/atakanozceviz/pong
```

## Getting Started

You will need to define your list of paths and commands in a json file.

##### Example JSON:

```json
{
    "paths": {
        "solutionPaths": [
            "framework",
            "modules/users",
            "modules/background-jobs"
        ],
        "testProjectPaths": [
            "framework/test/Authentication.OAuth.Tests",
            "framework/test/MultiTenancy.Tests",
            "framework/test/Mvc.Tests"
        ]
    },
    "commands": {
        "build": {
            "description":"Builds a project and all of its dependencies.",
            "workers": 1,
            "onerror": "exit",
            "name": "dotnet",
            "args": [
                "build"
            ]
        },
        "test": {
            "description":".NET test driver used to execute unit tests.",
            "workers": 3,
            "onerror": "continue",
            "name": "dotnet",
            "args": [
                "test",
                "--logger:trx"
            ]
        }
    }
}
```

In the above example there are 2 objects at the root, "paths" and "commands". These names can not be changed, however you can decide the name of path arrays and command objects as you want. You will use these names to exucute `pong` (In the example; `"solutionPaths"`, `"testProjectPaths"` are path arrays and `"build"`, `"test"` are commands)

`"test"` command in the example will be interpreted like so:

```command
$ dotnet test --logger:trx
```

`Commands` must be in this form:

```json
{
  "CommandName": {
    "description": Short description for this command.,
    "workers": An integer minimum of 1,
    "onerror": What to do on error. Can be "exit" or "continue",
    "name": Command to exucute,
    "args": [
      "Arguments",
      "for",
      "command"
    ]
  }
}
```

`"workers"` defines how many commads can be run in parallel. (Setting to 1 would run one command at a time)
`"description"` will be shown like this in cli: 

```console
...
Available Commands:
  build       Builds a project and all of its dependencies.
  test        .NET test driver used to execute unit tests.
...
```

`Paths` are just arrays of strings.

```json
{
  "NameOfYourPathArray": [
    "a/valid/path",
    "/another/valid/path",
    "and/one/valid/file.txt"
  ]
}
```

## How to use

Pong tries to load "pong.json" file in the current directory. You can specify json file path using the `-s` flag.

```console
$ pong
Pong is a cli application that can run user
defined commands in many user defined paths. 
Example:
pong [command] [path]

Usage:
  pong [command]

Available Commands:
  build       Builds a project and all of its dependencies.
  help        Help about any command
  test        .NET test driver used to execute unit tests.
  version     Print the version number of Pong

Flags:
      --config string   config file (default is $HOME/.pong.yaml)
  -h, --help            help for pong
  -t, --toggle          Help message for toggle

Use "pong [command] --help" for more information about a command.
```

For the above sample json:

```console
$ pong test testProjectPaths
```

command would run `dotnet test --logger:trx` in these paths 3 at the same time:

```
framework/test/Authentication.OAuth.Tests
framework/test/MultiTenancy.Tests
framework/test/Mvc.Tests
```
