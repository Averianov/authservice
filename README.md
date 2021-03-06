# AuthService

Trial project.

Note: The project assumes a pre-installed Mongo database.

The service takes two routes. First route - /auth/login - get guid user data, validate and, if ok returned the couple access and refresh tokens. Refresh token returned as protected cookie. Second route - /auth/refresh - get refresh token, validate him, compare with data from DB and returned the new couple access and refresh tokens.

Any results will return in response as JSON message.

### For launch:

* Go to home directory:

For Windows:
```bash
cd %UserProfile%/go/src
```
For Linux:
```bash
cd $HOME/go/src
```

* Get application from git:
```bash
git clone https://github.com/Averianov/authservice
cd ./authservice
```

* Check variables in .env file:

Variables: url_db - url to mongo DB; domain - dn for application and for secure tokens; token_password - password for encode token

For example:
```cfg
url_db = 127.0.0.1:27017
domain = myapp.domain.com
token_password = SomeVerySecretPasswordForTokens
```

* For init dependencies if need:
```bash
go get -u ./...
```

* Build application:
```bash
go build main.go
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
