# Build Guide 🏭

## 🧩 Makefile

You can use `Makefile` in the project inorder to build service as you like.

**Available commands guide**

```
make build-<service> [OPTIONS]   will only build desired service into ./bin
make run-<service> [OPTIONS]     will build and run desired service (depends on 'make build')
make protoc                      will generate protobuf files from ./proto and output into ./pb
clear-log                        will remove anything ending in '.log'

<service>  can be broker, destination, receiver, sender or 'all' in order to
           build all services in a single executable binary file (/bin/all)
           
[OPTIONS]  can be any runtime flags.
           see usage.md for more informations

also 'make build' is shorthand for 'make build-all'
```
see [Usage Guide](https://github.com/SamMHD/simple-broker/blob/main/docs/usage.md) for more information about runtime flags.

## 🐳 Dockerfile

You can easily build image containing all the services (same as ```make build```).
**Note!!** don't forget to expose desired port in Dockerfile

```
docker build -t simple-broker .
```
