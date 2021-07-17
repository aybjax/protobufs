package main

import (
	"fmt"

	pb "protobufs/ayb.jax"

	"google.golang.org/protobuf/proto"
)

func main() {
	p := &pb.Person{
		Id: 1234,
		Name: "Roger F",
		Email: "rf@example.com",
		Phones: []*pb.PhoneNumber {
			{Number: "555-4321", Type: pb.PhoneType_HOME},
		},
	}

	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)

	fmt.Println("Original struct loaded from proto file:", p)
	fmt.Println("Marshaled proto data:", body)
	fmt.Println("Unmarshaled struct:", p1)
}