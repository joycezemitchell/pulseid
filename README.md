John Mitchell Villanueva - PulseId Exam
==
Notes
--
Once you run the API server, it will automatically create two tables in the database and a default admin user (username: pulseid, password: qwer1234)
```
admin
client_tokens 
```

1. **admin** table will contain all the admin credentials.
2. **client_tokens** table will contain all the invitation tokens.





Setup
--


1. Git clone this repo.
2. Create a mysql database 
    ```
    CREATE DATABASE pulseid
    ```
3. Run the following in the terminal
    ```
    export mysql_users_password="your database password"
    export mysql_users_schema="your database name"
    export mysql_users_host="localhost"
    export mysql_users_username="your database username"
    ```
4. In the project root folder, run the following to start the API server
    ```
    go run main.go
    ```
    
Generating an admin token for authentication
--    

To generate a token, run the following in Postman
    ```
    POST http://localhost:7070/login
    ```
    
    
     {
        "username":"pulseid",
        "password":"qwer1234"
     }

The screen should look something like this
![Alt text](https://github.com/joycezemitchell/pulseid/blob/master/images/im1.jpg)



Generating an invite token
--    

To generate an invite token, run the following in Postman
```
GET http://localhost:7070/requesttoken
```

In Header, you will need to add the Authorization with the generated token

![Alt text](https://github.com/joycezemitchell/pulseid/blob/master/images/im2.jpg)


The screen should look something like this

![Alt text](https://github.com/joycezemitchell/pulseid/blob/master/images/im3.JPG)

```
{
    "id": 22,
    "token": "9nd9TX",
    "status": "Active"
}
```

A new record will be added in client_token table. The generated token will be converted into md5 hash for security purposes.
![Alt text](https://github.com/joycezemitchell/pulseid/blob/master/images/im4.jpg)


Validating an invite token
--    

To validate an invite token, run the following in Postman

```
GET http://localhost:7070/validateToken
```

Here is a sample json request

```
{
    "token":"eAvBGO"
}
```




![Alt text](https://github.com/joycezemitchell/pulseid/blob/master/images/im5.jpg)

If the invite token is valid, it will return the following result

![Alt text](https://github.com/joycezemitchell/pulseid/blob/master/images/im6.jpg)


Listing all invite tokens
-- 
To list all the generated invite tokens, run the following in Postman
```
GET http://localhost:7070/getalltokens
```

In Header, you will need to add the Authorization with the generated token

![Alt text](https://github.com/joycezemitchell/pulseid/blob/master/images/im7.jpg)


The result should look something like this with the list of all the tokens in md5 hash
![Alt text](https://github.com/joycezemitchell/pulseid/blob/master/images/im8.jpg)


Swagger
-- 

To view the Swagger documentation, open a browser and access this url

```
http://localhost:7070/swagger/index.html#/

```

The screen should be similar to this:

![Alt text](https://github.com/joycezemitchell/pulseid/blob/master/images/im9.jpg)



Basic functionality requirements
-- 

> The API should be REST-full

I have used **Golang GIN** framework to create a REST-full enpoint

> The admin enpoint should be authenticated. Proposed an easy auth mechanism. 

For admin authentication, I have used JWT and admin credentials(username and password - md5 encrypted).
You will need to generate an access token (JWT) before you can access all secured endpoint like enpoint for generating token invites.

> Invite tokens to expire after 7 days
> Invite tokens can be realled disabled

To validate if a token has expired or if token has been disabled, I have used the following SQL
```
SELECT 
  id, status, created_at 
FROM 
  client_tokens
WHERE
  token = ? AND
  status = 'Active' AND
  created_at >= (now() - INTERVAL 8 DAY)

```

> A public endpoint for validating the invite token

To validate a token, you can to go to this endpoint

```
GET http://localhost:7070/validateToken
```

Nice to have funcional requirements
-- 

To view all the invite tokens, you can go to 

```
GET http://localhost:7070/getalltokens
```


Basic nun functional requirements
-- 

I believed I have met all the basic non-functional requirements inclunding the Nice to have nonfuncatinal requirements.

> Design and document the APIS the will facilitate the workflow outlined above

This GIT readme will serve as my documentaion incluing the Swagger API documentation that I added.

> Develop the API in GO Lang

All the code in this repo are written in golang

> Use any framework or library that will help you develop the solutions

I have used the following packages for this projet

```
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/fsnotify/fsnotify v1.5.0 // indirect
	github.com/gin-gonic/gin v1.7.4
	github.com/go-openapi/analysis v0.20.1 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/runtime v0.19.30 // indirect
	github.com/go-openapi/strfmt v0.20.2 // indirect
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/go-swagger/go-swagger v0.27.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/viper v1.8.1 // indirect
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	github.com/swaggo/gin-swagger v1.3.1
	github.com/swaggo/swag v1.7.1
	github.com/ugorji/go v1.2.6 // indirect
	github.com/urfave/cli v1.22.5 // indirect
	go.mongodb.org/mongo-driver v1.7.1 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/mod v0.5.0 // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.5 // indirect
	google.golang.org/protobuf v1.27.1 // indirect

```

> Make sure your code is well-fomatted, clean and follows the best practice.

I believed that my have met all the above requirements

> Seperate concerns

I have used the classic MVC patter for the projects

> Document the APIs in Swagger or similar tool.

I have used go-swagger package to incorporate Swagger documentation

> Use an actuall DB.

I have used MySQL to save all the invite tokens and to store admin credentials.

> Provide deployment instructaions

The Setup section above will serve as my deployment instructions










