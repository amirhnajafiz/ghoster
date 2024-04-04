# :ghost: Ghoster

Implementing a simple serverless system for executing *Golang* applications, aka **Ghoster**. This system is a **FaaS**(Function as a Service). As the name suggests, it executes Golang projects on demand. Ghoster is a minimal version of serverless applications.

## Run using Docker

In order to execute Ghoster, you need to set some environment variables, mount the functions directory, and build a docker image.

### env variables

### functions dir mount

### docker command

## Interface

Ghoster uses **HTTP** requests for registring projects, listing projects, and executing them.

```sh
curl -i -X POST \
  -F "file_name=subtraction" \
  -F "file=@Archive.zip" \
  localhost:5002/files
```

## Metrics

## Suggestions

## Contribute
