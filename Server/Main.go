package main

import (
    "GoMailApp_Updated_Fixed/Api"
    "GoMailApp_Updated_Fixed/Common"
    "GoMailApp_Updated_Fixed/Proto"
    "google.golang.org/grpc"
    "log"
    "net"
)

func main() {
    Common.InitFirebase()

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    Proto.RegisterEmailServiceServer(grpcServer, &Api.EmailServiceServer{})

    log.Println("Server listening on port 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
