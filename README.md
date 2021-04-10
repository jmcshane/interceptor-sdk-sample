# Tekton Triggers Interceptor SDK

This project simplifies the process of creating a custom interceptor to be used
as a Tekton interceptor plugin. The project does the following:

* Templates the necessary Kubernetes resources to run the interceptor
* Sets up the go binary entrypoint and HTTP server
* Handles the interceptor request JSON object and parses into components

Consumers of this library need to write a single function that implements
the Interceptor interface in order to execute their custom interceptor logic.

## Prerequisites

This project depends on the use of [ytt](https://carvel.dev/ytt/) for templating
YAML documents. Please install that tool using the methods discussed [on the 
Carvel project site](https://carvel.dev/#install).

Once the templated files are created, this project builds the go binary into
a Kubernetes deployment using [ko](https://github.com/google/ko). Information
on how to set up a Ko deployment can be found in the [Tekton Triggers repo](https://github.com/tektoncd/triggers/blob/main/DEVELOPMENT.md#environment-setup).

## Getting Started

In order to use this SDK, fork this project and update the `go.mod` package
path as appropriate. Then, update the values in [`config-template/values.yaml`](config-template/values.yaml) to specify the namespace these resources will
be deployed in, the go package path, the name of the custom interceptor, 
and the paths that the interceptor request will be served on.

Once this file is updated, run the following command:

```
$ ytt -f config-template/ --output-files config/
```

This will create a set of config files that can be used in a `ko` build.

Once the `go.mod` file has been updated and the `ytt` generation is complete, you
can run the `ko` build to apply the interceptor to the cluster:

```
$ ko apply -f config/
```

This will create your interceptor and allow you to address it using the Triggers
Interceptor Plugin reference.

## Customization

### Go Module Path

In order to update the Go Module path for your forked repository, replace all instances of the
current module path with your desired path in the project.

### Interceptor Package

The package provided to you for implementing your interceptor is [pkg/interceptor](pkg/interceptor/).
This is referenced in the `server.go` file as an import and running the `NewInterceptor` function. If
you want to change the package name, you will also need to update the import path in the `server.go` 
file in `pkg/server`.

The logic of the interceptor can then be implemented inside of the `Process` function in the
`interceptor` package. This function is intended to encapsulate all the mechanism of passing
an interceptor request from the event listener to an HTTP endpoint, so that the only logic
that must be written is using the parameters of the interceptor request to parse the input
data (body, headers, and extensions) into an InterceptorResponse object. The InterceptorResponse
can determine if Trigger processing should continue, if any new extensions should be added,
or can return an error status message.