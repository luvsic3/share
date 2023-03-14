# share

Share directories and files from the CLI to iOS and Android devices without the need of an extra client app

We use Cobra for creating powerful modern CLI applications, [Taskfile](https://dev.to/stack-labs/introduction-to-taskfile-a-makefile-alternative-h92) (a Makefile alternative). 

## Usage

```
Usage:
  share /path/to/directory [flags]
  share [command]

Available Commands:
  clipboard   Share Clipboard content
  completion  generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help              help for share
      --ip string         Your machine public ip address
  -P, --password string   Set basic authentication password
  -U, --username string   Set basic authentication username

Use "share [command] --help" for more information about a command.
```

## Pre-requisits

Install Go in 1.16 version minimum.

## Build the app

`$ task build`

## Run the app

`$ task run`

## Credit

This project is deeply inspired by [sharing](https://github.com/parvardegr/sharing)