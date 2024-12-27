
## REST API Client  
[Extension link](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)


How to make GET Request 
```
### Get all users
GET http://localhost:3000/users
Accept: application/json
```

How to make POST Request with payload
```
POST http://localhost:3000/register
Content-Type: application/json

{
    "name": "Demo User",
    "email" : "demo.user@gmail.com",
    "password" : "Hello@123"
}
```

Using Variables
```
@baseUrl = http://localhost:3000
```

Naming Request
```
# Login
# @name tokenAPI
POST {{baseUrl}}/login
Content-Type: application/json

{
    "email" : "hello.123@gmail.com",
    "password": "Hello@123"
}
```

Storing the Response
```
@authToken = {{tokenAPI.response.body.token}}
```

Adding Authentication Headers
```
### Get my user profile
GET {{baseUrl}}/users/me
Content-Type: application/json
Authorization: Bearer {{authToken}}
```



