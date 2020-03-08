# insurance-gRPC-server

* How to run in local: 
     
     You have to clone this repository in the typical go workspace path: 
     
     ~/go/src/github.com$ `git clone https://github.com/Steampunk1453/insurance-gRPC-server`
     
     Get all project dependencies: ~/go/src/github.com/insurance-gRPC-server `go get ./...` 
     
     Run fake AMS API: ~/go/src/github.com/insurance-gRPC-server`go run api/ams_api.go` 
     
     Run server ~/go/src/github.com/insurance-gRPC-server `go run server/server.go` 
     
* To test gRPC server endpoints use this gRPC client with test data:
      
     Wait delay time (1 minute) and run client: ~/go/src/github.com/insurance-gRPC-server`go run client/client.go` 
      
* TODO
    
    Implement unit tests with mocks
     
    
