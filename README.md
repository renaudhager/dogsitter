# Dogsitter

## Description
This small tool, allow to export (pull) and import (push) dashboard from/to Datadog.

## Usage
```
$ dogsitter
NAME:
   dogsitter - A new cli application

USAGE:
   CLI tool to import and export Datadog dashboard. [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     pull     Pull dashboard configuration from Datadog API
     push     Import dashboard configuration to Datadog.
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --api-key value                           Datadog API key [$DD_API_KEY]
   --app-key value, --application-key value  Datadog Application key [$DD_APPLICATION_KEY]
   -l value, --log-level value               Setting log level (default: "INFO") [$DS_LOGLEVEL]
   --dh value, --datadog-host value          Datadog endpoint (default: "https://app.datadoghq.eu") [$DD_HOST]
   --help, -h                                show help
   --version, -v                             print the version
```

### Pull command
```
$ dogsitter pull -h
NAME:
   CLI tool to import and export Datadog dashboard. pull - Pull dashboard configuration from Datadog API

USAGE:
   CLI tool to import and export Datadog dashboard. pull [command options] [arguments...]

OPTIONS:
   --id value                id of dashboard
   -o value, --output value  output file for JSON payload. (default: "stdout")
```

### Push
```
$ dogsitter push -h
NAME:
   CLI tool to import and export Datadog dashboard. push - Import dashboard configuration to Datadog.

USAGE:
   CLI tool to import and export Datadog dashboard. push [command options] [arguments...]

OPTIONS:
   -f value, --file value  Dashboard file configuration. [$DS_IMPORT_FILE]
```
