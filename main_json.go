package main

import (
	"encoding/json"
	"fmt"

	pb "protobufs/ayb.jax"
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

	body, _ := json.Marshal(p)

	fmt.Println(string(body))
}