package server

import (
	"context"
	echo "kitexcall-issue/kitex_gen/echo"
	e "kitexcall-issue/kitex_gen/echo/echo"
	"net"

	"github.com/cloudwego/kitex/server"
)

func RunEchoServer(address string) error {
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return err
	}
	srv := e.NewServer(
		&echoImpl{},
		server.WithServiceAddr(addr),
	)
	return srv.Run()
}

var _ echo.Echo = &echoImpl{}

type echoImpl struct{}

// CMD:
//
// kitexcall -idl-path  idl/echo/echo.proto -m Echo/ConvertSint64 -d '{"message":"test", "num":2}' -e 127.0.0.1:10001
//
// Result:
//
// [Status]: Failed
// [ServerError]: RPC call failed: remote or network error: protobuf marshal message failed: [CATEGORY 0] convert failed: json convert to protobuf failed, MessageDescriptor:
// [JSON-TO-PROTOBUF] convert failed: sonic decode json bytes failed
// [JSON-TO-PROTOBUF] dismatched type: param isn't intType
func (echoImpl) ConvertSint64(ctx context.Context, req *echo.ConvertSint64Request) (*echo.ConvertSint64Response, error) {
	return &echo.ConvertSint64Response{
		Message: req.Message,
		Num:     req.Num,
	}, nil
}
