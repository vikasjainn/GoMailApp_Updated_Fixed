package main

import (
    "context"
    "GoMailApp_Updated_Fixed/Proto"
    "google.golang.org/grpc"
    "log"
    "time"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := Proto.NewEmailServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()

    response, err := client.SendEmail(ctx, &Proto.UserInput{Name: "John Doe", Email: "vikas.jain@bugsmirror.com"})
    if err != nil {
        log.Fatalf("Error calling SendEmail: %v", err)
    }

    log.Printf("Response: Success=%v, Message=%s", response.Success, response.Message)
}
