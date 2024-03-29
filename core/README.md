# Core

The core applications constructing the necessary habitat for tulip.
Core is a microservice architecture which revolves around gRPC (google remote procedure call.)

## Ports

Each microservice gets designated its own HTTP and GRPC port.

- api:
  - http: 8000
  - grpc: 8001
- broker:
  - tcp: 1883
  - ws: 1882
  - http: 1881
- discovery:
  - http: 8002 (not in use)
  - grpc: 8003
- gateway:
  - http: 8004
  - grpc: 8005
- pki:
  - http: 8006 (not in use)
  - grpc: 8007

## api

The api microservice is the one exposed both internally and externally. It allows other microservices the report changes
on an entity and CRUD. This allows us to only keep "1" connection the database (ideally). Both the web and the 
Tulip Interface both communicate over the http api with authenticated tokens.

## broker

The broker microservice is an "individual" microservice. It does not rely on anyone else, but it provides functionality
consumed in the discovery microservice.

## discovery

The discovery microservice is responsible for discovering new devices and performing the actual communication with them.
Devices might be discovered over Wind, mQTT or mDNS DNS-SD. In the case of mQTT devices a gRPC api is exposed so that
services such as the api can forward commands.

## gateway

The gateway microservice exposes a websocket for interfaces and the web ui to consume. It does not keep any state itself
it only forwards events generated by other services. Such as when entity state changes.

## pki

The pki microservice is responsible for generating certificates and tokens for Wind devices. It will also verify tokens.