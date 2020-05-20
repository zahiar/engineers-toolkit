# Engineer's Toolkit
An interactive CLI app designed to help automate some of the manual tasks you may need to do from time-to-time.

## Running the CLI app
### Requirements
1. Have GoLang installed
1. Have Docker installed - if you plan to test out the CLI app against the dockerised apps

### Instructions
To view the built-in help, run:
```
go run main.go
```

Example usages:
```
go run main.go logging help
go run main.go logging create
```

### Running the dockerised setup
```
docker-compose up
```

If you just want to run a single service, then you can do:
```
docker-compose run graylog
docker-compose run jenkins
```

## Errors
###### Graylog Docker ElasticSearch container errors on start up
Run this on the host machine, and then try re-launching the Docker containers
`sysctl -w vm.max_map_count=262144`
