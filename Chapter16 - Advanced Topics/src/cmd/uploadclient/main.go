package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
	"movieexample.com/gen"
)

func main() {
	conn, err := grpc.NewClient("localhost:8083")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := gen.NewMovieServiceClient(conn)

	filePath := "upload.txt"
	if err := uploadFile(client, filePath); err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}
}

func uploadFile(client gen.MovieServiceClient, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	stream, err := client.UploadFile(context.Background())
	if err != nil {
		return fmt.Errorf("failed to create upload stream: %w", err)
	}

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}

		if err := stream.Send(&gen.UploadRequest{
			Filename: filePath,
			Chunk:    buf[:n],
		}); err != nil {
			return fmt.Errorf("failed to send chunk: %w", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("error receiving response: %w", err)
	}

	fmt.Println("Server response:", resp.Message)
	return nil
}
