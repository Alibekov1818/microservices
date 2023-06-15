Phones Service

This service is used to create, delete, get phones for electronicstore.
Phones Service uses gRPC to handle requests. Also handler for getting 
single phone can be used with rabbitMQ.

gRPC server runs on port 8002

Methods:

1. `GetPhone` needs PhoneId data type as a request. Returns Phone proto type
2. `GetPhones` needs empty GetPhonesRequest data type as a request. Retuns PhoneList protoType
3. `CreatePhone` needs Phone data type as a request. Returns Phone proto type
4. `DeletePhone` needs PhoneId data type as a request. Returns Phone proto type

RabbitMQ:

You can also use RabbitMQ to get single Phone. Service's queue name is `phones`. In order
to get single phone you have to publish to the queue byte array of PhoneId proto type. The
service will publish result to `phones_client` queue. Don't forget to unmarshall the response.