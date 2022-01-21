package main

import (
	controller "github.com/AndyMile/articles/api/controllers"
	router "github.com/AndyMile/articles/api/provider/rest"
	pb "github.com/AndyMile/articles/app/proto"
	"google.golang.org/grpc"
) 

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	
	client := pb.NewArticleClient(conn)

    c := controller.NewBaseHandler(client)

	router.CreateRouter(c)
}