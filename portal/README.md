# NorthStar Portal (nsportal)

Provides a set of REST interfaces for clients to make use of NorthStar services.

Clients can be interactive user driven webservers;  or driver applications.

## REST enpoints exposed
<<tbd>>

## Environment Variables used

ADVERTISED_PORT": "8080",

NORTHSTARAPI_PROTOCOL ["http"]
NORTHSTARAPI_HOST_PORT 
ENABLE_DEBUG ["false"]
CONNECTION_BUFFER_SIZE" ["1024"]

ACCT_PROTOCOL   ["http"]
ACCT_HOST_PORT" 
ACCT_AUTH_HOST_PORT": "@ACCT_AUTH_HOST_PORT@",
ACCT_CLIENT_ID": "@CLIENT_ID@",
ACCT_SECRET": "@CLIENT_SECRET@",
ACCT_CLIENT_SCOPES": "ts.configuration",
ACCT_USER_SCOPES": "ts.user ts.user.ro ts.transformation ts.transformation.ro ts.notebook ts.notebook.ro ts.model.ro ts.nsobject.ro",


KAFKA_BROKERS_HOST_PORT": "@KAFKA_BROKERS_HOST_PORT@",
ZOOKEEPER_HOST_PORT": "@ZOOKEEPER_HOST_PORT@",
LOGGER_KAFKA_BROKERS_HOST_PORT": "@LOGGER_KAFKA_BROKERS_HOST_PORT@",
LOGGER_ZOOKEEPER_HOST_PORT": "@LOGGER_ZOOKEEPER_HOST_PORT@",

