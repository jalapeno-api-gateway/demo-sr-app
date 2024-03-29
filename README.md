<p align="center">
    <img src="https://github.com/jalapeno-api-gateway/.github/raw/main/img/logo.png">
</p>
<h1 align="center">Demo SR App</h1>
<p align="center">
The Demo-SR-App repository contains a demonstration app for the Jalapeno API Gateway.
</p>

---

## About
The application makes some requests and subscriptions to the Jalapeño API Gateway and prints the result to the console.

## Run Demo
To run the demo, install [golang](https://golang.org/doc/install) v1.16 or higher on your machine and clone this repository, then run the following command in the root directory of the project:

```bash
go run main.go <server-address> <request-service-port> <subscription-service-port>
```

The **server-address** points to the server where the Jalapeño API Gateway is installed.

The default ports for the services are:

| Service | Port |
| --- | --- |
| Request-Service | 30050 |
| Subscription-Service | 30051 |

## Disclaimer
This app is built and maintained for demonstration purposes for the INS Institute for Network and Security at the [OST - Eastern Switzerland University of Applied Sciences](https://www.ost.ch/en/). It only runs internally since it makes request and subscriptions to specific nodes in a local network. Other SR-App developers may use this application as a more elaborate sample app.
