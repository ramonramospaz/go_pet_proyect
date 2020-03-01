# Go pet proyect 1

## Info
Is a simple web service that use an in memory struct that save basic information.

## Create docker image
```bash
docker build -t go-petproject1 .
```

## Run docker image
```bash
docker run --name go-petproject1 -p 8081:8081 go-petproject1
```

## Update code in haroku
```bash
git push heroku master
```

# Services
>*LIST OF USER:* Show a list of users that are in the memory.
>*TYPE:* GET
>*URL:* http://server_ip:server_port/user
>*RESPONSE:*
```json
[
    {
        "id": 15281565,
        "name": "Ramon Ramos",
        "age": 38,
        "birthCountry": "Venezuela"
    },
    {
        "id": 15281567,
        "name": "Fabiola Ramos",
        "age": 37,
        "birthCountry": "Venezuela"
    }
]
```
>*CREATE USER:* Include a new user in the memory database.
>*TYPE:* POST
>*URL:*  http://server_ip:server_port/user
>*BODY:*
```json
{
        "id":15281565,          
        "name":"Ramon Ramos",        
        "age":38,        
        "birthCountry":"Venezuela"
}
```
>*RESPONSE:* 
```json
{
    "code": 200,
    "message": "User Created"
}
```
>*FIND USER:* Find an user in the memory database.
>*TYPE:* GET
>*URL:* http://server_ip:server_port/user/ID
>*RESPONSE:*
```json
{
        "id": 15281565,
        "name": "Ramon Ramos",
        "age": 38,
        "birthCountry": "Venezuela"
}
```
    or
```json
{
        "code": 400,
        "message": "User dont exist"
}
```
>*DELETE AN SPECIFIC USER:* Find and Delete an user in the memory database.
>*TYPE:* DELETE
>*URL:* http://server_ip:server_port/user/ID
>*RESPONSE:*
```json
{
    "code": 200,
    "message": "The user deleted"
}
```