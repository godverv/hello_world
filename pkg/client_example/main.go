package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/godverv/hello_world/pkg/api"
)

func main() {
	c, err := grpc.NewClient(":80",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatal(err)
	}

	apiClient := api.NewHelloWorldAPIClient(c)
	resp, err := apiClient.Version(context.Background(), &api.Version_Request{})
	if err != nil {
		logrus.Fatal(err.Error())
	} else {
		logrus.Println(resp)
	}

}
