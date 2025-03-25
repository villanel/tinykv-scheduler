package main

import (
	"context"
	"log"
	"strings"

	"github.com/villanel/tinykv-scheduler/kv/raftstore/scheduler_client"
)

func main() {
	schedulerClient, err := scheduler_client.NewClient(strings.Split("192.168.2.18:2379", ","), "")
	if err != nil {
		panic(err)
	}
	//check if   up
	tem := 7
	st := uint64(tem)
	store, err := schedulerClient.GetStoreState(context.TODO(), st)
	print(store.Capacity)
	if err != nil {
		log.Print(err.Error())
	}
}
