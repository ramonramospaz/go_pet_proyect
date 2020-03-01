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
>*LIST OF USER:* Show a list of users that are in the memory.<br>
>*TYPE:* GET<br>
>*URL:* http://server_ip:server_port/user<br>
>*RESPONSE:*<br>
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
>*CREATE USER:* Include a new user in the memory database.<br>
>*TYPE:* POST<br>
>*URL:*  http://server_ip:server_port/user<br>
>*BODY:*<br>
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
>*FIND USER:* Find an user in the memory database. <br>
>*TYPE:* GET<br>
>*URL:* http://server_ip:server_port/user/ID<br>
>*RESPONSE:*<br>
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
>*DELETE AN SPECIFIC USER:* Find and Delete an user in the memory database.<br>
>*TYPE:* DELETE<br>
>*URL:* http://server_ip:server_port/user/ID<br>
>*RESPONSE:*<br>
```json
{
    "code": 200,
    "message": "The user deleted"
}
```