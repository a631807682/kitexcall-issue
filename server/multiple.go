package server

import (
	"context"
	"kitexcall-issue/kitex_gen/multiple"
	"kitexcall-issue/kitex_gen/multiple/multiple1"
	"kitexcall-issue/kitex_gen/multiple/multiple2"
	"net"

	"github.com/cloudwego/kitex/server"
)

func RunMultipleServer(address string) error {
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return err
	}
	srv := server.NewServer(
		server.WithServiceAddr(addr),
	)

	srv.RegisterService(multiple1.NewServiceInfo(), &server1{})
	srv.RegisterService(multiple2.NewServiceInfo(), &server2{})
	return srv.Run()
}

var _ multiple.Multiple1 = &server1{}

type server1 struct{}

// CMD:
//
// kitexcall -idl-path  idl/multiple/multiple.proto -m Multiple1/Method1 -d '{}' -e 127.0.0.1:10002
//
// Result:
// [Status]: Failed
// [ServerError]: RPC call failed: missing method: Method1 in service
func (server1) Method1(ctx context.Context, req *multiple.Request) (*multiple.Response, error) {
	return &multiple.Response{}, nil
}

var _ multiple.Multiple2 = &server2{}

type server2 struct{}

// CMD:
//
// kitexcall -idl-path  idl/multiple/multiple.proto -m Multiple2/Method2 -d '{}' -e 127.0.0.1:10002
//
// Result:
// [Status]: Failed
// [ServerError]: RPC call failed: remote or network error[remote]: unknown method Method2
func (server2) Method2(ctx context.Context, req *multiple.Request) (*multiple.Response, error) {
	return &multiple.Response{}, nil
}
