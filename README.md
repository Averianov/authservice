# AuthService

Test project for the position of Junior Backend Developer at MEDODS Ltd.

Note: The project assumes a pre-installed Mongo database.

### For launch:

* Go to application workdirectory:
```bash
 cd ./some_application
```

* create main.go. For example:
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
```env
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
./myappname
```
### Example requests:

##### Login Example:
```bash
curl -d "{\"guid\":\"6F9619FF-8B86-D011-B42D-00CF4FC964FF\"}" -X POST http://localhost:8080/auth/login -H "Content-Type:application/json"
```

##### Refresh Example:
```bash
curl -X GET http://localhost:8080/auth/refresh -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SWQiOjF9.LrpWOP5Gi7Xn-vq-XBvR7dvnt-w8ZlhOS2qVfdv0t_M"
```