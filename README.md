# cmpe273-lab2

#Start the server
go get github.com/aggarwalsomya/cmpe273-lab2
cd cmpe273-lab2
go run *


# Making requests

Use curl to make GET and POST requests (Mac)

#GET REQUEST
curl.exe -H GET http://localhost:8080/hello/Somya

#POST REQUEST
curl.exe -H POST -d '{"name":"Somya Aggarwal"}' http://localhost:8080/hello
