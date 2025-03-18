package main

import (
	"context"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/villanel/tinykv-scheduler/kv/raftstore/scheduler_client"
	"github.com/villanel/tinykv-scheduler/proto/pkg/metapb"
)

func main() {
	schedulerClient, err := scheduler_client.NewClient(strings.Split("192.168.2.18:2379", ","), "")
	if err != nil {
		panic(err)
	}
	//check if   up
	tem := 2
	stores, err := schedulerClient.GetAllStores(context.TODO())
	if err != nil {
		log.Print(err.Error())
	}
	if stores == nil {
		log.Print("stores not found")
	}
	storeLink := map[int]uint64{}
	for _, store := range stores {
		str := store.Address
		re := regexp.MustCompile(`tinykv-(\d+)`)
		matches := re.FindStringSubmatch(str)
		if len(matches) >= 2 {
			num, err := strconv.Atoi(matches[1])
			if err != nil {
				panic("转换失败: " + err.Error())
			}
			storeLink[num] = store.GetId()
		}
	}
	st := storeLink[tem]
	store, err := schedulerClient.GetStore(context.TODO(), st)
	if err != nil {
		log.Print(err.Error())
	}
	if store.GetState() == metapb.StoreState_Up {

		schedulerClient.OfflineStore(context.TODO(), st)

	} else if store.GetState() == metapb.StoreState_Tombstone {
		schedulerClient.RemoveStore(context.TODO(), st)
	}
	println(store.GetState().String())
	schedulerClient.Close()
}
