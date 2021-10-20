# demo-sr-app

This repository contains a demo SR-App.

It makes some requests and subscriptions to the Jalapeño API Gateway and prints the result to the console.

## Run Demo

To run the demo, install [golang](https://golang.org/doc/install) on your machine and clone this repository, then run the following command in the root directory of the project:

```bash
go run main.go <server-address> <request-service-port> <subscription-service-port>
```

The **server-address** points to the server where the Jalapeño API Gateway is installed.

The default ports for the services are:

Service | Port
--- | ---
Request-Service | 30061
Subscription-Service | 30060
