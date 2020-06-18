package main

import (
	"flag"
	"fmt"
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/server"
	"log"
)

func main() {
	db := flag.String("database", "jsonDB", "possible database type. Valid values are `jsonDB` and `redis`.")
	flag.Parse()

	//go func() {
	//	time.Sleep(5 * time.Second)
	//	fmt.Println("Client sending request now")
	//	config := &tls.Config{InsecureSkipVerify: true}
	//	conn, err := grpc.Dial("localhost:8443", grpc.WithTransportCredentials(credentials.NewTLS(config)))
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	client := v1.NewSmartEnergyTableServiceClient(conn)
	//	resp, err := client.CreateRoom(context.TODO(), &v1.Empty{})
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	resp.UserId = uuid.New().String()
	//	stream, err := client.JoinRoom(context.TODO(), &v1.RoomUser{UserId: resp.UserId, Id: resp.Id})
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	go func(stream v1.SmartEnergyTableService_JoinRoomClient) {
	//		for {
	//			patch, err := stream.Recv()
	//			if err != nil {
	//				log.Println(err)
	//				break
	//			}
	//			fmt.Println(patch)
	//		}
	//		fmt.Println("Closing client")
	//	}(stream)
	//	time.Sleep(time.Second)
	//	client.SaveRoom(context.TODO(), &v1.RoomUser{UserId: resp.UserId, Id: resp.Id})
	//	time.Sleep(time.Second)
	//	client.GetEnergyData(context.TODO(), &v1.RoomUser{UserId: resp.UserId, Id: resp.Id})
	//	time.Sleep(time.Second)
	//	client.LeaveRoom(context.TODO(), &v1.RoomUser{UserId: resp.UserId, Id: resp.Id})
	//	time.Sleep(time.Second)
	//}()

	fmt.Println("using database:", *db)
	if err := server.Run(*db); err != nil {
		log.Fatal(err)
	}
}
