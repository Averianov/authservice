# AuthService

Test project for the position of Junior Backend Developer at MEDODS Ltd.

Note: The project assumes a pre-installed Mongo database.

The service takes two routes. First route - /auth/login - get guid user data, validate and, if ok returned the couple access and refresh tokens. Refresh token returned as protected cookie. Second route - /auth/refresh - get refresh token, validate him, compare with data from DB and returned the new couple access and refresh tokens.

Any results will return in response as JSON message.

### For launch:

* Get application from git:
```bash
git clone https://github.com/Averianov/authservice
cd ./authservice
```

* Check variables in .env file:

Variables: url_db - url to mongo DB, domain - dn for application and for secure tokens, token_password - password for encode token\n

For example:
```cfg
url_db = 127.0.0.1:27017
domain = myapp.domain.com
token_password = SomeVerySecretPasswordForTokens
```

* For test application controller:
```bash
go test ./controllers
```

* Init/update dependencies and build application:
```bash
go get -u ./...
go build .
```

* launch application:

For windows:
```bash
authservice.exe
```
For linux:
```bash
./authservice
```
### Example requests:

##### Login Example:
```bash
curl -d "{\"guid\":\"6F9619FF-8B86-D011-B42D-00CF4FC964FF\"}" -X POST http://localhost:8080/auth/login -H "Content-Type:application/json" -v
```

##### Refresh Example (NOTE: You must insert the bearer key you got from the previous response):
```bash
curl -X GET http://localhost:8080/auth/refresh -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SWQiOjF9.LrpWOP5Gi7Xn-vq-XBvR7dvnt-w8ZlhOS2qVfdv0t_M" -v
```