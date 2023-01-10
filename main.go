package main

import "interact-cosmos-data/grpc"

func main() {
	//err, err2 := grpc.QueryState()
	//err, err2 := grpc.QueryState02()
	//err := grpc.QueryState03()
	//err, _ := grpc.QueryState04()
	//err := grpc.QueryState05()
	//err := grpc.QueryState06()
	err := grpc.QueryState07()
	//err := rest.QueryRest01()
	if err != nil {
		println("error occured!! ", err.Error())
	}
}
