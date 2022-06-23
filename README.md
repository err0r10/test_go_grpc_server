# test_go_grpc_server

Run Server
1. Copy ```git clone git clone https://github.com/err0r10/test_go_grpc_server.git```
2. ```cd test_go_grpc_service```
3. Build ```docker build -t test_go_grpc_service ./server.go```
4. Run docker container ```docker run -d -p 5003:5003 test_go_grpc_service```

Run Client
1. ```cd test_go_grpc_service```
2. ```go run client.go```
