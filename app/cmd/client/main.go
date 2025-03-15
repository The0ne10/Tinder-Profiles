package main

import (
	pb "Profiles/my-helloservice/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	names := []string{"Андрей", "Слава", "Ваня", "Владлен"}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := pb.NewMyServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, name := range names {
		res, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("Ошибка запроса: %v", err)
		}
		log.Printf("Ответ сервера: %s", res.Message)
	}
}
