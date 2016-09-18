# Proxa

Proxa is a configurable HTTP reverse proxy.

It is still in development and should not be used in production.

## Install

    $ go get github.com/wridgers/proxa

## Usage ##

The default configuration path is `./proxa.toml`.

    $ proxa -config /path/to/configuration.toml

## Configuration

Proxa is configured with a simple [TOML](https://github.com/toml-lang/toml) file.

```toml
bind = ":8080"

################################################################################

[[routes]]
prefix = "/api"
backend = "api"

[[routes]]
prefix = "/"
backend = "default"

################################################################################

[[backends]]
name = "default"
addrs = ["localhost:8001"]

[[backends]]
name = "api"
addrs = ["localhost:9001", "localhost:9002"]
```
