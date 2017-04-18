package main

import (
	"fmt"
	"main/rpc/thrift"
	"testing"
)

func Test_Main(t *testing.T) {
	fmt.Println("start testing -----")
	t.Log("test ok")
	thrift.ClientMain("a", "b")
	fmt.Println("stop testing -----")
}
