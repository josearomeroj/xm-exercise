# xm-exercise

## Info

This is a programming exercise for XM.
Everything could be better but there is not enough time to make it perfect.

### Task: 
Build a microservice to handle companies. It should provide the following operations:
    
- [x] Create
- [x] Patch
- [x] Delete
- [x] Get

Plus:

- [x] On each mutating operation, an event should be produced.
- [x] Dockerize the application to be ready for building the production docker image
- [x] Use docker for setting up the external services such as the database
- [x] REST is suggested, but GRPC is also an option
- [x] JWT for authentication
- [x] Kafka for events
- [x] DB is up to you
- [x] Integration tests are highly appreciated*
- [x] Linter
- [x] Configuration file

## Requirements

- Docker version 19.03.0+
- Golang version 1.19+
- Protoc version 3+

### Protoc plugins

To generate code using this repository, you need to install the following protoc plugins:

- [protoc-gen-validate](https://github.com/bufbuild/protoc-gen-validate)
- [protoc-gen-gofullmethods](https://github.com/nicovogelaar/protoc-gen-gofullmethods)

### Protoc generation

[OPTIONAL] To regenerate the proto files use:

```
 go run ./scripts/gen_protobuf.go -gen
```

## Set up environment

````
git clone https://github.com/josearomeroj/xm-exercise.git
cd xm-exercise
docker-compose up -d
````

This will set up the environment using the defined services in the docker-compose.yaml file.
[MongoDB, Kafka, Zookeeper]

## Run

### Run tests

````
go test ./...
````

### Run server

This will run the server using the example config file, located at: example_config.yaml

````
go run ./cmd/xm-server
````

#### Custom config

You can set you own config using the yaml:

````
go run ./cmd/xm-server -config <path>
````

You can replace any configuration property by leaving it blank and setting an env variable.
Example:

ENV:

````
export AUTH_CONFIG_PRIVATE_KEY='value'
export AUTH_CONFIG_PUBLIC_KEY='value'
````

YAML:

````
AUTH_CONFIG:
    PRIVATE_KEY:
    PUBLIC_KEY:
    JWT_VALIDITY_MILLIS: 600000
````

## Docker build

To build this service for docker just execute:

````
docker build -t "xm-server:1.0.0" .
````

And set the config properties with env variables.