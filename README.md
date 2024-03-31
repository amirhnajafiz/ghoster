# :ghost: Ghoster

Simple serverless system for executing *Golang* applications, aka **Ghoster**. This system is a **FaaS**(Function as a Service). As the name suggests, it executes Golang projects on demand.

```sh
curl -i -X POST \
  -F "file_name=subtraction" \
  -F "file=@Archive.zip" \
  localhost:5002/files
```

## Components

Ghoster consists of three major components as descripted below:

1. file manager: handles the Golang projects
   1. upload
   2. index
   3. list
   4. delete
2. core: is a gateway for running projects
   1. execute
3. metrics server: manages the app metrics (metrics exporter)
