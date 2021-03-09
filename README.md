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


### Environment variables

|Variable Name | Default value | Description|
|-------------:|--------------:|-----------:|
|SERVER_PORT |8080 | Port, which our application listens |
|MONGO_HOST|localhost|The name of host  to connect Mongo DB|
|MONGO_PORT|27017|The number of port to connect Mongo DB|
|MONGO_USERNAME|mongoroot|Mongo DB authorization username|
|MONGO_PASSWORD|mongopass|Mongo DB authorization password|
|MONGO_DATABASE|admin|Mongo DB database|
|SECRET|0000-0000-0000-000| The secret value used in Crypto algorithm of token signature|
|ACCESS_TOKEN_EXPIRATION_MS|5 * 60 * 1000| Duration in milliseconds from access been issued till its expiration |
|REFRESH_TOKEN_EXPIRATION_MS|24 * 60 * 60 * 1000| Duration in milliseconds from access been issued till its expiration |


## REST API

## New user account creation

Method creates new user account

~~~
Method: POST 
url: http://localhost:8080/api/user/register

Headers:
Content-Type: application/json
Accept: application/json

Request Body

{
 "id":"8860b036-5fa0-458d-ba38-1928585a21a8",
"password":"1111"
}


Response:

{
  "message": "success",
  "timestamp": "2021-03-09T07:51:29.212Z"
}
~~~

## User Authentication

Method returns access and refresh tokens 


~~~
Method: POST 
url: http://localhost:8080/api/user/login

Headers:
Content-Type: application/json
Accept: application/json

Request Body

{
 "id":"8860b036-5fa0-458d-ba38-1928585a21a8",
"password":"1111"
}


Response:

{
  "authentication": {
    "id": "8860b036-5fa0-458d-ba38-1928585a21a8",
    "refreshToken": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzQXQiOjE2MTUzNjMwMDA2MDI2NzA5MDAsImlkIjoiNjA0NzJhMzg0NmJmNmM5YWFhYWZjMmI5IiwidXNlcklkIjoiODg2MGIwMzYtNWZhMC00NThkLWJhMzgtMTkyODU4NWEyMWE4In0.EfQHMI2qF_myX-n6ouiw-bJQbqFm5cQziTXAJNcaavk1YS2KCT249YvGKlhjD3-CBNJpVHSglRrZlcjaj9hKsQ",
    "accessToken": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzQXQiOjE2MTUyNzY5MDA2MDI1MDIyMDAsInVzZXJJZCI6Ijg4NjBiMDM2LTVmYTAtNDU4ZC1iYTM4LTE5Mjg1ODVhMjFhOCJ9.lBB1jScWqfgFqUzk3n_h_1icqPzUZoQGRUdAG54uxJy-U1y7iaZuRFhd0cqyekdmzIMGaPxK5Z18kYU4jirDTQ"
  },
  "message": "success",
  "timestamp": "2021-03-09T07:56:40.662Z"
}
~~~

## Refresh user access

Method refreshes access token which might be expired
It needs to present existent refresh token in request body
Any refresh token may be used only once
New refresh and access token are returned in response body


~~~
Method: POST 
url: http://localhost:8080/api/user/refresh

Headers:
Content-Type: application/json
Accept: application/json

Request Body

{
 "id":"8860b036-5fa0-458d-ba38-1928585a21a8",
"refreshToken": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzQXQiOjE2MTUzNjMwMDA2MDI2NzA5MDAsImlkIjoiNjA0NzJhMzg0NmJmNmM5YWFhYWZjMmI5IiwidXNlcklkIjoiODg2MGIwMzYtNWZhMC00NThkLWJhMzgtMTkyODU4NWEyMWE4In0.EfQHMI2qF_myX-n6ouiw-bJQbqFm5cQziTXAJNcaavk1YS2KCT249YvGKlhjD3-CBNJpVHSglRrZlcjaj9hKsQ"
}


Response:

{
  "authentication": {
    "id": "8860b036-5fa0-458d-ba38-1928585a21a8",
    "refreshToken": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzQXQiOjE2MTUzNjMwMDA2MDI2NzA5MDAsImlkIjoiNjA0NzJhMzg0NmJmNmM5YWFhYWZjMmI5IiwidXNlcklkIjoiODg2MGIwMzYtNWZhMC00NThkLWJhMzgtMTkyODU4NWEyMWE4In0.EfQHMI2qF_myX-n6ouiw-bJQbqFm5cQziTXAJNcaavk1YS2KCT249YvGKlhjD3-CBNJpVHSglRrZlcjaj9hKsQ",
    "accessToken": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzQXQiOjE2MTUyNzY5MDA2MDI1MDIyMDAsInVzZXJJZCI6Ijg4NjBiMDM2LTVmYTAtNDU4ZC1iYTM4LTE5Mjg1ODVhMjFhOCJ9.lBB1jScWqfgFqUzk3n_h_1icqPzUZoQGRUdAG54uxJy-U1y7iaZuRFhd0cqyekdmzIMGaPxK5Z18kYU4jirDTQ"
  },
  "message": "success",
  "timestamp": "2021-03-09T07:56:40.662Z"
}
~~~

## Logout 

This method intended to clear existent refresh token. 
To use this method user must be authenticated
so It needs to present existent valid access token in Authorization header  


~~~
Method: POST 
url: http://localhost:8080/api/user/logout

Headers:
Authorization: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzQXQiOjE2MTUyNzY5MzMyMTIyMTk1MDAsInVzZXJJZCI6Ijg4NjBiMDM2LTVmYTAtNDU4ZC1iYTM4LTE5Mjg1ODVhMjFhOCJ9.3O2NP1_q5Bq5tNzfYNzQJ-Zqi804h2ZNjXDYip_6ggrpU9E0g6VQ_SEHJvggfmIUFgVbVz3ExRFWXRr3u5dcnw
Accept: application/json

Response:

{
  "message": "success",
  "timestamp": "2021-03-09T07:57:57.603Z"
}
~~~
