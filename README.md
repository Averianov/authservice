# AuthService
Тестовый проект на позицию Junior Backend Developer компании MEDODS

Проект предполагает предустановленной базы данных Mongo. 

### For run:
Go to user`s GOPATH:
* cd %USERPROFILE%/go/src

Download repository
* go get https://github.com/Averianov/authservice.git@v0.0.1
* git clone https://github.com/Averianov/authservice.git
* cd ./authservice

Edit .env file as you need

Download dependencies
* go env -w GO111MODULE=on && go mod vendor

Build project
* cd ./cmd && go build

launch authservice.exe

### Two requests:

##### Login Example:
* curl -d "{\"guid\":\"6F9619FF-8B86-D011-B42D-00CF4FC964FF\"}" -X POST http://localhost:8080/auth/login -H "Content-Type:application/json"

##### Refresh Example:
* curl -X GET http://localhost:8080/auth/refresh -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SWQiOjF9.LrpWOP5Gi7Xn-vq-XBvR7dvnt-w8ZlhOS2qVfdv0t_M"
