package main

import (
    "context"
    "io/ioutil"
    "log"
    "os"
    "time"

    "google.golang.org/grpc"
    pb "github.com/SarathViswa09/image_search" // Update this import path to your generated package
)

const (
    address     = "localhost:50051"
    defaultName = "dog"
)

func main() {
    // Set up a connection to the server.
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewImageSearchServiceClient(conn)

    // Contact the server and print out its response.
    keyword := defaultName
    if len(os.Args) > 1 {
        keyword = os.Args[1]
    }
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.SearchImage(ctx, &pb.SearchRequest{Keyword: keyword})
    if err != nil {
        log.Fatalf("could not search: %v", err)
    }

    // Save the image to the 'received_output' directory
    outputPath := "received_output/" + keyword + "_output.jpg"
    if err := ioutil.WriteFile(outputPath, r.GetImage(), 0644); err != nil {
        log.Fatalf("failed to write image to file: %v", err)
    }

    log.Printf("Image saved to %s", outputPath)
}
