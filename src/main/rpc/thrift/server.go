package thrift

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"main/rpc/thrift/rpc"
	"os"
)

type RpcServiceImpl struct {
}

func (this *RpcServiceImpl) FunCall(callTime int64, funCode string, paramMap map[string]string) (r []string, err error) {
	fmt.Println("-->FunCall:", callTime, funCode, paramMap)

	for k, v := range paramMap {
		r = append(r, k+v)
	}
	return
}

func ServerMain(ip string, port string) {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(ip + ":" + port)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &RpcServiceImpl{}
	processor := rpc.NewRpcServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	//fmt.Println("thrift server in", NetworkAddr)
	server.Serve()
}
