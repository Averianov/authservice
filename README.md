# AuthService
Test project for the position of Junior Backend Developer at MEDODS Ltd

### Two requests:

##### Login Example:
* curl -d "{\"guid\":\"6F9619FF-8B86-D011-B42D-00CF4FC964FF\"}" -X POST http://localhost:8080/auth/login -H "Content-Type:application/json"

##### Refresh Example:
* curl -X GET http://localhost:8080/auth/refresh -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SWQiOjF9.LrpWOP5Gi7Xn-vq-XBvR7dvnt-w8ZlhOS2qVfdv0t_M"
