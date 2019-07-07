# Dogsitter

[![Build Status](https://cloud.drone.io/api/badges/renaudhager/dogsitter/status.svg)](https://cloud.drone.io/renaudhager/dogsitter) [![codecov](https://codecov.io/gh/renaudhager/dogsitter/branch/master/graph/badge.svg)](https://codecov.io/gh/renaudhager/dogsitter)

## Description
This small tool, allow to manipulate dashboard from Datadog.

## Usage
```
$ dogsitter
NAME:
   dogsitter - A new cli application

USAGE:
   CLI tool to import and export Datadog dashboard. [global options] command [command options] [arguments...]

COMMANDS:
     delete   Delete dashboard from Datadog.
     list     List dashboard existing in Datadog.
     pull     Pull dashboard configuration from Datadog API
     push     Import dashboard configuration to Datadog.
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --api-key value                           Datadog API key [$DATADOG_API_KEY]
   --app-key value, --application-key value  Datadog Application key [$DATADOG_APP_KEY]
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

### Push command
```
$ dogsitter push -h
NAME:
   CLI tool to import and export Datadog dashboard. push - Import dashboard configuration to Datadog.

USAGE:
   CLI tool to import and export Datadog dashboard. push [command options] [arguments...]

OPTIONS:
   -f value, --file value  Dashboard file configuration. [$DS_IMPORT_FILE]
```

### List command
```
$ dogsitter list -h
NAME:
   CLI tool to import and export Datadog dashboard. list - List dashboard existing in Datadog.

USAGE:
   CLI tool to import and export Datadog dashboard. list [command options] [arguments...]

OPTIONS:
   --format text             Format of the list of dashboard. Supported values are text or json. (default: "text")
   -o value, --output value  output file to print dashboard list. (default: "stdout")
   --id value                Get detail for a specific dashboard.
```

### Delete command
```
$ dogsitter delete -h
NAME:
   CLI tool to import and export Datadog dashboard. delete - Delete dashboard from Datadog.

USAGE:
   CLI tool to import and export Datadog dashboard. delete [command options] [arguments...]

OPTIONS:
   --id value  Dashboard id.
```
