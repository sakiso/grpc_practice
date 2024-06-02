package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	hellopb "mygrpc/pkg/grpc" // protoファイルから生成されたコードをimport

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// scanner *bufio.Scanner
	client hellopb.GreetingServiceClient
)

func main() {
	fmt.Println("start gRPC client")

	// 1. 標準入力からの読み込みを行うためのscannerを作成
	scanner := bufio.NewScanner(os.Stdin)

	// 2. gRPCサーバーに接続
	address := "localhost:8080"
	conn, err := grpc.Dial( // Dialでコネクションを張る(conn = コネクションのこと)
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
		return
	}

	defer conn.Close()

	// 3. gRPCサーバーのクライアントを作成
	client = hellopb.NewGreetingServiceClient(conn)

	for {
		fmt.Println("1: send request")
		fmt.Println("2: exit")
		fmt.Print("prealse enter > ")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			CallHello()
		case "2":
			fmt.Println("exit")
			goto M
		}
	}
M:
}

func CallHello() {
	fmt.Print("名前を入力してください: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	req := &hellopb.HelloRequest{
		Name: name,
	}

	res, err := client.Hello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetMessage())
	}
}
