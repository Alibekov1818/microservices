Token Service

This service is used to get tokens for electronicstore. Token Service uses 
RabbitMQ in order to get requests with queue name `token`.
The logic of the service was implemented by gRPC.

How to use

1. Grpc server runs in port 8001 
so you can pass gRPC request in `GetToken` method.
2. You can publish marshalled proto request to `token` queue,
which will return byte array of proto request.