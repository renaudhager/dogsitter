# Dogsitter

[![Build Status](https://cloud.drone.io/api/badges/renaudhager/dogsitter/status.svg)](https://cloud.drone.io/renaudhager/dogsitter) [![Go Report Card](https://goreportcard.com/badge/github.com/renaudhager/dogsitter)](https://goreportcard.com/report/github.com/renaudhager/dogsitter) [![codecov](https://codecov.io/gh/renaudhager/dogsitter/branch/master/graph/badge.svg)](https://codecov.io/gh/renaudhager/dogsitter)

## Description
A small command-line tool to manipulate Datadog dashboards.

## Usage
```
$ dogsitter
NAME:
   dogsitter - manipulate Datadog dashboards

USAGE:
   CLI tool to manipulate Datadog dashboards. [global options] command [command options] [arguments...]

COMMANDS:
     delete   Delete a dashboard
     list     List the dashboards
     pull     Retrieve a dashboard configuraton
     push     Upload a dashboard configuration
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --api-key value                           Datadog API key [$DATADOG_API_KEY]
   --app-key value, --application-key value  Datadog Application key [$DATADOG_APP_KEY]
   -l value, --log-level value               Log level (default: "INFO") [$DS_LOGLEVEL]
   --dh value, --datadog-host value          Datadog endpoint (default: "https://app.datadoghq.eu") [$DD_HOST]
   --help, -h                                Show help
   --version, -v                             Print the version
```

### Pull command
```
$ dogsitter pull -h
NAME:
   CLI tool to manipulate Datadog dashboards. pull - Retreive dashboard configuration from the Datadog API.

USAGE:
   CLI tool to manipulate Datadog dashboards. pull [command options] [arguments...]

OPTIONS:
   --id value                ID of dashboard
   -o value, --output value  Output file for JSON payload (default: "stdout")
```

### Push command
```
$ dogsitter push -h
NAME:
   CLI tool to manipulate Datadog dashboards. push - Upload dashboard configuration to Datadog.

USAGE:
   CLI tool to manipulate Datadog dashboards. push [command options] [arguments...]

OPTIONS:
   -f value, --file value  File to read configuration from [$DS_IMPORT_FILE]
```

### List command
```
$ dogsitter list -h
NAME:
   CLI tool to manipulate Datadog dashboards. list - List existing dashboards in Datadog.

USAGE:
   CLI tool to manipulate Datadog dashboards. list [command options] [arguments...]

OPTIONS:
   --format text             Output format. Supported values are "text or "json" (default: "text")
   -o value, --output value  Output file (default: "stdout")
   --id value                Get details for a specific dashboard
```

### Delete command
```
$ dogsitter delete -h
NAME:
   CLI tool to manipulate Datadog dashboards. delete - Delete dashboard from Datadog.

USAGE:
   CLI tool to manipulate Datadog dashboards. delete [command options] [arguments...]

OPTIONS:
   --id value  Dashboard ID
```
