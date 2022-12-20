package main

import (
	grpc "interact-cosmos-data/grpc"
)

func main() {
	err, err2 := grpc.QueryState()
	if err != nil {
		println("error occured!! ", err.Error(), err2.Error())
	}
}
