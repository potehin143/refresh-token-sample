# Refresh token usage sample

##Test task
###Description
Create an authentication service. There should be 4 REST API routes with the following functionality:
1. returning Access and Refresh tokens for the user whose identifier (UUID) is specified in the request parameter;
2. refreshing the Access token;
3. removing of the specified Refresh token from DB;
4. removing all Refresh tokens that relate to a certain user from DB.
There are several requirements to the tokens:
1. Refresh token must be protected from changes on the client-side
2. Refresh token cannot be reused more than once
3. Refresh operation for an Access token can be performed only with the Refresh token that was issued along with it.
###Technologies
1. Programming language: Go
2. DB:
- DBMS: MongoDB
Topology: Replica set
3. Access token
- Type: JWT
- Encryption algorithm: SHA 512
4. Refresh token
- Type: any
- Transfer format: base 64
- Storing in DB: bcrypt hash


### Build

How to get required dependencies required 
~~~
go get github.com/gorilla/mux
go get github.com/dgrijalva/jwt-go
go get go.mongodb.org/mongo-driver/mongo
go get github.com/google/uuid
go get golang.org/x/crypto/bcrypt
~~~

To build application docker image
~~~
docker build -t go-auth-server .
~~~

### Docker compose usage

To start project
~~~
docker-compose up
~~~

To shutdown project
~~~
docker-compose down
~~~


