# AuthService

Test project for the position of Junior Backend Developer at MEDODS Ltd.

Note: The project assumes a pre-installed Mongo database.

The service takes two routes. First route - /auth/login - get guid user data, validate and, if ok returned the couple access and refresh tokens. Refresh token returned as protected cookie. Second route - /auth/refresh - get refresh token, validate him, compare with data from DB and returned the new couple access and refresh tokens.

### For launch:

* Go to application workdirectory:
```bash
 cd ./some_application
```

* create main.go file. For example:
```go
package main

import (
	"github.com/Averianov/authservice"
)

func main() {
	err := authservice.Run()
	if err != nil {
		panic(err)
	}
}
```

* Create .env file. For example:
```cfg
domain = localhost
url_db = localhost:27017
token_password = $omE_e}{ample_$ecreT
```
* Init dependencies:
```bash
go mod init myappname
```
* Build application:
```bash
go build
```
* launch application:
```bash
myappname
```
### Example requests:

##### Login Example:
```bash
curl -d "{\"guid\":\"6F9619FF-8B86-D011-B42D-00CF4FC964FF\"}" -X POST http://localhost:8080/auth/login -H "Content-Type:application/json"
```
Note: this request can't take returning refresh token as cookie.

##### Refresh Example:
```bash
curl -X GET http://localhost:8080/auth/refresh -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SWQiOjF9.LrpWOP5Gi7Xn-vq-XBvR7dvnt-w8ZlhOS2qVfdv0t_M"
```