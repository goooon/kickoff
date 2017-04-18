package main

import (
	"fmt"
	//"main/utils"
	//	"main/utils"
	"main/rpc/thrift"
	"os"
	"reflect"
	"time"
)

type Node struct {
	expire uint32
	f      func()
}

func getInt(i int32) (int32, int32) {
	return 1 + i, 2
}

func (n *Node) String() string {
	return fmt.Sprintf("Node:expire,%d", n.expire)
}

func Producer(id int, item chan int) {
	for i := 0; i < 10; i++ {
		item <- i
		fmt.Printf("Producer %d produces data: %d\n", id, i)
		time.Sleep(10 * 1e6)
	}
}

func Consumer(id int, item chan int) {
	for i := 0; i < 20; i++ {
		c_item := <-item
		fmt.Printf("Consumer %d get data: %d\n", id, c_item)
		time.Sleep(10 * 1e6)
	}
}

func main() {
	args := os.Args

	fmt.Println(args)
	fmt.Println(cap(args))

	for v := range args {
		fmt.Println(v)
	}

	//	utils.Query()

	f := getInt
	fmt.Println(f(3))
	item := make(chan int, 6)

	go Producer(1, item)
	go Producer(2, item)
	go Consumer(1, item)
	fmt.Println(getInt(3))
	time.Sleep(1 * 1e9)
	//utils.Main()
	thrift.ClientMain("a", "b")
}
