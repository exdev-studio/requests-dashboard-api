# Requests Dashboard API

Collect and analyze incoming request

## Requirements

* Go >= 1.15 (earlier versions have not been tested)

## Install

```shell
git clone git@github.com:exdev-studio/requests-dashboard-api.git
```

## Test

To test the application before you build it you need to run

```shell
make test
```

## Build

In case you have `make` installed

```shell
make
```

Please ensure you have Go installed before you try to run the snippet above

## Install app

To install the app in your `$GOBIN` you need the followings:

```shell
make install
```

## Run

To run the server you need to execute the following

```shell
./bin/apiserver
```

## Help

In case you're wondering what arguments the server has you can run

```shell
./bin/apiserver -help
```

As soon as you run that snippet you will get a built help

## Deployments

There are two possible ways to deploy the app:

* Manual Docker
* Using Docker-compose

### Manual Docker

```shell
./scripts/docker-build.sh && ./scripts/docker-run.sh
```

### Using Docker-compose

```shell
./scripts/docker-compose-run.sh
```
